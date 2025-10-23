package node

import (
	"alat/pkg/core/connected"
	"alat/pkg/core/service/filesend"
	"context"
)

func (n *Node) QuerySendFiles(p *connected.Connected, files []string) <-chan error {
	channel := make(chan error)
	ctx := context.TODO()
	go func() {
		n.services.FileSend.AddPendingTransfers(&p.Info, files)
		for _, file := range files {
			channel <- n.services.FileSend.SendFile(ctx, p.IP, p.Port, &p.PairedDevice.Token, file)
		}
		close(channel)
	}()
	return channel
}

func (n *Node) GetFileTransfersStatus() *filesend.FileTransfersStatus {
	return n.services.FileSend.GetStatus()
}
