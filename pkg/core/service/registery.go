package service

import (
	"alat/pkg/core/config"
	"alat/pkg/core/pair"
	"alat/pkg/core/service/clipboardcontrol"
	"alat/pkg/core/service/filesend"
	"alat/pkg/core/service/sysinfo"
	"alat/pkg/core/service/webshare"
)

type Registry struct {
	SysInfo     sysinfo.Service
	FileSend    filesend.Service
	WebShare    webshare.Service
	ClipControl clipboardcontrol.Service
}

func (r *Registry) UpdateConfig(settings config.ServiceConfig) {
	r.FileSend.Configure(settings.FileSend)
	r.SysInfo.Configure(settings.SysInfo)
	r.WebShare.Configure(settings.FileSend)
	r.ClipControl.Configure(settings.ClipControl)
}

func CreateRegistry(settings *config.ServiceConfig, p *pair.PairManager) *Registry {
	return &Registry{
		FileSend:    filesend.CreateService(settings.FileSend, p),
		SysInfo:     sysinfo.CreateService(settings.SysInfo, p),
		WebShare:    webshare.CreateService(&settings.FileSend, p),
		ClipControl: clipboardcontrol.CreateService(settings.ClipControl, p),
	}
}
