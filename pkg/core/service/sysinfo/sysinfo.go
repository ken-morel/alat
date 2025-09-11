// Package sysinfo provides functions to gather system information.
package sysinfo

import (
	"time"

	"alat/pkg/pbuf"
)

type Config struct {
	Enabled   bool
	CacheTime time.Duration
}

type Service struct {
	config   Config
	ready    bool
	cache    *pbuf.SysInfo
	cacheAge time.Time
}

func (s *Service) Enabled() bool {
	return s.config.Enabled
}

func (s *Service) Get() (*pbuf.SysInfo, error) {
	var cache *pbuf.SysInfo
	var err error
	if s.cache == nil || time.Since(s.cacheAge) > s.config.CacheTime {
		cache, err = GetSysInfo()
		s.cacheAge = time.Now()
		if err != nil {
			s.cache = nil
			return nil, err
		} else {
			s.cache = cache
		}
	}
	return s.cache, nil
}

func (s *Service) Configure(c Config) {
	s.config = c
}

func CreateService(conf Config) Service {
	return Service{
		cacheAge: time.Now(),
		cache:    nil,
		ready:    true,
		config:   conf,
	}
}
