// Package libalat: Shared executable bindings for alat
package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"alat/pkg/core/config"
	"alat/pkg/core/device"
	"alat/pkg/core/node"
	"alat/pkg/core/pair"
	"alat/pkg/core/service"
	"alat/pkg/core/service/filesend"
	"alat/pkg/core/service/sysinfo"
	"alat/pkg/core/storage"
	"encoding/json"
	"path"
	"sync"
	"time"
	"unsafe"
)

// AlatInstance holds the entire state for a running instance of the Alat core.
type AlatInstance struct {
	node             *node.Node
	nodeStore        storage.NodeStorage
	serviceRegistery *service.Registry
	appSettings      *config.AppSettings
	serviceSettings  *config.ServiceSettings
	configPath       string
}

var (
	instances      = make(map[int]*AlatInstance)
	instancesMutex = &sync.Mutex{}
	nextInstanceID = 1
)

//export create_app
func create_app(configPath *C.char, deviceType C.int) C.int {
	instancesMutex.Lock()
	defer instancesMutex.Unlock()

	goConfigPath := C.GoString(configPath)

	// 1. Load Settings
	appSettings, err := config.LoadAppSettings(goConfigPath)
	if err != nil {
		return 0 // Error
	}
	serviceSettings, err := config.LoadServiceSettings(goConfigPath)
	if err != nil {
		return 0 // Error
	}

	// 2. Create Storage
	storagePath := path.Join(goConfigPath, "node.yml")
	nodeStore := storage.CreateYAMLNodeStorage(storagePath)

	// 3. Create Services
	registry := initServices(serviceSettings)

	// 4. Create Device Details
	details := &device.Details{
		Color:       appSettings.DeviceColor,
		Name:        appSettings.DeviceName,
		Type:        device.DeviceType(deviceType),
		Certificate: appSettings.Certificate,
	}

	// 5. Create Pair Manager
	pairManager, err := pair.NewManager(nodeStore, details)
	if err != nil {
		return 0 // Error
	}

	// 6. Create the Node
	node, err := node.NewNode(&registry, nodeStore, details, pairManager)
	if err != nil {
		return 0 // Error
	}

	// 7. Assemble the App and store it
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

// ... (start, stop, destroy functions remain the same) ...

//export get_app_settings_json
func get_app_settings_json(handle C.int) *C.char {
	instancesMutex.Lock()
	instance, ok := instances[int(handle)]
	instancesMutex.Unlock()

	if !ok {
		return nil
	}

	jsonBytes, err := json.Marshal(instance.appSettings)
	if err != nil {
		return nil
	}

	return C.CString(string(jsonBytes))
}

//export set_app_settings_json
func set_app_settings_json(handle C.int, settingsJson *C.char) C.int {
	instancesMutex.Lock()
	instance, ok := instances[int(handle)]
	instancesMutex.Unlock()

	if !ok {
		return -1 // Invalid handle
	}

	goJson := C.GoString(settingsJson)
	var newSettings config.AppSettings
	if err := json.Unmarshal([]byte(goJson), &newSettings); err != nil {
		return -2 // JSON parsing error
	}

	// Update and save
	instance.appSettings = &newSettings
	if err := config.SaveAppSettings(instance.appSettings, instance.configPath); err != nil {
		return -3 // File save error
	}

	// Propagate changes to the running node
	instance.node.SetDetails(&device.Details{
		Color:       instance.appSettings.DeviceColor,
		Name:        instance.appSettings.DeviceName,
		Type:        instance.node.PairManager.DeviceDetails().Type, // Preserve original type
		Certificate: instance.appSettings.Certificate,
	})

	return 0 // Success
}

func initServices(serviceSettings *config.ServiceSettings) service.Registry {
	return service.Registry{
		SysInfo: sysinfo.CreateService(sysinfo.Config{
			Enabled:   serviceSettings.SysInfo.Enabled,
			CacheTime: time.Duration(serviceSettings.SysInfo.CacheSeconds) * time.Second,
		}),
		FileSend: filesend.CreateService(filesend.Config{
			Enabled: true, // Assuming filesend is always enabled for libalat
		}),
	}
}

//export free_string
func free_string(s *C.char) {
	C.free(unsafe.Pointer(s))
}

func main() {}
