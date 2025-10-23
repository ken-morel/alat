package webshare

import (
	"slices"

	"github.com/google/uuid"
)

func (s *Service) CreateSession() string {
	s.sessionsLock.Lock()
	defer s.sessionsLock.Unlock()
	sessionID := uuid.New().String()

	if len(s.sessions) >= 50 {
		s.sessions = s.sessions[len(s.sessions)-49:] // Keep the last 49 sessions, plus the new one
	}

	s.sessions = append(s.sessions, sessionID)
	return sessionID
}

func (s *Service) SessionExists(sessionID string) bool {
	s.sessionsLock.RLock()
	defer s.sessionsLock.RUnlock()
	return slices.Contains(s.sessions, sessionID)
}
