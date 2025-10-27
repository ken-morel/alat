package webshare

import "fmt"

type Status struct {
	Running     bool         `json:"running"`
	Port        int          `json:"port"`
	Passcode    string       `json:"passcode"`
	SharedFiles []SharedFile `json:"sharedFiles"`
	ShareURLs   []string     `json:"shareURLs"`
}

func (s *Service) GetStatus() *Status {
	port := s.GetPort()
	url := ""
	if port != 0 {
		// This should ideally resolve the local IP, but localhost is fine for development.
		url = fmt.Sprintf("http://localhost:%d", port)
	}

	return &Status{
		Running:     s.IsRunning(),
		Port:        port,
		Passcode:    s.GetPasscode(),
		SharedFiles: s.GetSharedFiles(),
		ShareURLs:   []string{url},
	}
}
