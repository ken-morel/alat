package app

import (
	"alat/pkg/core/pair"
	"alat/pkg/core/pbuf"
	"alat/pkg/core/service/sysinfo"
)

func (app *App) GetThisDeviceSysInfo() (*pbuf.SysInfo, error) {
	return sysinfo.Get()
}

func (app *App) GetPairedDeviceSysInfo(device pair.Pair) (*pbuf.SysInfo, error) {
	return sysinfo.GetInfo(device.DeviceInfo.Address, device.Token)
}
