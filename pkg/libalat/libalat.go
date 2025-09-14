// Package libalat: Shared executable bindings for alat
package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"sync"
	"alat/pkg/core/storage"
)

var (
	storages      = make(map[int]*storage.NodeStorage)
	storagesMutex = &sync.Mutex{}
	nextStorageID = 1
)

//export create_node_storage
func create_node_storage(path *C.char) C.int {
	storagesMutex.Lock()
	defer storagesMutex.Unlock()

	goPath := C.GoString(path)

	// Assuming storage.CreateYAMLNodeStorage exists and returns *storage.NodeStorage
	nodeStorage := storage.CreateYAMLNodeStorage(goPath)

	handle := nextStorageID
	storages[handle] = nodeStorage
	nextStorageID++

	return C.int(handle)
}

//export destroy_node_storage
func destroy_node_storage(handle C.int) {
	storagesMutex.Lock()
	defer storagesMutex.Unlock()

	// We don't need to do anything to the object itself,
	// just remove it from the map. Go's garbage collector
	// will handle the rest once nothing references it anymore.
	delete(storages, int(handle))
}

// A main function is required to build a shared library.
func main() {}
