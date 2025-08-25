// Package transport: holds server and client methods
package transport

import (
	"alat/pkg/core/pair"
	"alat/pkg/core/service"
)

type Server struct {
	Services    *service.Registry
	PairManager *pair.PairManager
}
