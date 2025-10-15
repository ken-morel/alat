package app

import (
	"alat/pkg/core/connected"
	"alat/pkg/core/service"

	"alat/pkg/pbuf"
)

//FIX: Move service configuration to service registry, when updating config, creating, init, ...
//TODO: ^

func (app *App) initServices() error {
	app.serviceRegistery = service.CreateRegistry(*app.serviceSettings)
	return nil
}

func (app *App) ServiceSysInfoGet(dev connected.Connected) (*pbuf.SysInfo, error) {
	return dev.GetSysInfo()
}
