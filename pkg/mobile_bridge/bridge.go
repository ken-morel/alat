// Package mobile_bridge will act as a bridge between Dart FFI and the Alat core.
package main

import (
	"C"
	"alat/pkg/core/client"
	"alat/pkg/core/device"
	"encoding/json"
)
import "fmt"

//export SearchDevices
func SearchDevices() *C.char {
	channel := make(chan device.DeviceInfo)
	var devices []device.DeviceInfo

	go client.SearchDevices(channel)

	for info := range channel {
		devices = append(devices, info)
	}

	jsonBytes, err := json.Marshal(devices)
	if err != nil {
		// In case of an error, return an empty JSON array
		return C.CString("[]")
	}

	return C.CString(string(jsonBytes))
}

func main() {
	fmt.Println("Main running..., I don't know why?")
}
