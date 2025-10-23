package webshare

import (
	"slices"

	"github.com/google/uuid"
)

func (s *Service) CreateSession() string {
	s.sessionsLock.Lock()
	defer s.sessionsLock.Unlock()
	sessionID := uuid.New().String()
	if len(s.sessions) > 50 {
		newSessions := make([]string, 50)
		starts := len(s.sessions) - 50
		for i := range 50 {
			newSessions = append(newSessions, s.sessions[starts+i])
		}
		s.sessions = newSessions
	}

	s.sessions = append(s.sessions, sessionID)
	return sessionID
}

func (s *Service) SessionExists(sessionID string) bool {
	s.sessionsLock.Lock()
	defer s.sessionsLock.Unlock()
	return slices.Contains(s.sessions, sessionID)
}
