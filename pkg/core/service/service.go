// Package service: holds services and service deffinitions
package service

import "alat/pkg/core/device"

type (
	ServiceName string
)

const (
	FileSystem ServiceName = "filesystem"
	Clipboard  ServiceName = "clipboard"
)

type Service interface {
	Name() ServiceName
	Call(method string, params map[string]any) (any, error)
	Permissions() []device.PermissionName
}
