package node

import (
	"alat/pkg/core/connected"
	"alat/pkg/core/service/filesend"
	"context"
	"fmt"
)

func (n *Node) QuerySendFiles(p *connected.Connected, files []string) <-chan error {
	fmt.Println("Sending ", len(files), " files")
	channel := make(chan error)
	ctx := context.TODO()
	go func() {
		fmt.Println("In goroutine, adding pending transfers")
		n.services.FileSend.AddPendingTransfers(&p.Info, files)
		for _, file := range files {
			fmt.Print("Transfering ..", file)
			channel <- n.services.FileSend.SendFile(ctx, p.IP, p.Port, &p.PairedDevice.Token, file)
			fmt.Println("  Done")
		}
		fmt.Println("Done sending files, closing channel")
		close(channel)
	}()
	return channel
}

func (n *Node) GetFileTransfersStatus() *filesend.FileTransfersStatus {
	return n.services.FileSend.GetStatus()
}
