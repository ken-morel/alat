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
		setError(fmt.Errorf("alat instance %d does not exist", handle))
		return nil
	}
	return toJSON(*instance.node.GetFileTransfersStatus())
}

//export query_send_files_to_connected_device
func query_send_files_to_connected_device(handle C.int, deviceIDC *C.char, filesJSONC *C.char) C.int {
	filesJSON := C.GoString(filesJSONC)
	instance := getInstance(handle)
	if instance == nil {
		setError(fmt.Errorf("instance %d does not exist", handle))
		return -1
	}
	deviceID := C.GoString(deviceIDC)

	device := instance.node.GetConnectedDeviceByID(deviceID)

	if device == nil {
		setError(fmt.Errorf("device disconnected"))
		return -2
	}
	var files []string
	if err := json.Unmarshal([]byte(filesJSON), &files); err != nil {
		setError(fmt.Errorf("error decoding list of files to send from json: %v", err))
	}

	defer instance.node.QuerySendFiles(device, files)
	return 0
}
