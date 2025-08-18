// Package rcfile: Stores api and cliend info for the sendfile service
package rcfile

type ServiceConfig struct {
	Enabled     bool   `yaml:"enabled"`
	Destination string `yaml:"destination"`
	FileMaxSize uint64 `yaml:"filemaxsize"`
}

const ServiceName = "rcfile"

var Config ServiceConfig = ServiceConfig{
	Enabled: false,
}
var Ready bool = false

func Init(conf ServiceConfig) error {
	Config = conf
	Ready = true
	return nil
}
