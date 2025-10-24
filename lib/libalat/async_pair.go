package main

/*
#include "async_pair.h"
*/
import "C"

import (
	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"encoding/json"
	"sync"
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

		requestIDC := C.CString(requestID)
		tokenJSONC := C.CString(string(tokenBytes))
		detailsJSONC := C.CString(string(detailsBytes))
		C.call_async_callback_bridge(cb, handle, requestIDC, tokenJSONC, detailsJSONC)
	})

	return 0
}

//export submit_pair_response
func submit_pair_response(handle C.int, requestIDC *C.char, accepted C.bool, reasonC *C.char) C.int {
	instance := getInstance(handle)
	if instance == nil {
		return -1 // Instance not found
	}

	requestID := C.GoString(requestIDC)
	reason := C.GoString(reasonC)

	err := instance.node.SubmitPairResponse(requestID, bool(accepted), reason)
	if err != nil {
		setError(err)
		return -2 // Error submitting response
	}

	return 0
}
