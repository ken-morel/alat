package client

import (
	"context"
	"fmt"
	"net"
	"time"

	"alat/pkg/core/security"
	"alat/pkg/pbuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetSysInfo(addr net.IP, port int, token security.PairToken) (*pbuf.SysInfo, error) {
	fullAddress := net.JoinHostPort(addr.String(), fmt.Sprintf("%d", port))
	conn, err := grpc.NewClient(fullAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pbuf.NewAlatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	resp, err := client.GetSysInfo(ctx, &pbuf.GetSysInfoRequest{
		Token: token[:],
	})
	if err != nil {
		return nil, err
	}
	switch resp.Status {
	case pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_OK:
		return resp.Info, nil
	default:
		return nil, fmt.Errorf("error getting system info: %s", resp.Msg)
	}
}
