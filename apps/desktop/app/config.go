package app

import (
	"alat/apps/desktop/app/config"
	"alat/pkg/core/device"
	"alat/pkg/core/node"
	"alat/pkg/core/service"
	"fmt"
)

func (app *App) ConfigReady() bool {
	return app.settings.SetupComplete
}

func (app *App) initConfig() bool {
	err := config.Init()
	okay := true
	if err != nil {
		okay = false
		fmt.Println("Failed to initialize config:", err)
	}
	app.settings, err = config.LoadAppSettings()
	if err != nil {
		okay = false
		fmt.Println("Failed to load app settings:", err)
	}
	app.serviceSettings, err = config.LoadServiceSettings()
	if err != nil {
		okay = false
		fmt.Println("Failed to load service settings:", err)
	}
	app.nodeStore, err = config.GetNodeStorage()
	if err != nil {
		okay = false
		fmt.Println("Failed to initialize node storage:", err)
	}
	return okay
}

func (app *App) initNode() error {
	app.serviceRegistery = service.NewRegistery()
	app.nodeDetails = &device.Details{
		Color:       app.settings.DeviceColor,
		Name:        app.settings.DeviceName,
		Type:        device.DesktopDevice,
		Certificate: app.settings.Certificate,
	}

	node, err := node.NewNode(app.serviceRegistery, &app.nodeStore, app.nodeDetails)
	app.node = node
	return err
}

func (app *App) updateNode() error {
	app.nodeDetails.Color = app.settings.DeviceColor
	app.nodeDetails.Name = app.settings.DeviceName
	app.nodeDetails.Certificate = app.settings.Certificate
	if app.node != nil {
		app.node.SetDetails(app.nodeDetails)
	}
	return nil
}
