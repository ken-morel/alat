// Package libalat: Shared executable bindings for alat
package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

//export get_node_status_json
func get_node_status_json(handle C.int) *C.char {
	instance := getInstance(handle)
	if instance == nil {
		return nil
	}
	return toJSON(instance.node.GetStatus())
}

//export free_string
func free_string(s *C.char) {
	C.free(unsafe.Pointer(s))
}

func main() {}
