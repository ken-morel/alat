A Peer Manager is a device which manages peers connected to and by a [[Node]]. A peermanager stores information in a [[Peer Storage]] class in `storage module`.
A peermanager also holds some callback functions in a [[Peer Manager Callbacks]] structure. And an in memory cache of [[Paired Device]]s.

## callbacks
They are just stored in a callback array when any event occurs.
```go
type PeerEvent int

const (
	PeerDiscovered PeerEvent = iota
	PeerPaired
	PeerConnected
	PeerDisconnected
)

type PeerCallback func(event PeerEvent, device *PairedDevice)
```

## Impl

```go

// Manager manages the peers of the node.
type Manager struct {
	storage       storage.PeerStorage
	callbacks     []PeerCallback
	pairedDevices map[string]*PairedDevice
    pendingPairs map[string]*Pending...
}
interface Manager {
    OnPairRequest(deviceInfo *DeviceInfo) error
    OnPeerEvent(callback PeerCallback)
    fireCallbacks(event PeerEvent, device *PairedDevice)
}
