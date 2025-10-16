package main

import "C"
import "fmt"

//export query_connected_device_sysinfo
func query_connected_device_sysinfo(handle C.int, deviceIdC *C.char) *C.char {
	deviceId := C.GoString(deviceIdC)
	instance := getInstance(handle)
	if instance == nil {
		setError(fmt.Errorf("Instance %d does not exist", handle))
		return nil
	}
	device := instance.node.GetConnectedDeviceByID(deviceId)
	if device == nil {
		setError(fmt.Errorf("Connected device not found, device surely not connected"))
		return nil
	}
	info, err := device.GetSysInfo()
	if err != nil {
		setError(fmt.Errorf("Error getting system information from %s: %s", device.Info.Name, err.Error()))
		return nil
	}
	return toJSON(info)
}
