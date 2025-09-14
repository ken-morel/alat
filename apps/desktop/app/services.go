package app

import (
	"time"

	"alat/pkg/core/connected"
	"alat/pkg/core/service"
	"alat/pkg/core/service/filesend"
	"alat/pkg/core/service/sysinfo"
	"alat/pkg/pbuf"
)

func (app *App) initServices() error {
	app.serviceRegistery = &service.Registry{
		SysInfo: sysinfo.CreateService(sysinfo.Config{
			Enabled:   app.serviceSettings.SysInfo.Enabled,
			CacheTime: time.Duration(app.serviceSettings.SysInfo.CacheSeconds) * time.Second,
		}),
		FileSend: filesend.CreateService(filesend.Config{
			Enabled: app.serviceSettings.FileSend.Enabled,
		}),
	}
	return nil
}

func (app *App) ServiceSysInfoGet(dev connected.Connected) (*pbuf.SysInfo, error) {
	return dev.GetSysInfo()
}
