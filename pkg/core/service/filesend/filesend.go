// Package filesend: Holds file send service
package filesend

import "sync"

type Config struct {
	Enabled bool
}

type FileTransferStatus struct {
	Filename        string
	TotalSize       int64
	TransferredSize int64
	Status          string // e.g., "transferring", "completed", "failed"
}

type Service struct {
	config         Config
	ready          bool
	statuses       map[string][]*FileTransferStatus // Keyed by peer ID
	statusesMutex  sync.RWMutex
}

func (s *Service) Enabled() bool {
	return s.config.Enabled
}

func (s *Service) Configure(c Config) {
	s.config = c
}

func (s *Service) GetStatus(peerID string) []*FileTransferStatus {
	s.statusesMutex.RLock()
	defer s.statusesMutex.RUnlock()
	return s.statuses[peerID]
}

func (s *Service) updateStatus(peerID string, status *FileTransferStatus) {
	s.statusesMutex.Lock()
	defer s.statusesMutex.Unlock()

	if _, ok := s.statuses[peerID]; !ok {
		s.statuses[peerID] = []*FileTransferStatus{}
	}

	// If a transfer for the same file exists, update it. Otherwise, add it.
	for i, existingStatus := range s.statuses[peerID] {
		if existingStatus.Filename == status.Filename {
			s.statuses[peerID][i] = status
			return
		}
	}
	s.statuses[peerID] = append(s.statuses[peerID], status)
}


func CreateService(conf Config) Service {
	return Service{
		ready:    true,
		config:   conf,
		statuses: make(map[string][]*FileTransferStatus),
	}
}
