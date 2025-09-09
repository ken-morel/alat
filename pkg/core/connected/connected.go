// Package connected: manage currently  connected devices
package connected

import (
	"alat/pkg/core/device"
	"alat/pkg/core/discovery"
	"alat/pkg/core/pair"
)

type Connected struct {
	Info         device.Info
	PairedDevice device.PairedDevice
}
type Manager struct {
	devices     []Connected
	pairManager *pair.PairManager
	discoverer  *discovery.Discoverer
}

func NewManageer(pair *pair.PairManager, discoverer *discovery.Discoverer) *Manager {
	return &Manager{
		pairManager: pair,
		discoverer:  discoverer,
	}
}

func (m *Manager) GetConnectedDevices() []Connected {
	return m.devices
}

func (m *Manager) RefreshConnections() error {
	var connected []Connected
	for _, found := range m.discoverer.GetFoundDevices() {
		for _, device := range m.pairManager.GetPairedDevices() {
			if found.Info.ID == device.Certificate.ID() {
				connected = append(connected, Connected{
					Info:         found.Info,
					PairedDevice: device,
				})
				break
			}
		}
	}
	clear(m.devices)
	m.devices = connected
	return nil
}
