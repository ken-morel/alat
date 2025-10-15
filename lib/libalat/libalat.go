// Package libalat: Shared executable bindings for alat
package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"alat/pkg/core/config"
	"alat/pkg/core/device"
	"alat/pkg/core/device/color"
	"alat/pkg/core/node"
	"alat/pkg/core/pair"
	"alat/pkg/core/service"
	"alat/pkg/core/service/filesend"
	"alat/pkg/core/service/sysinfo"
	"alat/pkg/core/storage"
	"encoding/json"
	"fmt"
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
	alatErrorLock  = &sync.Mutex{}
	alatError      error
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

//export get_error
func get_error() *C.char {
	var message string
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	if alatError == nil {
		message = "Unknown error"
	} else {
		message = alatError.Error()
	}
	return C.CString(message)
}

// --- Lifecycle --- //

//export create_instance
func create_instance(configPath *C.char, deviceType *C.char) C.int {
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

// --- Settings --- //

//export get_app_settings_json
func get_app_settings_json(handle C.int) *C.char {
	instance := getInstance(handle)
	if instance == nil {
		return nil
	}
	return toJSON(instance.appSettings)
}

//export set_app_settings_json
func set_app_settings_json(handle C.int, settingsJSON *C.char) C.int {
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	instance := getInstance(handle)
	if instance == nil {
		return -1
	}

	var newSettings config.AppSettings
	if alatError = json.Unmarshal([]byte(C.GoString(settingsJSON)), &newSettings); alatError != nil {
		return -2
	}

	instance.appSettings = &newSettings
	if alatError = config.SaveAppSettings(instance.appSettings, path.Join(instance.configPath, "settings.yml")); alatError != nil {
		return -3
	}

	// Propagate changes to the running node
	instance.node.SetDetails(&device.Details{
		Color:       instance.appSettings.DeviceColor,
		Name:        instance.appSettings.DeviceName,
		Type:        instance.node.PairManager.DeviceDetails().Type, // Preserve original type
		Certificate: instance.appSettings.Certificate,
	})
	return 0
}

//export get_service_settings_json
func get_service_settings_json(handle C.int) *C.char {
	instance := getInstance(handle)
	if instance == nil {
		return nil
	}
	return toJSON(instance.serviceSettings)
}

//export set_service_settings_json
func set_service_settings_json(handle C.int, settingsJSON *C.char) C.int {
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	instance := getInstance(handle)
	if instance == nil {
		alatError = fmt.Errorf("instance %d does not exist, settings cannot be saved", handle)
		return -1
	}

	var newSettings config.ServiceSettings
	if alatError := json.Unmarshal([]byte(C.GoString(settingsJSON)), &newSettings); alatError != nil {
		return -2
	}

	instance.serviceSettings = &newSettings
	if alatError := config.SaveServiceSettings(instance.serviceSettings, path.Join(instance.configPath, "services.yml")); alatError != nil {
		return -3
	}

	instance.serviceRegistery.SysInfo.Configure(sysinfo.Config{
		Enabled:   newSettings.SysInfo.Enabled,
		CacheTime: time.Duration(newSettings.SysInfo.CacheSeconds) * time.Second,
	})
	return 0
}

// --- Device & Pairing --- //

//export get_found_devices_json
func get_found_devices_json(handle C.int) *C.char {
	instance := getInstance(handle)
	if instance == nil {
		return nil
	}
	return toJSON(instance.node.GetDiscoverer().GetFoundDevices())
}

//export get_paired_devices_json
func get_paired_devices_json(handle C.int) *C.char {
	instance := getInstance(handle)
	if instance == nil {
		return nil
	}
	paired, err := instance.nodeStore.GetPaired()
	if err != nil {
		return nil // Or return JSON error
	}
	return toJSON(paired)
}

//export get_connected_devices_json
func get_connected_devices_json(handle C.int) *C.char {
	instance := getInstance(handle)
	if instance == nil {
		return nil
	}
	return toJSON(instance.node.Connected.GetConnectedDevices())
}

// --- Core Constants --- //

//export get_alat_device_colors_json
func get_alat_device_colors_json() *C.char {
	return toJSON(color.Colors)
}

// --- Status --- //

//export get_node_status_json
func get_node_status_json(handle C.int) *C.char {
	instance := getInstance(handle)
	if instance == nil {
		return nil
	}
	return toJSON(instance.node.GetStatus())
}

// --- Helpers --- //

func getInstance(handle C.int) *AlatInstance {
	instancesMutex.Lock()
	defer instancesMutex.Unlock()
	return instances[int(handle)]
}

func toJSON(v any) *C.char {
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	bytes, alatError := json.Marshal(v)
	if alatError != nil {
		return nil
	}
	return C.CString(string(bytes))
}

func initServices(serviceSettings *config.ServiceSettings) service.Registry {
	return service.Registry{
		SysInfo: sysinfo.CreateService(sysinfo.Config{
			Enabled:   serviceSettings.SysInfo.Enabled,
			CacheTime: time.Duration(serviceSettings.SysInfo.CacheSeconds) * time.Second,
		}),
		FileSend: filesend.CreateService(filesend.Config{
			Enabled: true,
		}),
	}
}

//export free_string
func free_string(s *C.char) {
	C.free(unsafe.Pointer(s))
}

func main() {}
