package main

import "C"

import (
	"fmt"
	"sync"
)

var (
	_alatErrorLock = &sync.Mutex{}
	_alatError     error
)

//export get_error
func get_error() *C.char {
	var message string
	_alatErrorLock.Lock()
	defer _alatErrorLock.Unlock()
	if _alatError == nil {
		message = "unknown error"
	} else {
		message = _alatError.Error()
	}
	return C.CString(message)
}

func setError(err error) {
	_alatErrorLock.Lock()
	_alatError = err
	_alatErrorLock.Unlock()
}

func noSuchInstance(handle C.int) {
	setError(fmt.Errorf("alat instance id %d does not exist", handle))
}
