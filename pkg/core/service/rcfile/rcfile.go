// Package rcfile: Stores api and cliend info for the sendfile service
package rcfile

import "fmt"

type ServiceConfig struct {
	Enabled     bool   `yaml:"enabled"`
	Destination string `yaml:"destination"`
	FileMaxSize uint64 `yaml:"filemaxsize"`
}

type FileReceiveHandler func(chan *ReceiveFileStatus, string, int)

type ReceiveFileStatus struct {
	Finished bool
	Error    error
	Percent  float32
}

const ServiceName = "rcfile"

var config ServiceConfig = ServiceConfig{
	Enabled: false,
}
var Ready bool = false

var recvHandle FileReceiveHandler

func Init(conf ServiceConfig, handle FileReceiveHandler) error {
	config = conf
	Ready = true
	fmt.Println("Initializing rcfile service: ", conf)
	recvHandle = handle
	return nil
}
