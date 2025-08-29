package pair

import "alat/pkg/core/device"

type Event int

const (
	PeerDiscovered Event = iota
	PeerPaired
	PeerConnected
	PeerDisconnected
)

type PairCallback func(event Event, device *device.PairedDevice)

func (p *PairManager) OnPeerEvent(call PairCallback) {
	p.callbacks = append(p.callbacks, call)
}
