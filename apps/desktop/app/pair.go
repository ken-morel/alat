package app

import (
	"alat/pkg/core/device"
	"fmt"
)

func (app *App) SearchDevices() error {
	fmt.Println("[js call] Starting device search...")
	app.node.StartDeviceSearch()
	return nil
}

func (app *App) GetFoundDevices() ([]*device.Info, error) {
	return app.node.GetFoundDevices()
}

func (app *App) IsSearchingDevices() bool {
	return app.node.SearchingDevices()
}
