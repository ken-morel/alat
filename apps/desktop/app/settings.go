package app

import (
	"alat/apps/desktop/app/config"
	"alat/pkg/core/device/color"
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

func (app *App) GetAlatColors() []color.Color {
	return color.Colors
}
