package app

import (
	"alat/apps/desktop/app/config"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (app *App) SettingsGetDeviceName() string {
	return app.settings.DeviceName
}

func (app *App) SettingsSetDeviceName(name string) error {
	app.settings.DeviceName = name
	return config.SaveAppSettings(app.settings)
}

func (app *App) SettingsGetDeviceColor() string {
	return app.settings.DeviceColor
}

func (app *App) SettingsSetDeviceColor(color string) error {
	app.settings.DeviceColor = color
	return config.SaveAppSettings(app.settings)
}

func (app *App) SettingsGetFileSharingSettings() config.FileSharingSettings {
	return app.serviceSettings.FileSharing
}

func (app *App) SettingsSetFileSharingSettings(s config.FileSharingSettings) error {
	app.serviceSettings.FileSharing = s
	return config.SaveServiceSettings(app.serviceSettings)
}

func (app *App) AskFileSharingDestDirectory() string {
	dest, _ := rt.OpenDirectoryDialog(app.ctx, rt.OpenDialogOptions{
		Title: "Choose a location to save shared files",
	})
	return dest
}
