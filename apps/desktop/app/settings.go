package app

import (
	"fmt"
	"time"

	"alat/apps/desktop/app/config"
	"alat/pkg/core/device/color"
	"alat/pkg/core/service/sysinfo"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (app *App) SettingsGetDeviceName() string {
	return app.settings.DeviceName
}

func (app *App) SettingsSetDeviceName(name string) error {
	app.settings.DeviceName = name
	return config.SaveAppSettings(app.settings)
}

func (app *App) SettingsGetDeviceColorName() string {
	return app.settings.DeviceColor.Name
}

func (app *App) SettingsSetDeviceColorName(colName string) error {
	col := color.FromString(colName)
	if col == nil {
		return fmt.Errorf("color not registerred")
	} else {

		app.settings.DeviceColor = *col
		return config.SaveAppSettings(app.settings)
	}
}

func (app *App) AskFileSharingDestDirectory() string {
	dest, _ := rt.OpenDirectoryDialog(app.ctx, rt.OpenDialogOptions{
		Title: "Choose a location to save shared files",
	})
	return dest
}

func (app *App) SettingsSetSetupComplete(complete bool) error {
	app.settings.SetupComplete = complete
	if app.node != nil && !app.node.GetStatus().DiscoveryRunning && !app.node.GetStatus().ServerRunning {
		fmt.Println("Node is not running, starting")
		err := app.node.Start()
		if err != nil {
			fmt.Println("Error starting node after setup complete:", err)
		}
	} else {
		fmt.Println("Node is already running, not starting")
	}
	return config.SaveAppSettings(app.settings)
}

func (app *App) SettingsGetSysInfo() config.SysInfoSettings {
	return app.serviceSettings.SysInfo
}

func (app *App) SettingsSetSysInfo(conf config.SysInfoSettings) error {
	app.serviceRegistery.SysInfo.Configure(sysinfo.Config{
		Enabled:   conf.Enabled,
		CacheTime: time.Duration(conf.CacheSeconds) * time.Second,
	})
	app.serviceSettings.SysInfo = conf

	return config.SaveAppSettings(app.settings)
}
