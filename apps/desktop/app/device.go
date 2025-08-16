package app

import (
	"alat/apps/desktop/app/config"
	"alat/pkg/core/client"
	"alat/pkg/core/device"
	"alat/pkg/core/pair"
	"fmt"
)

func (app *App) SearchDevices() []device.DeviceInfo {
	channel := make(chan device.DeviceInfo)
	var infos []device.DeviceInfo
	go client.SearchDevices(channel)
	for info := range channel {
		infos = append(infos, info)
	}
	fmt.Println("Searched devices, got:", infos)
	return infos
}

func (app *App) GetPairedDevices() ([]pair.Pair, error) {
	var pairs []pair.Pair
	var err error
	fmt.Println("Getting paired devices")
	pairs, err = config.GetPairedDevices()
	fmt.Println("Returning paired devices", pairs)
	return pairs, err
}
