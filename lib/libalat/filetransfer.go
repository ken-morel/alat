package main

import "fmt"
import "C"

//export get_file_transfers_status
func get_file_transfers_status(handle C.int) *C.char {
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	instance := getInstance(handle)
	if instance == nil {
		alatError = fmt.Errorf("Alat instance %d does not exist", handle)
		return nil
	}
	return toJSON(*instance.node.GetFileTransfersStatus())
}
