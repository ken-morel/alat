package client

import (
	"context"
	"fmt"
	"net"
	"time"

	"alat/pkg/core/device"
	"alat/pkg/pbuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetDeviceInfo(addr net.IP, port int) (*device.Info, error) {
	// TODO: use connected connections
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
