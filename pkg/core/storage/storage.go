// Package storage: holds persistent file storage methods
package storage

import (
	"alat/pkg/core/config"
	"alat/pkg/core/device"
)

type NodeStorage interface {
	GetPairedDevices() ([]device.PairedDevice, error)
	AddPairedDevice(device.PairedDevice) error

	GetAppConfig() (*config.AppConfig, error)
	SetAppConfig(config.AppConfig) error

	GetServiceConfig() (*config.ServiceConfig, error)
	SetServiceConfig(config.ServiceConfig) error
}
