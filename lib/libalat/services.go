package main

import (
	"alat/pkg/core/config"
	"alat/pkg/core/service"
	"alat/pkg/core/service/filesend"
	"alat/pkg/core/service/sysinfo"
	"time"
)

func initServices(serviceSettings *config.ServiceSettings) service.Registry {

	return service.Registry{
		SysInfo: sysinfo.CreateService(sysinfo.Config{
			Enabled:   serviceSettings.SysInfo.Enabled,
			CacheTime: time.Duration(serviceSettings.SysInfo.CacheSeconds) * time.Second,
		}),
		FileSend: filesend.CreateService(filesend.Config{
			Enabled: true,
		}),
	}
}
