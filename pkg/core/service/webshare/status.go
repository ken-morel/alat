package webshare

import "fmt"

type Status struct {
	IsRunning   bool         `json:"isRunning"`
	Port        int          `json:"port"`
	Passcode    string       `json:"passcode"`
	SharedFiles []SharedFile `json:"sharedFiles"`
	ShareURL    string       `json:"shareURL"`
}

func (s *Service) GetStatus() *Status {
	port := s.GetPort()
	url := ""
	if port != 0 {
		// This should ideally resolve the local IP, but localhost is fine for development.
		url = fmt.Sprintf("http://localhost:%d", port)
	}

	return &Status{
		IsRunning:   s.IsRunning(),
		Port:        port,
		Passcode:    s.GetPasscode(),
		SharedFiles: s.GetSharedFiles(),
		ShareURL:    url,
	}
}
