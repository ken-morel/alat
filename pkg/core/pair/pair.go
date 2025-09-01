// Package pair: stores information about alraedy paired device
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/core/storage"
)

type PairRequest struct{}

type PairManager struct {
	storage       *storage.NodeStorage
	pairedDevices []device.PairedDevice
	details       *device.Details
	OnPairRequest func(*security.PairToken, *device.Details) (bool, string)
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
		pairedDevices: paired,
		details:       details,
	}, nil
}
