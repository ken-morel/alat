package node

import (
	"alat/pkg/core/config"
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
