// Package server: stores server handles
package server

import (
	"fmt"
	"net"

	"alat/pkg/core/service/filesend"

	"alat/pkg/core"
	"alat/pkg/core/pair"
	"alat/pkg/core/service"
	"alat/pkg/pbuf"

	"google.golang.org/grpc"
)

type Server struct {
	pbuf.UnimplementedAlatServiceServer
	Services    *service.Registry
	PairManager *pair.PairManager
	grpcServer  *grpc.Server
	listener    net.Listener
	Running     bool
	Port        int
}

func NewServer(registry *service.Registry, manager *pair.PairManager) *Server {
	return &Server{
		Services:    registry,
		PairManager: manager,
	}
}

func (s *Server) Start() (int, error) {
	var lis net.Listener
	var err error

	for s.Port = core.DefaultPort; s.Port < core.MaxPort; s.Port++ {
		lis, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.Port))
		if err == nil {
			break
		}
	}
	if err != nil {
		s.Port = 0
		return 0, fmt.Errorf("failed to listen, all alat ports used: %v", err)
	}
	s.listener = lis
	s.grpcServer = grpc.NewServer()
	pbuf.RegisterAlatServiceServer(s.grpcServer, s)
	pbuf.RegisterFileSendServiceServer(s.grpcServer, &filesend.FileSendServer{Service: &s.Services.FileSend})

	fmt.Printf("Server listening at %v\n", lis.Addr())
	go func() {
		s.Running = true
		if err := s.grpcServer.Serve(lis); err != nil {
			s.Running = false
			fmt.Printf("gRPC server error: %v\n", err)
		}
	}()

	return s.Port, nil
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
