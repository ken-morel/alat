package node

import (
	"alat/pkg/core/device"
	"alat/pkg/core/transport"
	"fmt"
)

func (n *Node) StartDeviceSearch() error {
	return n.discovery.Discoverer.Start()
}

func (n *Node) GetFoundDevices() (devices []*device.Info, rerr error) {
	found := n.discovery.Discoverer.Entries
	for _, entry := range found {
		info, err := transport.GetDeviceInfo(entry.AddrIPv4[0], entry.Port)
		if err != nil {
			rerr = err
			fmt.Println("Error getting device info:", err)
		} else {
			devices = append(devices, info)
		}
	}
	return
}

func (n *Node) SearchingDevices() bool {
	return n.discovery.Discoverer.Running
}
