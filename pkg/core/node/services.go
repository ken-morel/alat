package node

import "alat/pkg/core/service/filesend"

func (n *Node) GetFileTransfersStatus() *filesend.FileTransfersStatus {
	return n.services.FileSend.GetStatus()
}
