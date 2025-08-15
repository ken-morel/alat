// Package service holds definitions of services and usage
package service

type ServiceName string

const (
	FileTransfer ServiceName = "rcfile"
	FsBrowse     ServiceName = "fsbrowse"
	Clipboard    ServiceName = "clipboard"
)

type Service struct {
	Name      ServiceName `yaml:"name"`
	Enabled   bool        `yaml:"enabled"`
	Protected bool        `yaml:"protected"`
}
