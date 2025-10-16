package main

/*
#include "async_pair.h"
*/
import "C"

import (
	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"encoding/json"
	"fmt"
	"sync"
	"unsafe"
)

var (
	asyncCallbacks      = make(map[int]C.async_pair_request_callback)
	asyncCallbacksMutex = &sync.Mutex{}
)

func setAsyncCallback(handle int, callback C.async_pair_request_callback) {
	asyncCallbacksMutex.Lock()
	defer asyncCallbacksMutex.Unlock()
	asyncCallbacks[handle] = callback
}

func getAsyncCallback(handle int) C.async_pair_request_callback {
	asyncCallbacksMutex.Lock()
	defer asyncCallbacksMutex.Unlock()
	return asyncCallbacks[handle]
}

//export register_async_pair_request_callback
func register_async_pair_request_callback(handle C.int, callback C.async_pair_request_callback) C.int {
	instance := getInstance(handle)
	if instance == nil {
		return -1 // Instance not found
	}

	setAsyncCallback(int(handle), callback)

	instance.node.OnPairRequest(func(requestID string, token *security.PairToken, details *device.Details) {
		cb := getAsyncCallback(int(handle))
		if cb == nil {
			return
		}

		tokenBytes, err := json.Marshal(token)
		if err != nil {
			// Handle error appropriately
			return
		}
		detailsBytes, err := json.Marshal(details)
		if err != nil {
			// Handle error appropriately
			return
		}

		requestID_C := C.CString(requestID)
		tokenJSON_C := C.CString(string(tokenBytes))
		detailsJSON_C := C.CString(string(detailsBytes))
		fmt.Println("[libalat] received request, sending to dart handler")
		C.call_async_callback_bridge(cb, handle, requestID_C, tokenJSON_C, detailsJSON_C)
		fmt.Println("[libalat] handler called")
		C.free(unsafe.Pointer(detailsJSON_C))
		C.free(unsafe.Pointer(tokenJSON_C))
		C.free(unsafe.Pointer(requestID_C))
	})

	return 0
}

//export submit_pair_response
func submit_pair_response(handle C.int, requestID_C *C.char, accepted C.bool, reason_C *C.char) C.int {
	instance := getInstance(handle)
	if instance == nil {
		return -1 // Instance not found
	}

	requestID := C.GoString(requestID_C)
	reason := C.GoString(reason_C)

	err := instance.node.SubmitPairResponse(requestID, bool(accepted), reason)
	if err != nil {
		alatError = err
		return -2 // Error submitting response
	}

	return 0
}
