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
	Storage     storage.NodeStorage
	PairManager *pair.PairManager
	discovery   *discovery.Manager
	device      *device.Details
	services    *service.Registry
	server      *transport.Server
}

func (n *Node) SetDetails(details *device.Details) {
	n.device = details
	n.PairManager.SetInfo(details.GetInfo())
}
