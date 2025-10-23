package main

import "C"
import (
	"alat/pkg/core/config"
	"alat/pkg/core/node"
	"alat/pkg/core/storage"
	"encoding/json"
	"fmt"
	"sync"
)

// AlatInstance holds the entire state for a running instance of the Alat core.
type AlatInstance struct {
	node *node.Node
}

var (
	instances      = make(map[int]*AlatInstance)
	instancesMutex = &sync.Mutex{}
	nextInstanceID = 1
)

//export get_instances
func get_instances() *C.char {
	instancesMutex.Lock()
	defer instancesMutex.Unlock()
	keys := make([]int, 0, len(instances))
	for k := range instances {
		keys = append(keys, k)
	}
	return toJSON(keys)
}

//export create_instance
func create_instance(configPathC *C.char, appConfigC *C.char, serviceConfigC *C.char) C.int {
	instancesMutex.Lock()
	defer instancesMutex.Unlock()

	configPath := C.GoString(configPathC)
	appConfigBytes := []byte(C.GoString(appConfigC))
	serviceConfigBytes := []byte(C.GoString(serviceConfigC))

	var defaultAppConfig config.AppConfig
	err := json.Unmarshal(appConfigBytes, &defaultAppConfig)
	if err != nil {
		setError(err)
		return -1
	}

	var defaultServiceConfig config.ServiceConfig
	err = json.Unmarshal(serviceConfigBytes, &defaultServiceConfig)
	if err != nil {
		setError(err)
		return -2
	}

	store := storage.CreateYAMLNodeStorage(configPath, defaultAppConfig, defaultServiceConfig)

	node, err := node.CreateNode(store)
	if err != nil {
		setError(fmt.Errorf("Error creating node: %v", err))
		return -3
	}
	instance := &AlatInstance{
		node: node,
	}

	handle := nextInstanceID
	instances[handle] = instance
	nextInstanceID++

	return C.int(handle)
}

//export start_instance
func start_instance(handle C.int) C.int {
	instance := getInstance(handle)
	if instance == nil {
		setError(fmt.Errorf("instance %d cannot be started since it does not exist", handle))
		return -1
	}
	if err := instance.node.Start(); err != nil {
		setError(err)
		return -2
	}
	return 0
}

//export stop_instance
func stop_instance(handle C.int) {
	if instance := getInstance(handle); instance != nil {
		instance.node.Stop()
	}
}

//export get_port
func get_port(handle C.int) int {
	if instance := getInstance(handle); instance != nil {
		return instance.node.GetPort()
	} else {
		setError(fmt.Errorf("could not get instance's port since it does not exist"))
		return -1
	}
}

//export destroy_instance
func destroy_instance(handle C.int) {
	instancesMutex.Lock()
	defer instancesMutex.Unlock()
	if instance, ok := instances[int(handle)]; ok {
		instance.node.Stop()
		delete(instances, int(handle))

	}
}

func getInstance(handle C.int) *AlatInstance {
	instancesMutex.Lock()
	defer instancesMutex.Unlock()
	return instances[int(handle)]
}
