package main

import "alat/pkg/core/config"
import "C"

//export default_app_config
func default_app_config() *C.char {
	return toJSON(config.DefaultAppConfig())
}

//export default_service_config
func default_service_config() *C.char {
	return toJSON(config.DefaultServiceConfig())
}
