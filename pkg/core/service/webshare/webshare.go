// Package webshare provides the webshare service
package webshare

import (
	"alat/pkg/core"
	"alat/pkg/core/config"
	"alat/pkg/core/pair"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

const (
	DefaultPort = 80
	MaxPort     = core.MaxPort
)

type SharedFile struct {
	Path string `json:"path"`
	Size uint64 `json:"size"`
}
type Service struct {
	config          config.FileSendConfig
	sharedFiles     map[string]SharedFile
	sharedFilesLock sync.Mutex

	port        int
	runningLock sync.Mutex
	passcode    string

	sessionsLock sync.Mutex
	sessions     []string

	pairManager *pair.PairManager
}

func CreateService(c *config.FileSendConfig, p *pair.PairManager) Service {
	return Service{
		config:      *c,
		passcode:    CreatePasscode(),
		pairManager: p,
	}
}

func CreatePasscode() string {
	return uuid.NewString()[0:6]
}

func (s *Service) Start() (int, error) {
	s.runningLock.Lock()
	defer s.runningLock.Unlock()
	if s.port != 0 {
		return s.port, nil
	}

	server := http.NewServeMux()
	server.HandleFunc("api/sharedfiles", s.handleGetSharedFiles)
	server.HandleFunc("api/download", s.handleDownloadFile)
	server.HandleFunc("api/login", s.handleLogin)
	server.HandleFunc("", handleUIContent)

	var lis net.Listener
	var err error
	for s.port = DefaultPort; ; s.port += 1 {
		if s.port > MaxPort {
			return 0, fmt.Errorf("could not find usable port to create server")
		}
		lis, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.port))
		if err == nil {
			break
		}
	}
	go http.Serve(lis, server)
	return s.port, nil
}
