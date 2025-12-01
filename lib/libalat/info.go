package main

import "C"
import (
	"alat/pkg/core/transport/client"
	"encoding/json"
	"fmt"
	"net"
)

//export query_device_info_json
func query_device_info_json(ipJsonC *C.char, portC C.int) *C.char {
	port := int(portC)
	ipJson := C.GoString(ipJsonC)
	var ip net.IP
	if err := json.Unmarshal([]byte(ipJson), &ip); err != nil {
		setError(fmt.Errorf("could not parse device ip: %v", err))
		return nil
	}
	info, err := client.GetDeviceInfo(ip, port)
	if err != nil {
		setError(fmt.Errorf("error getting device info: %v", err))
		return nil
	}
	return toJSON(info)

}
