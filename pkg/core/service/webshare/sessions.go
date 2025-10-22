package webshare

import "github.com/google/uuid"

func (s *Service) CreateSession() string {
	s.sessionsLock.Lock()
	defer s.sessionsLock.Unlock()
	sessionId := uuid.New().String()
	s.sessions = append(s.sessions, sessionId)
	return sessionId
}

func (s *Service) SessionExists(sessionId string) bool {
	s.sessionsLock.Lock()
	defer s.sessionsLock.Unlock()
	for _, session := range s.sessions {
		if session == sessionId {
			return true
		}
	}

	return false
}
