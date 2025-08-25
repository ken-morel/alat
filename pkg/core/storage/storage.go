// Package storage: holds persistent file storage methods
package storage

import "alat/pkg/core/device"

type NodeStorage interface {
	Load() error
	Save() error
	GetPaired() (device.PairedDevice, error)
	AddPaired(device.PairedDevice) error
}

type YamlNodeStorage struct{}

type ChildPeerStorage struct {
	NodeStorage *NodeStorage
}
