package clipboardcontrol

import (
	"alat/pkg/core/security"
	"alat/pkg/pbuf"
	"context"
)

type ClipControlServer struct {
	pbuf.UnimplementedClipboardControlServiceServer
	Service *Service
}

func (s *ClipControlServer) GetClipboard(ctx context.Context, req *pbuf.GetClipboardRequest) (*pbuf.GetClipboardResponse, error) {
	dev := s.Service.pairManager.GetPairedDevice(security.PairToken(req.Token))
	if dev == nil {
		return &pbuf.GetClipboardResponse{
			Status: pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_UNAUTHORIZED,
			Msg:    "Pair token not recognised, device not paired",
		}, nil
	} else {
		content, err := s.Service.getHandler(*dev)
		if err != nil {
			return &pbuf.GetClipboardResponse{
				Status:  pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_ERROR,
				Msg:     err.Error(),
				Content: content,
			}, nil
		} else {
			return &pbuf.GetClipboardResponse{
				Status:  pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_OK,
				Msg:     "All okay",
				Content: content,
			}, err
		}
	}
}

func (s *ClipControlServer) SetClipboard(ctx context.Context, req *pbuf.SetClipboardRequest) (*pbuf.SetClipboardResponse, error) {
	if !s.Service.config.Enabled {
		return &pbuf.SetClipboardResponse{
			Status: pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_DISABLED,
			Msg:    "Service disabled",
		}, nil
	} else if !s.Service.config.CanReceive {
		return &pbuf.SetClipboardResponse{
			Status: pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_DISABLED,
			Msg:    "Clipboard receiving disabled",
		}, nil
	}
	dev := s.Service.pairManager.GetPairedDevice(security.PairToken(req.Token))
	if dev == nil {
		return &pbuf.SetClipboardResponse{
			Status: pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_UNAUTHORIZED,
			Msg:    "Pair token not recognised, device not paired",
		}, nil
	} else {
		err := s.Service.setHandler(*dev, req.Content)
		if err != nil {
			return &pbuf.SetClipboardResponse{
				Status: pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_ERROR,
				Msg:    err.Error(),
			}, nil
		} else {
			return &pbuf.SetClipboardResponse{
				Status: pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_OK,
				Msg:    "All okay",
			}, err
		}
	}
}
