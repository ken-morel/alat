// Package filesend: Holds file send service
package filesend

import (
	"alat/pkg/core/config"
	"alat/pkg/core/device"
	"alat/pkg/core/pair"
	"sync"
)

// TransferStatus defines the state of a file transfer.
type TransferStatus string

const (
	TransferStatusTransferring TransferStatus = "transferring"
	TransferStatusCompleted    TransferStatus = "completed"
	TransferStatusFailed       TransferStatus = "failed"
	TransferStatusPending      TransferStatus = "pending"
)

// FileTransferStatus holds the detailed progress of a single file transfer.
type FileTransferStatus struct {
	Filename        string
	TotalSize       int64
	TransferredSize int64
	Status          TransferStatus
}

// PeerTransferSession manages all transfers associated with a single peer.
type PeerTransferSession struct {
	PeerInfo          *device.Info
	IncomingTransfers map[string]*FileTransferStatus // Keyed by filename
	OutgoingTransfers map[string]*FileTransferStatus // Keyed by filename
	sync.RWMutex
}

// Service manages all file transfer sessions.
type Service struct {
	config        config.FileSendConfig
	ready         bool
	sessions      map[string]*PeerTransferSession // Keyed by peer ID
	sessionsMutex sync.RWMutex
	pairManager   *pair.PairManager
}

// getOrCreateSession retrieves an existing session or creates a new one for a peer.
func (s *Service) getOrCreateSession(peerInfo *device.Info) *PeerTransferSession {
	s.sessionsMutex.Lock()
	defer s.sessionsMutex.Unlock()

	if session, ok := s.sessions[peerInfo.ID]; ok {
		return session
	}

	session := &PeerTransferSession{
		PeerInfo:          peerInfo,
		IncomingTransfers: make(map[string]*FileTransferStatus),
		OutgoingTransfers: make(map[string]*FileTransferStatus),
	}
	s.sessions[peerInfo.ID] = session
	return session
}

func (s *Service) UpdateIncomingStatus(peerInfo *device.Info, status *FileTransferStatus) {
	session := s.getOrCreateSession(peerInfo)
	session.Lock()
	defer session.Unlock()
	session.IncomingTransfers[status.Filename] = status
}

func (s *Service) UpdateOutgoingStatus(peerInfo *device.Info, status *FileTransferStatus) {
	session := s.getOrCreateSession(peerInfo)
	session.Lock()
	defer session.Unlock()
	session.OutgoingTransfers[status.Filename] = status
}

func (s *Service) AddPendingTransfers(peerInfo *device.Info, files []string) {
	session := s.getOrCreateSession(peerInfo)
	for _, file := range files {
		if _, ok := session.OutgoingTransfers[file]; !ok {
			s.UpdateOutgoingStatus(peerInfo, &FileTransferStatus{
				Status:          TransferStatusPending,
				Filename:        file,
				TotalSize:       1,
				TransferredSize: 0,
			})
		}
	}
}

func (s *Service) GetSession(peerID string) (*PeerTransferSession, bool) {
	s.sessionsMutex.RLock()
	defer s.sessionsMutex.RUnlock()
	session, ok := s.sessions[peerID]
	return session, ok
}

func (s *Service) Enabled() bool {
	return s.config.Enabled
}

func (s *Service) Configure(c config.FileSendConfig) {
	s.config = c
}

func CreateService(conf config.FileSendConfig, p *pair.PairManager) Service {
	return Service{
		ready:       true,
		config:      conf,
		sessions:    make(map[string]*PeerTransferSession),
		pairManager: p,
	}
}
