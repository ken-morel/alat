// Package libalat: Shared executable bindings for alat
package main

import (
	"sync"

	"alat/pkg/core/storage"
)

/*
// You can include C definitions here if needed, but for now, it's empty.
#include <stdlib.h>
#include <string.h>
*/
import "C"

var nodeStorages = struct {
	Storages map[int]*storage.NodeStorage
	Count    int
	Mutex    *sync.Mutex
}{
	make(map[int]*storage.NodeStorage),
	0,
	&sync.Mutex{},
}

//export create_node_storage
func CreateNodeStorage(path *C.char) (nodeId C.int) {
	nodeStorages.Mutex.Lock()
	defer nodeStorages.Mutex.Unlock()
	nodeId = nodeStorages.Count
	nodeStorages.Storages[nodeId] = storage.CreateYAMLNodeStorage(C.GoString(path))
	nodeStorages.Count++
	return nodeId
}
