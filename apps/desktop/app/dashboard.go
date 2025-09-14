package app

import (
	"fmt"

	"alat/pkg/core/device"
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
