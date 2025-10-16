package filesend

import (
	"fmt"
	"io"
	"os"
	"path"

	"alat/pkg/core/device"
	"alat/pkg/core/pair"
	"alat/pkg/pbuf"

	"github.com/labstack/gommon/log"
)

type FileSendServer struct {
	pbuf.UnimplementedFileSendServiceServer
	Service     *Service
	PairManager *pair.PairManager
}

func rcfilepath(folder string, name string) string {
	newName := name
	ext := path.Ext(name)
	stem := name[:len(name)-len(ext)-1]
	for i := range 1000 {
		if i != 0 {
			newName = fmt.Sprintf("%s-%d.%s", stem, i, ext)
		}
		dest := path.Join(folder, newName)
		_, err := os.Stat(dest)
		if err != nil {
			return dest
		}
	}
	log.Errorf("Error: could not get file output path in downloads")
	return name
}

func (s *FileSendServer) SendFile(stream pbuf.FileSendService_SendFileServer) error {
	// The first message must be the initial request with metadata and sender info.
	req, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("failed to receive initial request: %w", err)
	}

	initialReq := req.GetInitialRequest()
	if initialReq == nil {
		return fmt.Errorf("protocol error: expected InitialSendFileRequest, got something else")
	}

	metadata := initialReq.GetMetadata()
	senderInfoPBUF := initialReq.GetSenderInfo()
	if metadata == nil || senderInfoPBUF == nil {
		return fmt.Errorf("protocol error: initial request is missing metadata or sender info")
	}

	senderInfo := device.PbufToInfo(senderInfoPBUF)

	// Create the file
	dest := rcfilepath(s.Service.config.SaveFolder, metadata.GetName())

	file, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, os.FileMode(metadata.Mode))
	fmt.Println("Saving file to: ", dest)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", metadata.Name, err)
	}
	defer file.Close()

	var transferredSize int64
	status := &FileTransferStatus{
		Filename:        metadata.Name,
		TotalSize:       metadata.Size,
		TransferredSize: 0,
		Status:          TransferStatusTransferring,
	}
	s.Service.UpdateIncomingStatus(senderInfo, status)

	// Process subsequent chunk messages
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break // End of stream
		}
		if err != nil {
			status.Status = TransferStatusFailed
			s.Service.UpdateIncomingStatus(senderInfo, status)
			return fmt.Errorf("error receiving chunk: %w", err)
		}

		chunk := req.GetChunk()

		if chunk == nil {
			// This could happen if the client sends an empty message. We'll just ignore it.
			continue
		}

		n, err := file.Write(chunk.Content)
		if err != nil {
			status.Status = TransferStatusFailed
			s.Service.UpdateIncomingStatus(senderInfo, status)
			return fmt.Errorf("failed to write chunk to file: %w", err)
		}
		transferredSize += int64(n)

		// Update status
		status.TransferredSize = transferredSize
		s.Service.UpdateIncomingStatus(senderInfo, status)
	}

	// Final status update
	status.Status = TransferStatusCompleted
	status.TransferredSize = status.TotalSize
	s.Service.UpdateIncomingStatus(senderInfo, status)

	return stream.SendAndClose(&pbuf.SendFileResponse{
		Status: pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_OK,
		Msg:    "File received successfully",
	})
}
