package filesend

import (
	"alat/pkg/core/security"
	"alat/pkg/pbuf"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// SendFile connects to a peer and sends a file.
// This is a self-contained method that handles the entire client-side lifecycle.
func (s *Service) SendFile(ctx context.Context, ip net.IP, port int, token *security.PairToken, filePath string) error {
	fullAddress := net.JoinHostPort(ip.To4().String(), strconv.Itoa(port))

	conn, err := grpc.NewClient(fullAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to peer: %w", err)
	}
	defer conn.Close()

	fileSendClient := pbuf.NewFileSendServiceClient(conn)

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	stream, err := fileSendClient.SendFile(ctx)
	if err != nil {
		return fmt.Errorf("failed to create send stream: %w", err)
	}

	senderInfo := s.pairManager.GetDeviceDetails().GetInfo()
	senderInfoPBUF := senderInfo.ToPBUF()
	initialReq := &pbuf.SendFileRequest{
		Data: &pbuf.SendFileRequest_InitialRequest{
			InitialRequest: &pbuf.InitialSendFileRequest{
				Metadata: &pbuf.FileMetadata{
					Name: fileInfo.Name(),
					Size: fileInfo.Size(),
					Mode: int32(fileInfo.Mode().Perm()),
				},
				SenderInfo: senderInfoPBUF,
				Token:      token[:],
			},
		},
	}
	if err := stream.Send(initialReq); err != nil {
		return fmt.Errorf("failed to send initial request: %w", err)
	}

	// 4. Stream file chunks and update status
	buffer := make([]byte, 1024*1024) // 1MB chunks
	var transferredSize int64
	status := &FileTransferStatus{
		Filename:        filePath,
		TotalSize:       fileInfo.Size(),
		TransferredSize: 0,
		Status:          TransferStatusTransferring,
	}
	s.UpdateOutgoingStatus(senderInfo, status)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			status.Status = TransferStatusFailed
			s.UpdateOutgoingStatus(senderInfo, status)
			return fmt.Errorf("failed to read file chunk: %w", err)
		}

		chunkReq := &pbuf.SendFileRequest{
			Data: &pbuf.SendFileRequest_Chunk{
				Chunk: &pbuf.FileChunk{
					Content: buffer[:n],
				},
			},
		}
		if err := stream.Send(chunkReq); err != nil {
			status.Status = TransferStatusFailed
			s.UpdateOutgoingStatus(senderInfo, status)
			return fmt.Errorf("failed to send chunk: %w", err)
		}

		transferredSize += int64(n)
		status.TransferredSize = transferredSize
		s.UpdateOutgoingStatus(senderInfo, status)
	}

	// 5. Close stream and get response
	resp, err := stream.CloseAndRecv()
	if err != nil {
		status.Status = TransferStatusFailed
		s.UpdateOutgoingStatus(senderInfo, status)
		return fmt.Errorf("failed to receive response: %w", err)
	}

	if resp.Status != pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_OK {
		status.Status = TransferStatusFailed
		s.UpdateOutgoingStatus(senderInfo, status)
		return fmt.Errorf("file transfer failed on server: %s", resp.Msg)
	}

	status.Status = TransferStatusCompleted
	s.UpdateOutgoingStatus(senderInfo, status)
	return nil
}
