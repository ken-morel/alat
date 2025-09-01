package node

import (
	"alat/pkg/core/device"
	"alat/pkg/core/transport"
	"fmt"
)

func (n *Node) StartDeviceSearch() error {
	// The actual search is now blocking, but we don't want to block the
	// frontend. The caller in app/pair.go will handle running this in a goroutine.
	err := n.discovery.Discoverer.Start()
	if err != nil {
		fmt.Printf("Device search failed: %v\n", err)
		return err
	}

	n.foundDevicesLock.Lock()
	n.foundDevices = n.discovery.Discoverer.Entries
	n.foundDevicesLock.Unlock()

	return nil
}

func (n *Node) GetFoundDevices() (devices []*device.Info, rerr error) {
	n.foundDevicesLock.Lock()
	defer n.foundDevicesLock.Unlock()

	for _, entry := range n.foundDevices {
		// Assuming AddrIPv4 has at least one element.
		// Production code should have more robust error handling here.
		if len(entry.AddrIPv4) > 0 {
			info, err := transport.GetDeviceInfo(entry.AddrIPv4[0], entry.Port)
			if err != nil {
				// Maybe collect errors instead of returning only the last one
				rerr = err
				fmt.Println("Error getting device info:", err)
			} else {
				devices = append(devices, info)
			}
		}
	}
	return
}

func (n *Node) SearchingDevices() bool {
	return n.discovery.Discoverer.Running
}
