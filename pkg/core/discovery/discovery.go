// Package discovery: handles mdns network discovery
package discovery

import "fmt"

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
		if err := m.Server.Start(port); err != nil {
			// Don't return the error, as this may not be a fatal error.
			// For example, on mobile, the Go mdns server may fail to start,
			// but the Dart mdns server will be used instead.
			fmt.Printf("discovery server failed to start: %v\n", err)
		}
	}
	return nil
}

func (m *Manager) Stop() {
	if m.Server != nil {
		go m.Server.Stop()
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
