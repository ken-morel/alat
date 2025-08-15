package app

import (
	"alat/pkg/core/address"
	"alat/pkg/core/client"
	"alat/pkg/core/device"
	"net"

	"github.com/wailsapp/wails/v2/pkg/options"
)

func (app *App) SearchDevices() []device.DeviceInfo {
	channel := make(chan device.DeviceInfo)
	var infos []device.DeviceInfo
	go client.SearchDevices(channel)
	for info := range channel {
		infos = append(infos, info)
	}
	return infos
}

func (app *App) GetPairedDevices() []device.DeviceInfo {
	return []device.DeviceInfo{
		{
			Name: "jealomy",
			Address: address.Address{
				Port: 307,
				IP:   net.IP{192, 168, 1, 1},
			},
			Color: options.RGBA{
				R: 29,
				G: 49,
				B: 100,
				A: 1,
			},
			Code: "showig awfioanfuiofunpiaondf",
			Type: device.DeviceTypeTV,
		},
	}
}
