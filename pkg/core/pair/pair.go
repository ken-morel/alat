// Package pair: stores information about alraedy paired device
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/core/storage"
	"errors"
)

type PairingSession struct {
	Initiator device.DeviceInfo
	Responder device.DeviceInfo
	Token     security.PairToken
}

type PairManager struct {
	storage        *storage.NodeStorage
	callbacks      []PairCallback
	pairedDevices  map[string]*device.PairedDevice
	pairingSessions map[string]*PairingSession
}

func (p *PairManager) InitiatePairing(info device.DeviceInfo) (*PairingSession, error) {
	token, err := security.GeneratePairToken()
	if err != nil {
		return nil, err
	}
	session := &PairingSession{
		Initiator: info,
		Token:     token,
	}
	p.pairingSessions[info.ID] = session
	return session, nil
}

func (p *PairManager) FinalizePairing(token security.PairToken, responderInfo device.DeviceInfo) (*device.PairedDevice, error) {
	var initiatorID string
	for id, session := range p.pairingSessions {
		if session.Token == token {
			initiatorID = id
			break
		}
	}

	if initiatorID == "" {
		return nil, errors.New("pairing session not found")
	}

	delete(p.pairingSessions, initiatorID)

	cert, err := security.GenerateCertificate()
	if err != nil {
		return nil, err
	}

	pairedDevice := &device.PairedDevice{
		Certificate: cert,
		Token:       token,
	}

	p.pairedDevices[responderInfo.ID] = pairedDevice

	return pairedDevice, nil
}

func (p *PairManager) OnPeerEvent(call PairCallback) {
	p.callbacks = append(p.callbacks, call)
}
