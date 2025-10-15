package main

import "C"
import (
	"alat/pkg/core/config"
	"alat/pkg/core/device"
	"alat/pkg/core/node"
	"alat/pkg/core/pair"
	"alat/pkg/core/service"
	"alat/pkg/core/storage"
	"fmt"
	"path"
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
func create_instance(configPath *C.char, deviceType *C.char, appConfigC *C.char, serviceConfigC *C.char) C.int {
	instancesMutex.Lock()
	defer instancesMutex.Unlock()
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()

	goConfigPath := C.GoString(configPath)
	goDeviceType := C.GoString(deviceType)

	var appSettings *config.AppSettings
	appSettings, alatError = config.LoadAppSettings(path.Join(goConfigPath, "settings.yml"))
	if alatError != nil {
		return -1
	}
	var serviceSettings *config.ServiceSettings
	serviceSettings, alatError = config.LoadServiceSettings(path.Join(goConfigPath, "services.yml"))
	if alatError != nil {
		return -2
	}

	storagePath := path.Join(goConfigPath, "node.yml")
	nodeStore := storage.CreateYAMLNodeStorage(storagePath)

	registry := initServices(serviceSettings)

	details := &device.Details{
		Color:       appSettings.DeviceColor,
		Name:        appSettings.DeviceName,
		Type:        device.DeviceTypeFromString(goDeviceType),
		Certificate: appSettings.Certificate,
	}

	var pairManager *pair.PairManager
	pairManager, alatError = pair.NewManager(nodeStore, details)
	if alatError != nil {
		return -3
	}

	node, err := node.NewNode(&registry, nodeStore, details, pairManager)
	if err != nil {
		alatError = err
		return -4
	}

	instance := &AlatInstance{
		node:             node,
		nodeStore:        nodeStore,
		serviceRegistery: &registry,
		appSettings:      appSettings,
		serviceSettings:  serviceSettings,
		configPath:       goConfigPath,
	}

	handle := nextInstanceID
	instances[handle] = instance
	nextInstanceID++

	return C.int(handle)
}

//export start_instance
func start_instance(handle C.int) C.int {
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	instance := getInstance(handle)
	if instance == nil {
		alatError = fmt.Errorf("instance %d cannot be started since it does not exist", handle)
		return -1
	}
	if alatError = instance.node.Start(); alatError != nil {
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
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	if instance := getInstance(handle); instance != nil {
		return instance.node.GetPort()
	} else {
		alatError = fmt.Errorf("could not get instance's port since it does not exist")
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
