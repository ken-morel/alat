package main

import "C"
import (
	"encoding/json"
)

func toJSON(v any) *C.char {
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	bytes, alatError := json.Marshal(v)
	if alatError != nil {
		return nil
	}
	return C.CString(string(bytes))
}
