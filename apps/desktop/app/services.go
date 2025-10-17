package app

import (
	"alat/pkg/core/connected"
	"alat/pkg/core/service/sysinfo"
)

func (app *App) QueryDeviceSysInfo(dev connected.Connected) (*sysinfo.SysInfo, error) {
	return app.node.QueryDeviceSysInfo(&dev)
}
