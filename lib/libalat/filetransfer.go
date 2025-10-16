package main

import "fmt"
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

func query_send_files_to_connected_device(handle C.int, deviceIdC *C.char, filesJsonC *C.char) C.int {

	instance := getInstance(handle)
	if instance == nil {
		return -1
	}
	deviceId := C.GoString(deviceIdC)

}
