// Package discovery: handles mdns network discovery
package discovery

type Discoverer struct{}

type Server struct{}

type Manager struct {
	Discoverer Discoverer
	Server     Server
}
