// Package pair: stores information about alraedy paired device
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/storage"
)

type PairManager struct {
	storage       *storage.NodeStorage
	callbacks     []PairCallback
	pairedDevices map[string]*device.PairedDevice
	info          *device.Info
}

func (p *PairManager) SetInfo(info *device.Info) {
	p.info = info
}

func (p *PairManager) OnPeerEvent(call PairCallback) {
	p.callbacks = append(p.callbacks, call)
}
