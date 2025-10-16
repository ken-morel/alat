// Package sysinfo provides functions to gather system information.
package sysinfo

import (
	"time"

	"alat/pkg/core/config"
	"alat/pkg/core/pair"
)

type Service struct {
	config   config.SysInfoConfig
	ready    bool
	cache    *SysInfo
	cacheAge time.Time

	pairManager *pair.PairManager
}

func (s *Service) Enabled() bool {
	return s.config.Enabled
}

func (s *Service) Get() (*SysInfo, error) {
	var cache *SysInfo
	var err error
	if s.cache == nil || time.Since(s.cacheAge) > time.Duration(s.config.CacheSeconds)*time.Second {
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

func (s *Service) Configure(c config.SysInfoConfig) {
	s.config = c
}

func CreateService(conf config.SysInfoConfig, p *pair.PairManager) Service {
	return Service{
		cacheAge:    time.Now(),
		cache:       nil,
		ready:       true,
		config:      conf,
		pairManager: p,
	}
}
