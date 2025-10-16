package server

import (
	"context"

	"alat/pkg/pbuf"
)

func (s *Server) GetDetails(ctx context.Context, req *pbuf.GetDetailsRequest) (*pbuf.GetDetailsResponse, error) {
	return &pbuf.GetDetailsResponse{
		Details: s.PairManager.GetDeviceDetails().ToPBUF(),
	}, nil
}

func (s *Server) GetInfo(ctx context.Context, req *pbuf.GetInfoRequest) (*pbuf.GetInfoResponse, error) {
	return &pbuf.GetInfoResponse{
		Info: s.PairManager.GetDeviceDetails().GetInfo().ToPBUF(),
	}, nil
}
