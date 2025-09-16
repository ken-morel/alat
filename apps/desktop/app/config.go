package app

import (
	"alat/pkg/core"
	core_config "alat/pkg/core/config"
	"alat/pkg/core/device"
	"alat/pkg/core/node"
	"alat/pkg/core/pair"
	"alat/pkg/core/storage"
	"fmt"
	"os"
	"path"
)

func (app *App) ConfigReady() bool {
	return app.settings.SetupComplete
}

func (app *App) initConfig() bool {
	configDir, err := initAndGetConfigDir()
	if err != nil {
		fmt.Printf("Error initializing config directory: %v\n", err)
		return false
	}

	appSettingsPath := path.Join(configDir, "settings.yml")
	app.settings, err = core_config.LoadAppSettings(appSettingsPath)
	if err != nil {
		fmt.Println("Failed to load app settings:", err)
		return false
	}

	serviceSettingsPath := path.Join(configDir, "services.yml")
	app.serviceSettings, err = core_config.LoadServiceSettings(serviceSettingsPath)
	if err != nil {
		fmt.Println("Failed to load service settings:", err)
		return false
	}

	app.nodeStore, err = GetNodeStorage(configDir)
	if err != nil {
		fmt.Println("Failed to initialize node storage:", err)
		return false
	}

	fmt.Println("Initialized app config")
	return true
}

// initAndGetConfigDir determines the appropriate configuration directory, creates it if it doesn't exist,
// and returns the path.
func initAndGetConfigDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = path.Join(os.TempDir(), core.AppID)
	} else {
		dir = path.Join(dir, core.AppID)
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return dir, nil
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
	pairManager, err := pair.NewManager(app.nodeStore, app.nodeDetails)
	if err != nil {
		fmt.Println("Failed to initialize pair manager:", err)
		return err
	}
	pairManager.OnPairRequest(app.handlePairRequest)

	node, err := node.NewNode(app.serviceRegistery, app.nodeStore, app.nodeDetails, pairManager)
	app.node = node
	fmt.Println("Initialized node, got error: ", err)
	return err
}

func GetNodeStorage(configDir string) (storage.NodeStorage, error) {
	return storage.CreateYAMLNodeStorage(
		path.Join(configDir, "node.yml"),
	), nil
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
