package server

import (
	"alat/pkg/core/service/filesend"
	"alat/pkg/pbuf"
	"io"
	"os"
)

type FileSendServer struct {
	pbuf.UnimplementedFileSendServiceServer
	Service *filesend.Service
}

func (s *FileSendServer) SendFile(stream pbuf.FileSendService_SendFileServer) error {
	// The first message should be metadata
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	metadata := req.GetMetadata()
	if metadata == nil {
		return io.ErrUnexpectedEOF // Or a more specific error
	}

	// Create the file
	file, err := os.OpenFile(metadata.Name, os.O_CREATE|os.O_WRONLY, os.FileMode(metadata.Mode))
	if err != nil {
		return err
	}
	defer file.Close()

	var transferredSize int64

	// Update status
	status := &filesend.FileTransferStatus{
		Filename:        metadata.Name,
		TotalSize:       metadata.Size,
		TransferredSize: 0,
		Status:          "transferring",
	}
	s.Service.updateStatus("peerID", status) // TODO: Get peerID from context

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			status.Status = "failed"
			s.Service.updateStatus("peerID", status)
			return err
		}

		chunk := req.GetChunk()
		if chunk == nil {
			// Handle error: expected a chunk
			continue
		}

		n, err := file.Write(chunk.Content)
		if err != nil {
			status.Status = "failed"
			s.Service.updateStatus("peerID", status)
			return err
		}
		transferredSize += int64(n)

		// Update status
		status.TransferredSize = transferredSize
		s.Service.updateStatus("peerID", status)
	}

	status.Status = "completed"
	s.Service.updateStatus("peerID", status)

	return stream.SendAndClose(&pbuf.SendFileResponse{
		Status: pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_OK,
		Msg:    "File received successfully",
	})
}