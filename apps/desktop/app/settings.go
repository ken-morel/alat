package app

import (
	"alat/pkg/core/config"
	"alat/pkg/core/device/color"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (app *App) ConfigReady() bool {
	conf, err := app.node.GetAppConfig()
	return err == nil && conf.SetupComplete
}
func (app *App) SettingsGetDeviceName() string {
	conf, err := app.node.GetAppConfig()
	if err != nil {
		return ""
	} else {
		return conf.DeviceName
	}
}

func (app *App) SettingsSetDeviceName(name string) error {
	conf, err := app.node.GetAppConfig()
	if err != nil {
		return err
	}
	conf.DeviceName = name
	return app.node.SetAppConfig(*conf)
}

func (app *App) SettingsGetDeviceColorName() (string, error) {
	conf, err := app.node.GetAppConfig()
	if err != nil {
		return "", err
	} else {
		return conf.DeviceColor.Name, nil
	}
}

func (app *App) SettingsSetDeviceColorName(colName string) error {
	col := color.FromString(colName)
	conf, err := app.node.GetAppConfig()
	if err != nil {
		return err
	}
	conf.DeviceColor = *col
	return app.node.SetAppConfig(*conf)
}

func (app *App) AskFileSharingDestDirectory() string {
	dest, _ := rt.OpenDirectoryDialog(app.ctx, rt.OpenDialogOptions{
		Title: "Choose a location to save shared files",
	})
	return dest
}

func (app *App) SettingsSetSetupComplete(complete bool) error {
	conf, err := app.node.GetAppConfig()
	if err != nil {
		return err
	} else {
		conf.SetupComplete = complete
		return app.node.SetAppConfig(*conf)
	}
}

func (app *App) SettingsGetSysInfo() (*config.SysInfoConfig, error) {
	conf, err := app.node.GetServiceConfig()
	if err != nil {
		return nil, err
	} else {
		return &conf.SysInfo, err
	}
}

func (app *App) SettingsSetSysInfo(newConfig config.SysInfoConfig) error {
	conf, err := app.node.GetServiceConfig()
	if err != nil {
		return err
	} else {
		conf.SysInfo = newConfig
		return app.node.SetServiceConfig(*conf)
	}
}

func (app *App) SettingsGetFileSend() (*config.FileSendConfig, error) {
	conf, err := app.node.GetServiceConfig()
	if err != nil {
		return nil, err
	} else {
		return &conf.FileSend, err
	}
}

func (app *App) SettingsSetFileSend(newConfig config.FileSendConfig) error {
	conf, err := app.node.GetServiceConfig()
	if err != nil {
		return err
	} else {
		conf.FileSend = newConfig
		return app.node.SetServiceConfig(*conf)
	}
}
