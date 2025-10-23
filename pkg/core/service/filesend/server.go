package filesend

import (
	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/pbuf"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/labstack/gommon/log"
)

type FileSendServer struct {
	pbuf.UnimplementedFileSendServiceServer
	Service *Service
}

func rcfilepath(folder string, name string) string {
	newName := name
	ext := path.Ext(name)
	stem := name[:len(name)-len(ext)]
	for i := range 1_000_000 {
		if i != 0 {
			newName = fmt.Sprintf("%s-%d%s", stem, i, ext)
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
	if !s.Service.Enabled() {
		return fmt.Errorf("sending files to this device(%s) is disabled", s.Service.pairManager.GetDeviceDetails().Name)
	}
	req, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("failed to receive initial request: %w", err)
	}

	initialReq := req.GetInitialRequest()
	if initialReq == nil {
		return fmt.Errorf("protocol error: expected InitialSendFileRequest, got something else")
	}

	token := security.PairToken(initialReq.GetToken())
	if !s.Service.pairManager.IsTokenValid(token) {
		return fmt.Errorf("file sending unauthorized(device is not authorized to send files to %s), device received invalid pair token", s.Service.pairManager.GetDeviceDetails().Name)
	}

	metadata := initialReq.GetMetadata()
	senderInfoPBUF := initialReq.GetSenderInfo()
	if metadata == nil || senderInfoPBUF == nil {
		return fmt.Errorf("protocol error: initial request is missing metadata or sender info")
	}

	senderInfo := device.PbufToInfo(senderInfoPBUF)

	dest := rcfilepath(s.Service.config.SaveFolder, metadata.GetName())

	file, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, os.FileMode(metadata.Mode))
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

		status.TransferredSize = transferredSize
		s.Service.UpdateIncomingStatus(senderInfo, status)
	}

	status.Status = TransferStatusCompleted
	status.TransferredSize = status.TotalSize
	s.Service.UpdateIncomingStatus(senderInfo, status)

	return stream.SendAndClose(&pbuf.SendFileResponse{
		Status: pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_OK,
		Msg:    "File received successfully",
	})
}
