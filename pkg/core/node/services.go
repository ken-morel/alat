package node

import (
	"alat/pkg/core/service"
	"alat/pkg/core/service/filesend"
)

func (n *Node) GetFileTransfersStatus() *filesend.FileTransfersStatus {
	return n.services.FileSend.GetStatus()
}

func (n *Node) GetServices() *service.Registry {
	return n.services
}
