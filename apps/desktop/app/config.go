package app

import (
	"alat/apps/desktop/app/config"
	"alat/pkg/core/device"
	"alat/pkg/core/node"
	"alat/pkg/core/pair"
	"fmt"
)

func (app *App) ConfigReady() bool {
	return app.settings.SetupComplete
}

func (app *App) initConfig() bool {
	fmt.Println("Initializing app")
	err := config.Init()
	fmt.Println("Initialized config, got")
	if err != nil {
		fmt.Printf("Error initializing config: %v\n", err)
	}
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
	fmt.Println("Initialized app config")
	return okay
}

func (app *App) initNode() error {
	fmt.Println("Initializing node")
	app.initServices()
	app.nodeDetails = &device.Details{
		Color:       app.settings.DeviceColor,
		Name:        app.settings.DeviceName,
		Type:        device.DesktopDevice,
		Certificate: app.settings.Certificate,
	}
	pairManager, err := pair.NewManager(&app.nodeStore, app.nodeDetails)
	if err != nil {
		fmt.Println("Failed to initialize pair manager:", err)
		return err
	}
	pairManager.OnPairRequest = app.handlePairRequest

	node, err := node.NewNode(app.serviceRegistery, &app.nodeStore, app.nodeDetails, pairManager)
	app.node = node
	fmt.Println("Initialized node, got error: ", err)
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
