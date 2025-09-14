package main

import "C"

import (
	"sync"

	"alat/pkg/core/storage"
)

var (
	storages      = make(map[int]*storage.YAMLNodeStorage)
	storagesMutex = &sync.Mutex{}
	nextStorageID = 1
)

//export create_node_storage
func create_node_storage(path *C.char) C.int {
	storagesMutex.Lock()
	defer storagesMutex.Unlock()

	goPath := C.GoString(path)
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
	delete(storages, int(handle))
}
