package core

type DeviceAddress struct {
	Port int
	IP   string
}

type DeviceInfo struct {
	Name    string
	Address DeviceAddress
}
