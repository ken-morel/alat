// Package pair: stores information about alraedy paired device
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/core/storage"
	"alat/pkg/core/transport/client"
	"alat/pkg/pbuf"
	"fmt"
	"net"
)

type PairManager struct {
	deviceDetails *device.Details
	storage       storage.NodeStorage
	pairedDevices []device.PairedDevice
	onPairRequest func(*security.PairToken, *device.Details) (bool, string)
}

func (p *PairManager) GetDeviceDetails() *device.Details {
	return p.deviceDetails
}
func (p *PairManager) SetDeviceDetails(d *device.Details) {
	p.deviceDetails = d
}

func (p *PairManager) OnPairRequest(handle func(*security.PairToken, *device.Details) (bool, string)) {
	p.onPairRequest = handle
}
func (p *PairManager) AddPairedDevice(dev device.PairedDevice) {
	p.storage.AddPairedDevice(dev)
	p.pairedDevices = append(p.pairedDevices, dev)
}

func (p *PairManager) HandlePairRequest(token *security.PairToken, details *device.Details) (bool, string) {
	if p.onPairRequest != nil && p.storage != nil {
		accepted, reason := p.onPairRequest(token, details)
		if accepted {
			p.AddPairedDevice(device.PairedDevice{
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

func NewManager(stor storage.NodeStorage, details *device.Details) (*PairManager, error) {
	paired, err := stor.GetPairedDevices()
	if err != nil {
		return nil, err
	}
	return &PairManager{
		storage:       stor,
		pairedDevices: paired,
		deviceDetails: details,
	}, nil
}

func (p *PairManager) GetPairedDevices() []device.PairedDevice {
	return p.pairedDevices
}

func (p *PairManager) RequestPair(ip net.IP, port int) (*pbuf.RequestPairResponse, error) {
	pairToken, err := security.GeneratePairToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate pair token: %w", err)
	}
	response, err := client.RequestPair(ip, port, &pairToken, p.deviceDetails)
	if response.GetAccepted() {
		p.AddPairedDevice(device.PairedDevice{
			Token:       security.PairToken(response.GetToken()),
			Certificate: security.Certificate(response.GetDetails().GetCertificate()),
		})
	}
	return response, err
}
