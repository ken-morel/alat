// Package discovery: handles mdns network discovery
package discovery

type Status int

const (
	Exposing Status = iota
	Stopped
)

type Manager struct {
	Discoverer *Discoverer
	Server     *Server
}

func (m *Manager) Expose(port int) error {
	return m.Server.Start(port)
}

func (m *Manager) Stop() {
	if m.Server != nil {
		m.Server.Stop()
	}
}

func NewManager() (*Manager, error) {
	discoverer := NewDiscoverer()

	server := NewServer()

	return &Manager{
		Discoverer: discoverer,
		Server:     server,
	}, nil
}
