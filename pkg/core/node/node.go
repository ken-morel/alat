// Package node: the core's core
package node

import (
	"fmt"

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
	PairManager *pair.PairManager
	discovery   *discovery.Manager
	Services    *service.Registry
	server      *server.Server
	connected   *connected.Manager
}

func CreateNode(store storage.NodeStorage) (*Node, error) {
	serviceConfig, err := store.GetServiceConfig()
	if err != nil {
		return nil, fmt.Errorf("error getting initial service configuraiton: %v", err)
	}
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

	registry := service.CreateRegistry(serviceConfig, pairManager)
	if err != nil {
		return nil, fmt.Errorf("error creating pair manager: %v", err)
	}
	discoveryManager, err := discovery.NewManager()
	if err != nil {
		return nil, fmt.Errorf("error creating discovery manager: %v", err)
	}

	return &Node{
		server:      server.NewServer(registry, pairManager),
		Services:    registry,
		storage:     store,
		PairManager: pairManager,
		discovery:   discoveryManager,
		connected:   connected.NewManageer(pairManager, discoveryManager.Discoverer),
		workerState: workerState{},
	}, nil
}

func (n *Node) SetDetails(details *device.Details) {
	n.PairManager.SetDeviceDetails(details)
}

func (n *Node) GetDetails() *device.Details {
	return n.PairManager.GetDeviceDetails()
}

func (n *Node) GetPort() int {
	return n.server.Port
}

func (n *Node) Start() error {
	status := n.GetStatus()
	listeningPort := status.Port
	if !status.ServerRunning {
		port, err := n.server.Start()
		if err != nil {
			return err
		}
		listeningPort = port

	}
	if !status.DiscoveryRunning {
		err := n.discovery.Server.Start(listeningPort)
		if err != nil {
			return err
		}
	}
	if !status.WorkerRunning {
		n.StartWorker()
	}
	return nil
}

func (n *Node) Stop() {
	status := n.GetStatus()
	if status.ServerRunning {
		n.server.Stop()
	}
	if status.DiscoveryRunning {
		n.discovery.Stop()
	}
	if status.WorkerRunning {
		n.StopWorker()
	}
}

func (n *Node) GetFoundDevices() []discovery.FoundDevice {
	return n.discovery.Discoverer.GetFoundDevices()
}

func (n *Node) GetConnectedDeviceByID(id string) *connected.Connected {
	for _, device := range n.GetConnectedDevices() {
		if device.Info.ID == id {
			return &device
		}
	}
	return nil
}

func (n *Node) SetFoundDevices(devices []discovery.FoundDevice) {
	n.discovery.Discoverer.ProvideFoundDevices(devices)
}

func (n *Node) DisableDiscovery() {
	n.discovery.Discoverer.Disable()
}
func (n *Node) DiscoveryEnabled() bool {
	return n.discovery.Discoverer.IsEnabled()
}
