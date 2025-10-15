package main

import "C"
import "sync"

var (
	alatErrorLock = &sync.Mutex{}
	alatError     error
)

//export get_error
func get_error() *C.char {
	var message string
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	if alatError == nil {
		message = "Unknown error"
	} else {
		message = alatError.Error()
	}
	return C.CString(message)
}
