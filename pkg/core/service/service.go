// Package service: holds services and service deffinitions
package service

type (
	Name string
)

const (
	FileSystem Name = "filesystem"
	Clipboard  Name = "clipboard"
	SysInfo    Name = "sysinfo"
)
