// Package node: the core's core
package node

import (
	"alat/pkg/core"
	"alat/pkg/core/connected"
	"alat/pkg/core/device"
	"alat/pkg/core/discovery"
	"alat/pkg/core/pair"
	"alat/pkg/core/service"
	"alat/pkg/core/storage"
	"alat/pkg/core/transport/server"
)

type Node struct {
	workerState workerState
	Storage     *storage.NodeStorage
	PairManager *pair.PairManager
	discovery   *discovery.Manager
	device      *device.Details
	services    *service.Registry
	server      *server.Server
	Connected   *connected.Manager
}

func NewNode(registry *service.Registry, store *storage.NodeStorage, details *device.Details, manager *pair.PairManager) (*Node, error) {
	server := server.NewServer(registry, manager)
	discoveryManager, err := discovery.NewManager()
	if err != nil {
		return nil, err
	}

	return &Node{
		server:      server,
		device:      details,
		services:    registry,
		Storage:     store,
		PairManager: manager,
		discovery:   discoveryManager,
		Connected:   connected.NewManageer(manager, discoveryManager.Discoverer),
		workerState: workerState{},
	}, nil
}

func (n *Node) GetDiscoverer() *discovery.Discoverer {
	return n.discovery.Discoverer
}

func (n *Node) SetDetails(details *device.Details) {
	n.device = details
	n.PairManager.SetDetails(details)
}

func (n *Node) Start() error {
	err := n.discovery.Server.Start(core.AlatPort)
	if err != nil {
		return err
	}
	err = n.server.Start()
	if err != nil {
		return err
	}
	n.StartWorker()
	return nil
}

func (n *Node) Stop() {
	n.server.Stop()
	n.discovery.Stop()
	n.StopWorker()
}
