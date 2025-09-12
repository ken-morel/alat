package node

import (
	"fmt"

	"alat/pkg/core/security"
	"alat/pkg/core/transport/client"
	"alat/pkg/pbuf"
)

func (n *Node) RequestPairFoundDevice(id string) (*pbuf.RequestPairResponse, error) {
	pairToken, err := security.GeneratePairToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate pair token: %w", err)
	}
	for _, found := range n.discovery.Discoverer.GetFoundDevices() {
		if found.Info.ID == id {
			return client.RequestPair(found.IP, found.Port, &pairToken, n.device)
		}
	}
	return nil, fmt.Errorf("device with ID %s not found", id)
}
