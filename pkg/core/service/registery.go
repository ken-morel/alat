package service

import (
	"alat/pkg/core/service/filesend"
	"alat/pkg/core/service/sysinfo"
)

type Registry struct {
	SysInfo  sysinfo.Service
	FileSend filesend.Service
}
