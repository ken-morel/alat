package main

import "C"
import (
	"alat/pkg/core/discovery"
	"encoding/json"
	"fmt"
)

//export discovery_provide_found_devices_json
func discovery_provide_found_devices_json(handle C.int, devicesJsonC *C.char) C.int {
	if instance := getInstance(handle); instance != nil {
		devicesJson := C.GoString(devicesJsonC)
		var devices []discovery.FoundDevice
		if err := json.Unmarshal([]byte(devicesJson), &devices); err != nil {
			setError(fmt.Errorf("error unmarshalling found devices json %v", err))
			return -2
		}
		instance.node.SetFoundDevices(devices)
		return 0
	} else {
		return -1
	}
}
