package discovery

import (
	"fmt"
	"os"

	"github.com/grandcat/zeroconf"
)

func (s *Server) Stop() {
	s.Running = false
	if s.zero != nil {
		s.zero.Shutdown()
	}
}

type Server struct {
	zero    *zeroconf.Server
	Running bool
}

func (s *Server) Start(port int) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	s.Running = true
	s.zero, err = zeroconf.Register(fmt.Sprintf("%s%d", hostname, port), "_alat._tcp", "local.", port, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		s.Running = false
		return fmt.Errorf("failed to register mdns server: %w", err)
	}
	return nil
}

func NewServer() *Server {
	return &Server{}
}
