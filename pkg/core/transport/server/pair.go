package server

import (
	"context"

	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/pbuf"
)

func (s *Server) RequestPair(ctx context.Context, req *pbuf.RequestPairRequest) (*pbuf.RequestPairResponse, error) {
	accepted, reason := s.PairManager.HandlePairRequest((*security.PairToken)(req.GetToken()), device.PbufToDetails(req.GetDetails()))
	return &pbuf.RequestPairResponse{
		Token:    req.GetToken(),
		Accepted: accepted,
		Reason:   reason,
		Details:  s.PairManager.GetDeviceDetails().ToPBUF(),
	}, nil
}
