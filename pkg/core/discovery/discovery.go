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
	disabled   bool
}

func (m *Manager) Expose(port int) error {
	if m.IsEnabled() {
		return m.Server.Start(port)
	}
	return nil
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

func (m *Manager) Disable() {
	m.Stop()
	m.disabled = true
	m.Discoverer.resolver = nil
}
func (m *Manager) IsEnabled() bool {
	return !m.disabled
}
