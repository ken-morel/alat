// package clibpardcontrol: provides clipboard control service
package clipboardcontrol

import (
	"alat/pkg/core/config"
	"alat/pkg/core/device"
	"alat/pkg/core/pair"
	"alat/pkg/pbuf"
)

type Service struct {
	config config.ClipControlConfig
	ready  bool

	pairManager *pair.PairManager

	setHandler func(device.PairedDevice, *pbuf.ClipboardContent) error
	getHandler func(device.PairedDevice) (*pbuf.ClipboardContent, error)
}

func (s *Service) Configure(c config.ClipControlConfig) {
	s.config = c
}
func (s *Service) Initialize(setter func(device.PairedDevice, *pbuf.ClipboardContent) error, getter func(device.PairedDevice) (*pbuf.ClipboardContent, error)) {
	s.setHandler = setter
	s.getHandler = getter
}

func CreateService(c config.ClipControlConfig, p *pair.PairManager) Service {
	return Service{
		ready:       false,
		config:      c,
		pairManager: p,
	}
}
