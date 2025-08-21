package storage

import (
	"fmt"
	"os"

	"github.com/gurkankaymak/hocon"
)

// PeerStorage defines the interface for storing and retrieving paired devices.
type PeerStorage interface {
	Save(device *PairedDevice) error
	Load() (map[string]*PairedDevice, error)
	Delete(deviceID string) error
}

// PairedDevice represents a device that has been paired with the node.
type PairedDevice struct {
	ID   string
	Name string
	// Add other device properties here
}

// HoconPeerStorage is an implementation of PeerStorage that uses HOCON for storage.
type HoconPeerStorage struct {
	filePath string
}

// NewHoconPeerStorage creates a new HoconPeerStorage.
func NewHoconPeerStorage(filePath string) (*HoconPeerStorage, error) {
	return &HoconPeerStorage{filePath: filePath}, nil
}

// Save saves a paired device to the HOCON file.
func (s *HoconPeerStorage) Save(device *PairedDevice) error {
	// This is a simplified implementation. A real implementation would need to
	// merge the new device with existing devices in the file.
	configString := fmt.Sprintf("peers.%s = { name = \"%s\" }", device.ID, device.Name)
	return os.WriteFile(s.filePath, []byte(configString), 0644)
}

// Load loads all paired devices from the HOCON file.
func (s *HoconPeerStorage) Load() (map[string]*PairedDevice, error) {
	conf, err := hocon.ParseFromFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]*PairedDevice), nil
		}
		return nil, err
	}

	devices := make(map[string]*PairedDevice)
	peersObject := conf.GetObject("peers")
	if peersObject == nil {
		return devices, nil
	}

	for deviceID, deviceObject := range peersObject.Items() {
		deviceConfig, ok := deviceObject.(*hocon.HoconObject)
		if !ok {
			continue
		}
		deviceName := deviceConfig.GetString("name")
		devices[deviceID] = &PairedDevice{
			ID:   deviceID,
			Name: deviceName,
		}
	}

	return devices, nil
}

// Delete removes a paired device from the HOCON file.
func (s *HoconPeerStorage) Delete(deviceID string) error {
	// This is a simplified implementation. A real implementation would need to
	// read the file, remove the device, and write the file back.
	return fmt.Errorf("delete not implemented")
}
