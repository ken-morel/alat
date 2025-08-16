package server

import (
	"alat/pkg/core/device"
	"alat/pkg/core/pbuf"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var Running bool = false

// PairingServer implements the pbuf.PairingServer interface.
type PairingServer struct {
	pbuf.UnimplementedPairingServer
}

func (s *PairingServer) RequestPair(ctx context.Context, req *pbuf.PairRequest) (*pbuf.PairResponse, error) {
	fmt.Printf("Received pairing request from device: %s\n", req.Device.Name)

	// TODO: Ask user authorization
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	token := fmt.Sprintf("%s-%s", req.Device.Id, uid)
	info, err := device.ThisDeviceInfo.ToPBuf()
	if err != nil {
		return nil, err
	}

	return &pbuf.PairResponse{
		Accepted: false,
		Token:    token,
		Device:   &info,
	}, nil
}

func startServer() {
	// TODO: Get port from config
	lis, err := net.Listen("tcp", ":60001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pbuf.RegisterPairingServer(s, &PairingServer{})

	fmt.Println("gRPC server listening on :60001")
	Running = true
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	Running = false
}

func Start() {
	go startServer()
}

