package node

import (
	"fmt"

	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/pbuf"
)

func (n *Node) RequestPairFoundDevice(id string) (*pbuf.RequestPairResponse, error) {
	for _, found := range n.discovery.Discoverer.GetFoundDevices() {
		if found.Info.ID == id {
			return n.PairManager.RequestPair(found.IP, found.Port)
		}
	}
	return nil, fmt.Errorf("device with ID %s not found", id)
}

func (n *Node) OnPairRequest(handle func(requestID string, token *security.PairToken, details *device.Details)) {
	n.PairManager.OnPairRequest(handle)
}

func (n *Node) SubmitPairResponse(requestID string, accepted bool, reason string) error {
	return n.PairManager.SubmitPairResponse(requestID, accepted, reason)
}
