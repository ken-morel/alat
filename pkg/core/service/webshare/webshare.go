// Package webshare provides the webshare service
package webshare

import (
	"net/http"
	"sync"

	"alat/pkg/core/config"
)

type SharedFile struct {
	Path string `json:"path"`
	Size uint64 `json:"size"`
}
type Service struct {
	config          config.FileSendConfig
	sharedFiles     map[string]SharedFile
	sharedFilesLock sync.Mutex
	passcode        string

	shouldStop   bool
	port         int
	runningLock  sync.Mutex
	sessions     []string
	password     string
	sessionsLock sync.RWMutex
}

func (s *Service) StartServer() {
	s.runningLock.Lock()
	defer s.runningLock.Unlock()
	if s.port != 0 {
		return
	}

	server := http.NewServeMux()
	server.HandleFunc("api/sharedfiles", s.handleGetSharedFiles)
	server.HandleFunc("api/download", s.handleDownloadFile)
	server.HandleFunc("api/login", s.handleLogin)
	server.HandleFunc("", handleUIContent)
}
