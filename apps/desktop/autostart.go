package main

import (
	"os"
	"strings"
)

func isAutoStarted() bool {
	for _, arg := range os.Args {
		if strings.EqualFold(arg, "--autostart") {
			return true
		}
	}
	return false
}

func autostartCheck() {
	if !appConfig.Autostart && isAutoStarted() {
		os.Exit(0)
	}
}
