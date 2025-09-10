package app

import (
	"fmt"

	"alat/apps/desktop/app/config"
	"alat/pkg/core/device/color"

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
	rt.WindowReloadApp(app.ctx)
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
