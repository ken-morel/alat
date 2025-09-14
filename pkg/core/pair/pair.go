// Package pair: stores information about alraedy paired device
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/core/storage"
)

type PairRequest struct{}

type PairManager struct {
	storage       storage.NodeStorage
	pairedDevices []device.PairedDevice
	details       *device.Details
	onPairRequest func(*security.PairToken, *device.Details) (bool, string)
}

func (p *PairManager) OnPairRequest(handle func(*security.PairToken, *device.Details) (bool, string)) {
	p.onPairRequest = handle
}

func (p *PairManager) HandlePairRequest(token *security.PairToken, details *device.Details) (bool, string) {
	if p.onPairRequest != nil && p.storage != nil {
		accepted, reason := p.onPairRequest(token, details)
		if accepted {
			p.storage.AddPaired(device.PairedDevice{
				Certificate: details.Certificate,
				Token:       *token,
			})
		}
		return accepted, reason
	} else {
		return false, "App misconfigured, no pairing handler available"
	}
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

func NewManager(stor storage.NodeStorage, details *device.Details) (*PairManager, error) {
	paired, err := stor.GetPaired()
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
