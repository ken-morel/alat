package server

import (
	"context"

	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/pbuf"
)

func (s *Server) RequestPair(ctx context.Context, req *pbuf.RequestPairRequest) (*pbuf.RequestPairResponse, error) {
	var accepted bool
	var reason string
	if s.PairManager.OnPairRequest != nil {
		accepted, reason = s.PairManager.OnPairRequest((*security.PairToken)(req.GetToken()), device.PbufToDetails(req.GetDetails()))
	} else {
		accepted = false
		reason = "App misconfigured, no pairing handler available"
	}
	return &pbuf.RequestPairResponse{
		Token:    req.GetToken(),
		Accepted: accepted,
		Reason:   reason,
		Details:  s.PairManager.DeviceDetails().ToPBUF(),
	}, nil
}
