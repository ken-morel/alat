package node

import (
	"alat/pkg/core/service"
)

func (n *Node) GetServices() *service.Registry {
	return n.services
}
