// Package pair: stores information about alraedy paired device
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/core/storage"
	"alat/pkg/core/transport/client"
	"alat/pkg/pbuf"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
)

type pairResponse struct {
	accepted bool
	reason   string
}

type PairManager struct {
	deviceDetails   *device.Details
	storage         storage.NodeStorage
	pairedDevices   []device.PairedDevice
	onPairRequest   func(requestID string, token *security.PairToken, details *device.Details)
	pendingRequests map[string]chan pairResponse
	pendingMutex    sync.Mutex
}

func (p *PairManager) GetDeviceDetails() *device.Details {
	return p.deviceDetails
}
func (p *PairManager) SetDeviceDetails(d *device.Details) {
	p.deviceDetails = d
}

func (p *PairManager) OnPairRequest(handle func(requestID string, token *security.PairToken, details *device.Details)) {
	p.onPairRequest = handle
}
func (p *PairManager) AddPairedDevice(dev device.PairedDevice) {
	p.storage.AddPairedDevice(dev)
	p.pairedDevices = append(p.pairedDevices, dev)
}

func (p *PairManager) HandlePairRequest(token *security.PairToken, details *device.Details) (bool, string) {
	if p.onPairRequest == nil {
		return false, "Device cannot be paired because it is misconfigured, no pairing handler available"
	}

	requestID := uuid.New().String()
	responseChan := make(chan pairResponse, 1)

	p.pendingMutex.Lock()
	p.pendingRequests[requestID] = responseChan
	p.pendingMutex.Unlock()

	defer func() {
		p.pendingMutex.Lock()
		delete(p.pendingRequests, requestID)
		p.pendingMutex.Unlock()
	}()

	p.onPairRequest(requestID, token, details)

	select {
	case response := <-responseChan:
		if response.accepted {
			p.AddPairedDevice(device.PairedDevice{
				Certificate: details.Certificate,
				Token:       *token,
			})
		}
		return response.accepted, response.reason
	case <-time.After(time.Minute): // 30-second timeout
		return false, string(ResponseReasonTimeOut)
	}
}

func (p *PairManager) SubmitPairResponse(requestID string, accepted bool, reason string) error {
	p.pendingMutex.Lock()
	responseChan, ok := p.pendingRequests[requestID]
	p.pendingMutex.Unlock()

	if !ok {
		return fmt.Errorf("no pending pair request found for ID: %s", requestID)
	}

	responseChan <- pairResponse{accepted: accepted, reason: reason}
	return nil
}

func (p *PairManager) IsTokenValid(token security.PairToken) bool {
	return p.GetPairedDevice(token) != nil
}
func (p *PairManager) GetPairedDevice(token security.PairToken) *device.PairedDevice {
	for _, dev := range p.pairedDevices {
		if dev.Token == token {
			return &dev
		}
	}
	return nil
}

func NewManager(stor storage.NodeStorage, details *device.Details) (*PairManager, error) {
	paired, err := stor.GetPairedDevices()
	if err != nil {
		return nil, err
	}
	return &PairManager{
		storage:         stor,
		pairedDevices:   paired,
		deviceDetails:   details,
		pendingRequests: make(map[string]chan pairResponse),
	}, nil
}

func (p *PairManager) GetPairedDevices() []device.PairedDevice {
	return p.pairedDevices
}

func (p *PairManager) RequestPair(ip net.IP, port int) (*pbuf.RequestPairResponse, error) {
	pairToken, err := security.GeneratePairToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate pair token: %w", err)
	}
	response, err := client.RequestPair(ip, port, &pairToken, p.deviceDetails)
	if response.GetAccepted() {
		p.AddPairedDevice(device.PairedDevice{
			Token:       security.PairToken(response.GetToken()),
			Certificate: security.Certificate(response.GetDetails().GetCertificate()),
		})
	}
	return response, err
}
