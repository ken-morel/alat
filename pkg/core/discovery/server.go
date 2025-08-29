package discovery

import (
	"fmt"
	"os"

	"github.com/grandcat/zeroconf"
)

func (s *Server) Stop() {
	if s.zero != nil {
		s.zero.Shutdown()
	}
}

type Server struct {
	zero *zeroconf.Server
}

func (s *Server) Start(port int) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	s.zero, err = zeroconf.Register(hostname, "_alat._tcp", "local.", port, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		return fmt.Errorf("failed to register mdns server: %w", err)
	}
	return nil
}

func NewServer() *Server {
	return &Server{}
}
