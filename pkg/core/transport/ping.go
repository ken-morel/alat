package transport

import (
	"alat/pkg/core/device"
	"alat/pkg/pbuf"
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetDeviceInfo(addr net.IP, port int) (*device.Info, error) {
	fullAddress := net.JoinHostPort(addr.String(), fmt.Sprintf("%d", port))
	conn, err := grpc.NewClient(fullAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pbuf.NewAlatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetInfo(ctx, &pbuf.GetInfoRequest{})
	if err != nil {
		return nil, err
	} else {
		return device.PbufToInfo(resp.GetInfo()), nil
	}
}
