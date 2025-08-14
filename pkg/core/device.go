package core

type DeviceInfo struct {
	Name    string
	Address Address
}

func GetAvailableAddresses() ([]Address, error) {
	return GetLocalAddresses()
}
