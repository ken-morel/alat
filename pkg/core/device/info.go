// Package device: holds peer related things
package device

type DeviceInfo struct {
	ID   string
	Name string
	Type DeviceType
}

type DeviceType string

const (
	MobileDevice  DeviceType = "mobile"
	DesktopDevice DeviceType = "desktop"
	TVDevice      DeviceType = "tv"
)
