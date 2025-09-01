package transport

import (
	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/pbuf"
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RequestPair(addr net.IP, port int, token *security.PairToken, details *device.Details) (*pbuf.RequestPairResponse, error) {
	fullAddress := net.JoinHostPort(addr.String(), fmt.Sprintf("%d", port))
	conn, err := grpc.NewClient(fullAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pbuf.NewAlatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.RequestPair(ctx, &pbuf.RequestPairRequest{
		Token:   token[:],
		Details: details.ToPBUF(),
	})
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
