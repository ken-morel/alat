// Package transport: holds server and client methods
package transport

import (
	"alat/pkg/core/discovery"
	"alat/pkg/core/pair"
	"alat/pkg/core/service"
	"alat/pkg/pbuf"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pbuf.UnimplementedAlatServiceServer
	Services    *service.Registry
	PairManager *pair.PairManager
	grpcServer  *grpc.Server
	listener    net.Listener
	Running     bool
}

func NewServer(registry *service.Registry, manager *pair.PairManager) *Server {
	return &Server{
		Services:    registry,
		PairManager: manager,
	}
}

func (s *Server) RequestPair(ctx context.Context, req *pbuf.RequestPairRequest) (*pbuf.RequestPairResponse, error) {
	fmt.Println("Pair request received: ", req)
	// TODO: Insert some magick here
	return &pbuf.RequestPairResponse{
		Token:       req.GetToken(),
		Accepted:    false,
		Info:        s.PairManager.DeviceDetails().GetInfo().ToPBUF(),
		Certificate: s.PairManager.DeviceDetails().Certificate[:],
	}, nil
}

func (s *Server) GetDetails(ctx context.Context, req *pbuf.GetDetailsRequest) (*pbuf.GetDetailsResponse, error) {
	return &pbuf.GetDetailsResponse{
		Details: s.PairManager.DeviceDetails().ToPBUF(),
	}, nil
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", discovery.DefaultPort))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	s.listener = lis
	s.grpcServer = grpc.NewServer()
	pbuf.RegisterAlatServiceServer(s.grpcServer, s)
	fmt.Printf("Server listening at %v\n", lis.Addr())
	go func() {
		s.Running = true
		if err := s.grpcServer.Serve(lis); err != nil {
			s.Running = false
			fmt.Printf("gRPC server error: %v\n", err)
		}
	}()

	return nil
}

func (s *Server) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
		fmt.Println("gRPC server stopped.")
	}
	if s.listener != nil {
		s.listener.Close()
	}
	s.Running = false
}
