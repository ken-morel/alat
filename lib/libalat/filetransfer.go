package main

import (
	"encoding/json"
	"fmt"
)
import "C"

//export get_file_transfers_status
func get_file_transfers_status(handle C.int) *C.char {
	instance := getInstance(handle)
	if instance == nil {
		setError(fmt.Errorf("Alat instance %d does not exist", handle))
		return nil
	}
	return toJSON(*instance.node.GetFileTransfersStatus())
}

//export query_send_files_to_connected_device
func query_send_files_to_connected_device(handle C.int, deviceIdC *C.char, filesJsonC *C.char) C.int {
	filesJson := C.GoString(filesJsonC)
	instance := getInstance(handle)
	if instance == nil {
		setError(fmt.Errorf("Instance %d does not exist", handle))
		return -1
	}
	deviceId := C.GoString(deviceIdC)

	device := instance.node.GetConnectedDeviceByID(deviceId)

	if device == nil {
		setError(fmt.Errorf("Device disconnected"))
		return -2
	}
	var files []string
	if err := json.Unmarshal([]byte(filesJson), &files); err != nil {
		setError(fmt.Errorf("Error decoding list of files to send from json: %v", err))
	}

	defer instance.node.QuerySendFiles(device, files)
	return 0
}
