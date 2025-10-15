// Package node: the core's core
package node

import (
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
	storage     storage.NodeStorage
	pairManager *pair.PairManager
	discovery   *discovery.Manager
	services    *service.Registry
	server      *server.Server
	connected   *connected.Manager
}

func CreateNode(store storage.NodeStorage) (*Node, error) {
	serviceConfig, err := store.GetServiceConfig()
	if err != nil {
		return nil, err
	}
	registry := service.CreateRegistry(serviceConfig)
	appConfig, err := store.GetAppConfig()
	if err != nil {
		return nil, err
	}
	pairManager, err := pair.NewManager(store, &device.Details{
		Certificate: appConfig.Certificate,
		Color:       appConfig.DeviceColor,
		Name:        appConfig.DeviceName,
		Type:        appConfig.DeviceType,
	})
	if err != nil {
		return nil, err
	}

	discoveryManager, err := discovery.NewManager()
	if err != nil {
		return nil, err
	}

	return &Node{
		server:      server.NewServer(registry, pairManager),
		services:    registry,
		storage:     store,
		pairManager: pairManager,
		discovery:   discoveryManager,
		connected:   connected.NewManageer(pairManager, discoveryManager.Discoverer),
		workerState: workerState{},
	}, nil
}

func (n *Node) SetDetails(details *device.Details) {
	n.pairManager.SetDeviceDetails(details)
}
func (n *Node) GetDetails() *device.Details {
	return n.pairManager.GetDeviceDetails()
}

func (n *Node) GetPort() int {
	return n.server.Port
}

func (n *Node) Start() error {
	port, err := n.server.Start()
	if err != nil {
		return err
	}
	err = n.discovery.Server.Start(port)
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
func (n *Node) GetFoundDevices() []discovery.FoundDevice {
	return n.discovery.Discoverer.GetFoundDevices()
}
