package app

import (
	"alat/pkg/core/device"
	"fmt"
)

func (app *App) GetPairedDevices() []device.PairedDevice {
	paired, err := app.nodeStore.GetPaired()
	if err != nil {
		fmt.Println("Error getting paired devices:", err)
		return nil
	} else {
		return paired
	}
}
