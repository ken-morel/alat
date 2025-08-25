```go
type PeerStorage interface {
	Save(device *PairedDevice) error
	Load() (map[string]*PairedDevice, error)
	Delete(deviceID string) error
}
```