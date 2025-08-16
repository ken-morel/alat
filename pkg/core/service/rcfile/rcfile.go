// Package rcfile: Stores api and cliend info for the sendfile service
package rcfile

import "fmt"

type ServiceConfig struct {
	Enabled     bool   `yaml:"enabled"`
	Destination string `yaml:"destination"`
	FileMaxSize uint64 `yaml:"filemaxsize"`
}

const ServiceName = "rcfile"

var config ServiceConfig = ServiceConfig{
	Enabled: false,
}
var Ready bool = false

func Init(conf ServiceConfig) error {
	config = conf
	Ready = true
	fmt.Println("Initializing rcfile service: ", conf)
	return nil
}
