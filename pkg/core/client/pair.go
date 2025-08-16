package client

import (
	"alat/pkg/core/device"
	"alat/pkg/core/pbuf"
	"alat/pkg/core/service"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RequestPair sends a pairing request to another device.
func RequestPair(localDeviceInfo device.DeviceInfo, target device.DeviceInfo, services []service.Service) (*pbuf.PairResponse, error) {
	// 1. Establish a connection
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", target.Address.IP[0], target.Address.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("did not connect: %w", err)
	}
	defer conn.Close()

	// 2. Create the client
	client := pbuf.NewPairingClient(conn)

	// 3. Prepare the request
	localDevice := &pbuf.Device{
		Id:   localDeviceInfo.Code,
		Name: localDeviceInfo.Name,
		Type: pbuf.Device_DeviceType(localDeviceInfo.Type),
	}

	// Convert core.Service to pbuf.Service
	pbufServices := make([]*pbuf.Service, len(services))
	for i, s := range services {
		pbufServices[i] = &pbuf.Service{Name: s.Name, Enabled: s.Enabled}
	}

	req := &pbuf.PairRequest{
		Device:   localDevice,
		Services: pbufServices,
	}

	// 4. Send the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := client.RequestPair(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not send pair request: %w", err)
	}

	return res, nil
}

