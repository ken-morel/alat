package app

import (
	"alat/apps/desktop/app/config"
	"alat/pkg/core/service/rcfile"
	"alat/pkg/core/service/sysinfo"
)

func (app *App) SetupServices() {
	conf := config.GetConfig().Services
	rcfile.Init(conf.RCFile, app.onFileReceive)
	sysinfo.Init(conf.SysInfo)
}
