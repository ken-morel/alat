package node

import (
	"alat/pkg/core/config"
	"alat/pkg/core/connected"
	"alat/pkg/core/device"
)

func (n *Node) SetServiceConfig(conf config.ServiceConfig) error {
	n.services.UpdateConfig(conf)
	return n.storage.SetServiceConfig(conf)
}
func (n *Node) SetAppConfig(conf config.AppConfig) error {
	n.pairManager.SetDeviceDetails(&device.Details{
		Color:       conf.DeviceColor,
		Name:        conf.DeviceName,
		Type:        conf.DeviceType,
		Certificate: conf.Certificate,
	})
	return n.storage.SetAppConfig(conf)
}

func (n *Node) GetAppConfig() (*config.AppConfig, error) {
	return n.storage.GetAppConfig()
}
func (n *Node) GetServiceConfig() (*config.ServiceConfig, error) {
	return n.storage.GetServiceConfig()
}

func (n *Node) GetPairedDevices() []device.PairedDevice {
	return n.pairManager.GetPairedDevices()
}
func (n *Node) GetConnectedDevices() []connected.Connected {
	return n.connected.GetConnectedDevices()
}
