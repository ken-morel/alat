// Package pair: stores information about alraedy paired device
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/core/storage"
)

type PendingPair struct {
	Token security.PairToken
}

type PairManager struct {
	storage       *storage.NodeStorage
	callbacks     []PairCallback
	pairedDevices map[string]*device.PairedDevice
	pendingPairs  map[string]*PendingPair
}

func (p *PairManager) OnPairRequest(info device.DeviceInfo) error {
	return nil
}

func (p *PairManager) OnPeerResponse(info device.DeviceInfo) error {
	return nil
}

func (p *PairManager) OnPeerEvent(call PairCallback) {
	p.callbacks = append(p.callbacks, call)
}
func fireCallbacks(event *Event)

// type Manager interface {
// 	OnPairRequest(deviceInfo *DeviceInfo) error
// 	OnPeerEvent(callback PeerCallback)
// 	fireCallbacks(event PeerEvent, device *PairedDevice)
// }
