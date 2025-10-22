package webshare

import (
	"embed"
	"encoding/json"
	"net/http"
	"strings"
)

//go:embed all:ui
var ui embed.FS

func (s *Service) handleGetSharedFiles(w http.ResponseWriter, r *http.Request) {
	s.sharedFilesLock.Lock()
	defer s.sharedFilesLock.Unlock()
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form "+err.Error(), http.StatusBadRequest)
	}
	if !s.SessionExists(r.FormValue("session")) {
		http.Error(w, "Unauthorized, invalid session", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(s.sharedFiles)
}

func (s *Service) handleDownloadFile(w http.ResponseWriter, r *http.Request) {
	s.sharedFilesLock.Lock()
	defer s.sharedFilesLock.Unlock()
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !s.SessionExists(r.FormValue("session")) {
		http.Error(w, "Unauthorized, invalid session", http.StatusUnauthorized)
		return
	}

	fileUUid := r.FormValue("file")
	file, exists := s.sharedFiles[fileUUid]
	if exists {
		http.ServeFile(w, r, file.Path)
	} else {
		http.Error(w, "File does not exist", http.StatusNotFound)
	}
}

func (s *Service) handleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Could not parse form: "+err.Error(), http.StatusBadRequest)
		return
	}
	passcode := r.FormValue("passcode")
	if passcode != s.passcode {
		http.Error(w, "Invalid passcode", http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(s.CreateSession()))
}

func handleUIContent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	path, _ := strings.CutPrefix(r.URL.Path, "/")
	http.ServeFileFS(w, r, ui, path)
}
