package main

import "C"
import "fmt"

//export query_connected_device_sysinfo
func query_connected_device_sysinfo(handle C.int, deviceIDC *C.char) *C.char {
	deviceID := C.GoString(deviceIDC)
	instance := getInstance(handle)
	if instance == nil {
		setError(fmt.Errorf("instance %d does not exist", handle))
		return nil
	}
	device := instance.node.GetConnectedDeviceByID(deviceID)
	if device == nil {
		setError(fmt.Errorf("connected device not found, device surely not connected"))
		return nil
	}

	info, err := instance.node.QueryDeviceSysInfo(device)
	if err != nil {
		setError(fmt.Errorf("error getting system information from %s: %s", device.Info.Name, err.Error()))
		return nil
	}
	return toJSON(info)
}
