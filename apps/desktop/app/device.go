package app

import (
	"alat/pkg/core"
	"net"
)

func (app *App) GetAvailableAddresses() ([]core.Address, error) {
	return core.GetAvailableAddresses()
}

func (app *App) GetPairedDevices() []core.DeviceInfo {
	return []core.DeviceInfo{
		{
			Name: "jealomy",
			Address: core.Address{
				Port: 307,
				IP:   net.IP{},
			},
		},
	}
}
