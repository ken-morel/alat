package node

import (
	"alat/pkg/core/security"
	"alat/pkg/core/transport"
	"alat/pkg/pbuf"
	"fmt"
)

func (n *Node) RequestPairFoundDevice(id string) (*pbuf.RequestPairResponse, error) {
	pairToken, err := security.GeneratePairToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate pair token: %w", err)
	}
	for _, found := range n.discovery.Discoverer.GetFoundDevices() {
		if found.Info.ID == id {
			return transport.RequestPair(found.IP, found.Port, &pairToken, n.device)
		}
	}
	return nil, fmt.Errorf("device with ID %s not found", id)
}
