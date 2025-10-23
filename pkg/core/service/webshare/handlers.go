package webshare

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

//go:embed all:ui
var ui embed.FS

func rcfilepath(folder string, name string) string {
	newName := name
	ext := path.Ext(name)
	stem := name[:len(name)-len(ext)]
	for i := range 1_000_000 {
		if i != 0 {
			newName = fmt.Sprintf("%s-%d%s", stem, i, ext)
		}
		dest := path.Join(folder, newName)
		_, err := os.Stat(dest)
		if err != nil {
			return dest
		}
	}
	log.Errorf("Error: could not get file output path in downloads")
	return name
}

func (s *Service) handleGetSharedFiles(w http.ResponseWriter, r *http.Request) {
	s.sharedFilesLock.RLock()
	defer s.sharedFilesLock.RUnlock()

	if !s.isValidSession(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.sharedFiles)
}

func (s *Service) handleDownloadFile(w http.ResponseWriter, r *http.Request) {
	if !s.isValidSession(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fileUUID := r.URL.Query().Get("uuid")
	if fileUUID == "" {
		http.Error(w, "Missing file uuid", http.StatusBadRequest)
		return
	}

	s.sharedFilesLock.RLock()
	var filePath string
	for _, f := range s.sharedFiles {
		if f.UUID == fileUUID {
			filePath = f.Path
			break
		}
	}
	s.sharedFilesLock.RUnlock()

	if filePath == "" {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filePath)
}

func (s *Service) handleUploadFile(w http.ResponseWriter, r *http.Request) {
	if !s.isValidSession(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	maxFileSize := int64(s.config.MaxSize) * 1024 * 1024
	if maxFileSize == 0 {
		maxFileSize = 10 << 30 // Default to 10GB if 0 (unlimited) is configured, to prevent abuse
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)
	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		http.Error(w, fmt.Sprintf("File too large: %v", err), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Use the configured SaveFolder
	saveFolder := s.config.SaveFolder
	if saveFolder == "" {
		// Fallback to user's downloads if SaveFolder is not configured
		homeDir, err := os.UserHomeDir()
		if err != nil {
			http.Error(w, "Could not determine user home directory", http.StatusInternalServerError)
			return
		}
		saveFolder = filepath.Join(homeDir, "Downloads")
	}
	os.MkdirAll(saveFolder, os.ModePerm) // Ensure it exists

	// Determine the destination path using rcfilepath to handle duplicates
	dstPath := rcfilepath(saveFolder, handler.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating file: %v", err), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file data
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, fmt.Sprintf("Error saving file: %v", err), http.StatusInternalServerError)
		return
	}

	// Add the file to the shared list
	if err := s.AddSharedFiles([]string{dstPath}); err != nil {
		http.Error(w, fmt.Sprintf("Error sharing file: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "File uploaded and shared successfully")
}

func (s *Service) handleLogin(w http.ResponseWriter, r *http.Request) {
	s.sessionsLock.Lock()
	defer s.sessionsLock.Unlock()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	passcode := r.FormValue("passcode")
	if passcode != s.passcode {
		http.Error(w, "Invalid passcode", http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.NewString()
	s.sessions = append(s.sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(12 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	w.WriteHeader(http.StatusOK)
}

func (s *Service) isValidSession(r *http.Request) bool {
	s.sessionsLock.RLock()
	defer s.sessionsLock.RUnlock()

	c, err := r.Cookie("session_token")
	if err != nil {
		return false
	}
	sessionToken := c.Value
	return slices.Contains(s.sessions, sessionToken)
}

func handleUIContent(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "index.html"
	}
	path = strings.TrimPrefix(path, "/")

	data, err := ui.ReadFile("ui/" + path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	contentType := "text/plain"
	if strings.HasSuffix(path, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(data)
}
