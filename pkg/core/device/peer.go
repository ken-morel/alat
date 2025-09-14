// Package device: holds peer related things
package device

import (
	"alat/pkg/core/security"
)

type PairedDevice struct {
	Certificate security.Certificate `yaml:"certificate" json:"certificate"`
	Token       security.PairToken   `yaml:"token"       json:"token"`
}
