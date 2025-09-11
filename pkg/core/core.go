// Package core store's alat core
package core

const (
	AppID    = "cm.engon.alat"
	AlatPort = 12121
)

type (
	ServiceName string
)

const (
	FileSystem ServiceName = "filesystem"
	Clipboard  ServiceName = "clipboard"
	SysInfo    ServiceName = "sysinfo"
)
