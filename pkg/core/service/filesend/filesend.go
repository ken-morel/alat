// Package filesend: Holds file send service
package filesend

import (
	"sync"

	"alat/pkg/pbuf"
)

type Config struct {
	Enabled     bool
	SaveFolder  string
	FileMaxSize uint32
}

// TransferStatus defines the state of a file transfer.
type TransferStatus string

const (
	TransferStatusTransferring TransferStatus = "transferring"
	TransferStatusCompleted    TransferStatus = "completed"
	TransferStatusFailed       TransferStatus = "failed"
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
	PeerInfo          *pbuf.DeviceInfo
	IncomingTransfers map[string]*FileTransferStatus // Keyed by filename
	OutgoingTransfers map[string]*FileTransferStatus // Keyed by filename
	sync.RWMutex
}

// Service manages all file transfer sessions.
type Service struct {
	config        Config
	ready         bool
	sessions      map[string]*PeerTransferSession // Keyed by peer ID
	sessionsMutex sync.RWMutex
}

// getOrCreateSession retrieves an existing session or creates a new one for a peer.
func (s *Service) getOrCreateSession(peerInfo *pbuf.DeviceInfo) *PeerTransferSession {
	s.sessionsMutex.Lock()
	defer s.sessionsMutex.Unlock()

	if session, ok := s.sessions[peerInfo.Id]; ok {
		return session
	}

	session := &PeerTransferSession{
		PeerInfo:          peerInfo,
		IncomingTransfers: make(map[string]*FileTransferStatus),
		OutgoingTransfers: make(map[string]*FileTransferStatus),
	}
	s.sessions[peerInfo.Id] = session
	return session
}

// UpdateIncomingStatus updates the status of an incoming file transfer.
func (s *Service) UpdateIncomingStatus(peerInfo *pbuf.DeviceInfo, status *FileTransferStatus) {
	session := s.getOrCreateSession(peerInfo)
	session.Lock()
	defer session.Unlock()
	session.IncomingTransfers[status.Filename] = status
}

// GetSession returns the transfer session for a given peer ID.
func (s *Service) GetSession(peerID string) (*PeerTransferSession, bool) {
	s.sessionsMutex.RLock()
	defer s.sessionsMutex.RUnlock()
	session, ok := s.sessions[peerID]
	return session, ok
}

func (s *Service) Enabled() bool {
	return s.config.Enabled
}

func (s *Service) Configure(c Config) {
	s.config = c
}

func CreateService(conf Config) Service {
	return Service{
		ready:    true,
		config:   conf,
		sessions: make(map[string]*PeerTransferSession),
	}
}
