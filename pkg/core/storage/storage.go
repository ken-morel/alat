// Package storage: holds persistent file storage methods
package storage

import "alat/pkg/core/device"

type NodeStorage interface {
	GetPaired() ([]device.PairedDevice, error)
	AddPaired(device.PairedDevice) error
}
