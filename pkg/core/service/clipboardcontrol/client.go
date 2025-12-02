package clipboardcontrol

import (
	"alat/pkg/core/connected"
	"alat/pkg/pbuf"
	"context"
	"fmt"
	"time"
)

func (s *Service) RequestGetClipboard(dev *connected.Connected) (*pbuf.ClipboardContent, error) {
	conn, err := connected.GetDeviceClientConnection(dev)
	if err != nil {
		return nil, err
	}
	client := pbuf.NewClipboardControlServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	clip, err := client.GetClipboard(ctx, &pbuf.GetClipboardRequest{
		Token:     dev.PairedDevice.Token[:],
		Requester: s.pairManager.GetDeviceDetails().GetInfo().ToPBUF(),
	})
	if err != nil {
		return nil, err
	} else if clip.GetStatus() != pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_OK {
		return nil, fmt.Errorf("%s", clip.GetMsg())
	} else {
		return clip.GetContent(), nil
	}
}

func (s *Service) RequestSetClipboard(dev *connected.Connected, content *pbuf.ClipboardContent) error {
	conn, err := connected.GetDeviceClientConnection(dev)
	if err != nil {
		return err
	}
	client := pbuf.NewClipboardControlServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	clip, err := client.SetClipboard(ctx, &pbuf.SetClipboardRequest{
		Token:     dev.PairedDevice.Token[:],
		Requester: s.pairManager.GetDeviceDetails().GetInfo().ToPBUF(),
		Content:   content,
	})
	if err != nil {
		return err
	} else if clip.GetStatus() != pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_OK {
		return fmt.Errorf("%s", clip.GetMsg())
	} else {
		return nil
	}
}
