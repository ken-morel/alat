// Package service holds definitions of services and usage
package service

import "alat/pkg/core/pbuf"

type ServiceName string

const (
	RCFile  ServiceName = "rcfile"
	SysInfo ServiceName = "sysinfo"
)

type Service struct {
	Name    ServiceName `yaml:"name"`
	Enabled bool        `yaml:"enabled"`
}

func (s *Service) ToPBuf() pbuf.Service {
	return pbuf.Service{
		Name:    string(s.Name),
		Enabled: s.Enabled,
	}
}

func FromPBuf(pb *pbuf.Service) Service {
	return Service{
		Name:    ServiceName(pb.GetName()),
		Enabled: pb.GetEnabled(),
	}
}
