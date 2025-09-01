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
	m.Server.Start(port)
	return nil
}

func (m *Manager) Stop() {
	if m.Server != nil {
		m.Server.Stop()
	}
}

func NewManager() (*Manager, error) {
	discoverer, err := NewDiscoverer()
	if err != nil {
		return nil, err
	}
	server := NewServer()
	return &Manager{
		Discoverer: discoverer,
		Server:     server,
	}, nil
}
