package main

import (
	"alat/pkg/core/config"
	"alat/pkg/core/service"
)

func initServices(serviceSettings *config.ServiceSettings) *service.Registry {
	return service.CreateRegistry(serviceSettings)
}
