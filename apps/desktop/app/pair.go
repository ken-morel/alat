package app

import (
	"alat/pkg/core/client"
	"alat/pkg/core/device"
	"alat/pkg/core/pair"
	"alat/pkg/core/pbuf"
	"alat/pkg/core/service"
	"fmt"
)

func (app *App) HandlePairRequest(req *pbuf.PairRequest) {
	fmt.Println("Pair request received", req)
}

func (app *App) HandlePairResponse(res *pbuf.PairResponse) {
	fmt.Println("Pair response received", res)
}

func (app *App) RequestPair(deviceInfo device.DeviceInfo, services []service.Service) error {
	fmt.Println("Requesting pair from frontend")
	token := pair.GeneratePairToken()
	res, err := client.SendPairRequest(
		deviceInfo.Address,
		token,
		services,
	)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	if res {
		app.pandingPairs = append(app.pandingPairs, PendingPair{
			Device:   deviceInfo,
			Token:    token,
			Services: services,
		})
		fmt.Println("Succesfully send pair request")
		return nil
	} else {
		return fmt.Errorf("pairing request rejected by device")
	}
}
