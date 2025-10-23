// Package webshare provides the webshare service
package webshare

import (
	"alat/pkg/core/config"
	"alat/pkg/core/pair"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	DefaultPort = 8080
	MaxPort     = 8090
)

type SharedFile struct {
	UUID string `json:"uuid"`
	Path string `json:"path"`
	Name string `json:"name"`
	Size uint64 `json:"size"`
}
type Service struct {
	config          config.FileSendConfig
	sharedFiles     []SharedFile
	sharedFilesLock sync.RWMutex

	port     int
	passcode string
	running  bool

	runningLock sync.Mutex
	server      *http.Server
	listener    net.Listener

	sessionsLock sync.RWMutex
	sessions     []string

	pairManager *pair.PairManager
}

func CreateService(c *config.FileSendConfig, p *pair.PairManager) Service {
	return Service{
		config:      *c,
		passcode:    CreatePasscode(),
		pairManager: p,
		sharedFiles: []SharedFile{},
	}
}

func CreatePasscode() string {
	return uuid.NewString()[0:6]
}

// AddSharedFiles adds files to the list of shared files.
func (s *Service) AddSharedFiles(paths []string) error {
	s.sharedFilesLock.Lock()
	defer s.sharedFilesLock.Unlock()

	// Keep track of existing paths to avoid duplicates
	existingPaths := make(map[string]bool)
	for _, f := range s.sharedFiles {
		existingPaths[f.Path] = true
	}

	for _, p := range paths {
		if existingPaths[p] {
			continue // Skip duplicates
		}

		info, err := os.Stat(p)
		if err != nil {
			return fmt.Errorf("could not get file info for %s: %w", p, err)
		}
		if info.IsDir() {
			continue // Skip directories
		}
		file := SharedFile{
			UUID: uuid.NewString(),
			Path: p,
			Name: filepath.Base(p),
			Size: uint64(info.Size()),
		}
		s.sharedFiles = append(s.sharedFiles, file)
		existingPaths[p] = true
	}
	return nil
}

// RemoveSharedFile removes a file from the list of shared files by its UUID.
func (s *Service) RemoveSharedFile(uuid string) {
	s.sharedFilesLock.Lock()
	defer s.sharedFilesLock.Unlock()

	newFiles := []SharedFile{}
	for _, f := range s.sharedFiles {
		if f.UUID != uuid {
			newFiles = append(newFiles, f)
		}
	}
	s.sharedFiles = newFiles
}

// ClearSharedFiles removes all files from the list.
func (s *Service) ClearSharedFiles() {
	s.sharedFilesLock.Lock()
	defer s.sharedFilesLock.Unlock()
	s.sharedFiles = []SharedFile{}
}

// GetSharedFiles returns the list of shared files.
func (s *Service) GetSharedFiles() []SharedFile {
	s.sharedFilesLock.RLock()
	defer s.sharedFilesLock.RUnlock()
	// Return a copy to be safe
	filesCopy := make([]SharedFile, len(s.sharedFiles))
	copy(filesCopy, s.sharedFiles)
	return filesCopy
}

// GetPasscode returns the current passcode.
func (s *Service) GetPasscode() string {
	return s.passcode
}

// SetPasscode sets a new passcode.
func (s *Service) SetPasscode(passcode string) {
	s.passcode = passcode
}

// GetPort returns the port the server is running on.
func (s *Service) GetPort() int {
	if !s.IsRunning() {
		return 0
	}
	return s.port
}

// IsRunning returns true if the server is running.
func (s *Service) IsRunning() bool {
	s.runningLock.Lock()
	defer s.runningLock.Unlock()
	return s.running
}

func (s *Service) Start() (int, error) {
	s.runningLock.Lock()
	defer s.runningLock.Unlock()

	if s.running {
		return s.port, nil
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/shared-files", s.handleGetSharedFiles)
	mux.HandleFunc("/api/download", s.handleDownloadFile)
	mux.HandleFunc("/api/upload", s.handleUploadFile)
	mux.HandleFunc("/api/login", s.handleLogin)
	mux.HandleFunc("/", handleUIContent)

	var lis net.Listener
	var err error
	port := DefaultPort
	for {
		if port > MaxPort {
			return 0, fmt.Errorf("could not find a free port to start webshare server")
		}
		lis, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			break
		}
		port++
	}

	s.port = port
	s.listener = lis
	s.server = &http.Server{Handler: mux}
	s.running = true

	go func() {
		if err := s.server.Serve(s.listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Webshare server error: %v\n", err)
			s.runningLock.Lock()
			s.running = false
			s.runningLock.Unlock()
		}
	}()

	return s.port, nil
}

// Stop gracefully shuts down the webshare server.
func (s *Service) Stop() error {
	s.runningLock.Lock()
	defer s.runningLock.Unlock()

	if !s.running {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("webshare server shutdown failed: %w", err)
	}

	s.running = false
	s.port = 0
	s.server = nil
	s.listener = nil
	fmt.Println("Webshare server stopped.")
	return nil
}
