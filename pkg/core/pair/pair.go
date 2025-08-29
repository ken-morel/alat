// Package pair: stores information about alraedy paired device
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/storage"
)

type PairManager struct {
	storage       *storage.NodeStorage
	callbacks     []PairCallback
	pairedDevices []device.PairedDevice
	details       *device.Details
}

func (p *PairManager) DeviceDetails() *device.Details {
	return p.details
}

func (p *PairManager) SetDetails(details *device.Details) {
	p.details = details
}

func NewManager(stor *storage.NodeStorage, details *device.Details) (*PairManager, error) {
	paired, err := (*stor).GetPaired()
	if err != nil {
		return nil, err
	}
	return &PairManager{
		storage:       stor,
		callbacks:     nil,
		pairedDevices: paired,
		details:       details,
	}, nil
}
