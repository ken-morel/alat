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

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"files":  nil,
			"status": 400, // http.StatusBadRequest
		})

		return
	}
	if !s.SessionExists(r.FormValue("session")) {
		json.NewEncoder(w).Encode(map[string]any{
			"files":  s.sharedFiles,
			"status": 401, // http.StatusUnauthorized
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"status": 200, // http.StatusOK
		"files":  s.sharedFiles,
	})
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

	file, exists := s.sharedFiles[r.FormValue("file")]
	if exists {
		http.ServeFile(w, r, file.Path)
	} else {
		http.Error(w, "File does not exist", http.StatusNotFound)
	}
}

func (s *Service) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"status":  400, // http.StatusBadRequest,
			"session": "",
		})
		return
	}
	passcode := r.FormValue("passcode")
	if passcode != s.passcode {
		http.Error(w, "Invalid passcode", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  406, // http.StatusNotAcceptable,
			"session": "",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"status":  200, // http.StatusOK,
		"session": s.CreateSession(),
	})
}

func handleUIContent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	path, _ := strings.CutPrefix(r.URL.Path, "/")
	if path == "" {
		path = "index.html"
	}
	http.ServeFileFS(w, r, ui, path)
}
