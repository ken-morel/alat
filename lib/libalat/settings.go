package main

import "C"

import (
	"alat/pkg/core/config"
	"alat/pkg/core/device"
	"alat/pkg/core/device/color"
	"alat/pkg/core/service/sysinfo"
	"encoding/json"
	"fmt"
	"path"
	"time"
)

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
