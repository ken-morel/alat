// Package node: the core's core
package node

import (
	"alat/pkg/core/device"
	"alat/pkg/core/discovery"
	"alat/pkg/core/pair"
	"alat/pkg/core/service"
	"alat/pkg/core/storage"
	"alat/pkg/core/transport"
)

type Node struct {
	Storage     *storage.NodeStorage
	PairManager *pair.PairManager
	discovery   *discovery.Manager
	device      *device.Details
	services    *service.Registry
	server      *transport.Server
}

func NewNode(registry *service.Registry, store *storage.NodeStorage, details *device.Details) (*Node, error) {
	mamager, err := pair.NewManager(store, details)
	if err != nil {
		return nil, err
	}
	server := transport.NewServer(registry, mamager)
	discoveryManager, err := discovery.NewManager()
	if err != nil {
		return nil, err
	}

	return &Node{
		server:      server,
		device:      details,
		services:    registry,
		Storage:     store,
		PairManager: mamager,
		discovery:   discoveryManager,
	}, nil
}

func (n *Node) SetDetails(details *device.Details) {
	n.device = details
	n.PairManager.SetDetails(details)
}
