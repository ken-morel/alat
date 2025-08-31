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

func (app *App) SettingsGetUniversalClipboardSettings() config.UniversalClipboardSettings {
	return app.serviceSettings.UniversalClipboard
}

func (app *App) SettingsSetUniversalClipboardSettings(s config.UniversalClipboardSettings) error {
	app.serviceSettings.UniversalClipboard = s
	return config.SaveServiceSettings(app.serviceSettings)
}

func (app *App) SettingsGetNotificationSyncSettings() config.NotificationSyncSettings {
	return app.serviceSettings.NotificationSync
}

func (app *App) SettingsSetNotificationSyncSettings(s config.NotificationSyncSettings) error {
	app.serviceSettings.NotificationSync = s
	return config.SaveServiceSettings(app.serviceSettings)
}

func (app *App) SettingsGetMediaControlSettings() config.MediaControlSettings {
	return app.serviceSettings.MediaControl
}

func (app *App) SettingsSetMediaControlSettings(s config.MediaControlSettings) error {
	app.serviceSettings.MediaControl = s
	return config.SaveServiceSettings(app.serviceSettings)
}

func (app *App) SettingsGetRemoteInputSettings() config.RemoteInputSettings {
	return app.serviceSettings.RemoteInput
}

func (app *App) SettingsSetRemoteInputSettings(s config.RemoteInputSettings) error {
	app.serviceSettings.RemoteInput = s
	return config.SaveServiceSettings(app.serviceSettings)
}

func (app *App) SettingsGetFolderSyncSettings() config.FolderSyncSettings {
	return app.serviceSettings.FolderSync
}

func (app *App) SettingsSetFolderSyncSettings(s config.FolderSyncSettings) error {
	app.serviceSettings.FolderSync = s
	return config.SaveServiceSettings(app.serviceSettings)
}

func (app *App) AskFileSharingDestDirectory() string {
	dest, _ := rt.OpenDirectoryDialog(app.ctx, rt.OpenDialogOptions{
		Title: "Choose a location to save shared files",
	})
	return dest
}

func (app *App) SettingsSetSetupComplete(complete bool) error {
	app.settings.SetupComplete = complete
	return config.SaveAppSettings(app.settings)
}
