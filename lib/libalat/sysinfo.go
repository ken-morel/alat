package main

import "C"
import "fmt"

//export query_connected_device_sysinfo
func query_connected_device_sysinfo(handle C.int, deviceIdC *C.char) *C.char {
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	deviceId := C.GoString(deviceIdC)
	instance := getInstance(handle)
	if instance == nil {
		alatError = fmt.Errorf("Instance %d does not exist", handle)
		return nil
	}
	device := instance.node.GetConnectedDeviceByID(deviceId)
	if device == nil {
		alatError = fmt.Errorf("Connected device not found, device surely not connected")
		return nil
	}
	info, err := device.GetSysInfo()
	if err != nil {
		alatError = fmt.Errorf("Error getting system information from %s: %s", device.Info.Name, err.Error())
		return nil
	}
	return toJSON(info)
}
