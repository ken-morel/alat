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

func (p *PairManager) IsTokenValid(token security.PairToken) bool {
	for _, dev := range p.pairedDevices {
		if dev.Token == token {
			return true
		}
	}
	return false
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

func (p *PairManager) GetPairedDevices() []device.PairedDevice {
	return p.pairedDevices
}
