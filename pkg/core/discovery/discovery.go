// Package discovery: handles mdns network discovery
package discovery

const (
	// DefaultPort is the default port for the alat service.
	DefaultPort = 12121
)

type Status int

const (
	Exposing Status = iota
	Stopped
)

type Manager struct {
	Discoverer *Discoverer
	Server     *Server
	Status     Status
}

func (m *Manager) Expose(port int) error {
	m.Status = Exposing
	m.Server.Start(port)
	return nil
}

func (m *Manager) Stop() {
	if m.Server != nil {
		m.Server.Stop()
	}
	if m.Discoverer != nil {
		m.Discoverer.Stop()
	}
	m.Status = Stopped
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
		Status:     Stopped,
	}, nil
}
