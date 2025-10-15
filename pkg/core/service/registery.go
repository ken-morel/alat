package service

import (
	"alat/pkg/core/config"
	"alat/pkg/core/service/filesend"
	"alat/pkg/core/service/sysinfo"
)

type Registry struct {
	SysInfo  sysinfo.Service
	FileSend filesend.Service
}

func (r *Registry) UpdateConfig(settings config.ServiceSettings) {
	r.FileSend.Configure(settings.FileSend)
	r.SysInfo.Configure(settings.SysInfo)
}
func CreateRegistry(settings config.ServiceSettings) *Registry {
	return &Registry{
		FileSend: filesend.CreateService(settings.FileSend),
		SysInfo:  sysinfo.CreateService(settings.SysInfo),
	}
}
