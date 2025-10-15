package node

import (
	"fmt"

	"alat/pkg/pbuf"
)

func (n *Node) RequestPairFoundDevice(id string) (*pbuf.RequestPairResponse, error) {
	for _, found := range n.discovery.Discoverer.GetFoundDevices() {
		if found.Info.ID == id {
			return n.pairManager.RequestPair(found.IP, found.Port)
		}
	}
	return nil, fmt.Errorf("device with ID %s not found", id)
}
