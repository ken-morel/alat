package app

import (
	"time"

	"alat/pkg/core/service"
	"alat/pkg/core/service/sysinfo"
)

func (app *App) initServices() error {
	app.serviceRegistery = &service.Registry{
		SysInfo: sysinfo.CreateService(sysinfo.Config{
			Enabled:   app.serviceSettings.SysInfo.Enabled,
			CacheTime: time.Duration(app.serviceSettings.SysInfo.CacheSeconds) * time.Second,
		}),
	}
	return nil
}
