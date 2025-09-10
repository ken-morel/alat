// Package sysinfo provides functions to gather system information.
package sysinfo

import (
	"time"

	"alat/pkg/core/device"
	"alat/pkg/core/service"
	"alat/pkg/pbuf"
)

type Config struct {
	Enabled   bool          `yaml:"enabled"`
	CacheTime time.Duration `yaml:"cachetime"`
}
type Service struct {
	Config   Config
	Ready    bool
	cache    pbuf.SysInfo
	cacheAge time.Time
}

func (*Service) Name() service.ServiceName {
	return service.SysInfo
}

func (*Service) Permissions() []device.PermissionName {
	return nil
}

func (s *Service) Enabled() bool {
	return s.Config.Enabled
}

func (s *Service) Call(method string, args map[string]any) (any, error) {
	if method == "get" {
	}
}

func CreateService(conf Config) Service {
	return Service{
		Config: conf,
	}
}
