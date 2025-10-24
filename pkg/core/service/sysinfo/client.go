package sysinfo

import (
	"context"
	"fmt"
	"net"
	"time"

	"alat/pkg/core/connected"
	"alat/pkg/pbuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *Service) Query(device *connected.Connected) (*SysInfo, error) {
	fullAddress := net.JoinHostPort(device.IP.String(), fmt.Sprintf("%d", device.Port))
	conn, err := grpc.NewClient(fullAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pbuf.NewSysInfoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	resp, err := client.GetSysInfo(ctx, &pbuf.GetSysInfoRequest{
		Token: device.PairedDevice.Token[:],
	})
	if err != nil {
		return nil, err
	}
	switch resp.Status {
	case pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_OK:
		return SysInfoFromPBUF(resp.Info), nil
	default:
		return nil, fmt.Errorf("error getting system info: %s", resp.Msg)
	}
}
