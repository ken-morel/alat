package main

import "C"
import (
	"alat/pkg/core/pair"
	"fmt"
)

type RequestPairFoundDeviceResponse struct {
	Status   int                 `json:"status"`
	Error    string              `json:"error"`
	Accepted bool                `json:"accepted"`
	Reason   pair.ResponseReason `json:"reason"`
}

//export request_pair_found_device
func request_pair_found_device(handleC C.int, deviceIdC *C.char) *C.char {
	instance := getInstance(handleC)
	deviceId := C.GoString(deviceIdC)
	if instance == nil {
		setError(fmt.Errorf("Handle %d not found", handleC))
		return toJSON(RequestPairFoundDeviceResponse{
			Status:   -1,
			Error:    "Handle not found",
			Accepted: false,
			Reason:   pair.ResponseReasonUnknown,
		})
	}
	res, err := instance.node.RequestPairFoundDevice(deviceId)
	if err != nil {
		setError(err)

		return toJSON(RequestPairFoundDeviceResponse{
			Status:   -5,
			Error:    err.Error(),
			Accepted: false,
			Reason:   pair.ResponseReasonUnknown,
		})
	} else {
		return toJSON(RequestPairFoundDeviceResponse{
			Status:   0,
			Error:    "",
			Accepted: res.GetAccepted(),
			Reason:   pair.ResponseReason(res.GetReason()),
		})
	}
}
