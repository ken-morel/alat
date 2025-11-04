// Package discovery: handles mdns network discovery
package discovery

import (
	"fmt"
	"net"
)

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
	// --- Start Diagnostic Code ---
	var e string
	e += "--- Network Interface Diagnostics ---"
	ifaces, err := net.Interfaces()
	if err != nil {
		e += fmt.Sprintf("Error getting network interfaces: %v\n", err)
	} else {
		e += fmt.Sprintf("Found %d interfaces:\n", len(ifaces))
		for _, i := range ifaces {
			addrs, err := i.Addrs()
			if err != nil {
				e += fmt.Sprintf("  Interface %s: Error getting addresses: %v\n", i.Name, err)
			} else {
				e += fmt.Sprintf("  Interface: %s, Flags: %s, MTU: %d, Addrs: %v\n", i.Name, i.Flags.String(), i.MTU, addrs)
			}
		}
	}
	e += fmt.Sprintln("------------------------------------")
	// --- End Diagnostic Code ---
	//
	fmt.Println(e)

	discoverer, err := NewDiscoverer()
	if err != nil {
		return nil, fmt.Errorf("%v;\n;%v", e, err)
	}
	server := NewServer()
	return &Manager{
		Discoverer: discoverer,
		Server:     server,
	}, nil
}
