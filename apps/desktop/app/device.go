package app

import "alat/pkg/core"

func (app *App) GetAvailableAddresses() ([]core.DeviceAddress, error) {
	return core.GetAvailableAddresses()
}

func (app *App) GetPairedDevices() []core.DeviceInfo {
	return []core.DeviceInfo{
		{
			Name: "Jealomy",
			Address: core.DeviceAddress{
				IP:   "192.168.1.192",
				Port: 17,
			},
		},
		{
			Name: "Eloyd",
			Address: core.DeviceAddress{
				IP:   "192.168.1.122",
				Port: 17,
			},
		},
	}
}
