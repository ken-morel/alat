package main

import "C"

import (
	"alat/pkg/core/config"
	"alat/pkg/core/device/color"
	"encoding/json"
	"fmt"
)

//export get_app_config_json
func get_app_config_json(handle C.int) *C.char {
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	instance := getInstance(handle)
	if instance == nil {
		return nil
	}
	config, err := instance.node.GetAppConfig()
	if err != nil {
		alatError = err
		return nil
	}
	return toJSON(config)
}

//export set_app_config_json
func set_app_config_json(handle C.int, settingsJSON *C.char) C.int {
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	instance := getInstance(handle)
	if instance == nil {
		return -1
	}

	var newSettings config.AppConfig
	if alatError = json.Unmarshal([]byte(C.GoString(settingsJSON)), &newSettings); alatError != nil {
		return -2
	}
	alatError := instance.node.SetAppConfig(newSettings)
	if alatError != nil {
		return -1
	}

	return 0
}

//export get_service_config_json
func get_service_config_json(handle C.int) *C.char {
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	instance := getInstance(handle)
	if instance == nil {
		return nil
	}
	conf, alatError := instance.node.GetServiceConfig()
	if alatError != nil {
		return nil
	}
	return toJSON(conf)
}

//export set_service_config_json
func set_service_config_json(handle C.int, settingsJSON *C.char) C.int {
	alatErrorLock.Lock()
	defer alatErrorLock.Unlock()
	instance := getInstance(handle)
	if instance == nil {
		alatError = fmt.Errorf("instance %d does not exist, settings cannot be saved", handle)
		return -1
	}

	var newSettings config.ServiceConfig
	if alatError := json.Unmarshal([]byte(C.GoString(settingsJSON)), &newSettings); alatError != nil {
		return -1
	}

	alatError = instance.node.SetServiceConfig(newSettings)
	if alatError != nil {
		return -1
	}
	return 0
}

// --- Device & Pairing --- //

//export get_found_devices_json
func get_found_devices_json(handle C.int) *C.char {
	instance := getInstance(handle)
	if instance == nil {
		return nil
	}
	return toJSON(instance.node.GetFoundDevices())
}

//export get_paired_devices_json
func get_paired_devices_json(handle C.int) *C.char {
	instance := getInstance(handle)
	if instance == nil {
		return nil
	}

	return toJSON(instance.node.GetPairedDevices())
}

//export get_connected_devices_json
func get_connected_devices_json(handle C.int) *C.char {
	instance := getInstance(handle)
	if instance == nil {
		return nil
	}
	return toJSON(instance.node.GetConnectedDevices())
}

//export get_alat_device_colors_json
func get_alat_device_colors_json() *C.char {
	return toJSON(color.Colors)
}
