// Package device: holds peer related things
package device

import (
	"alat/pkg/core/device/color"
	"alat/pkg/pbuf"
)

type Info struct {
	ID    string      `yaml:"id"    json:"id"`
	Name  string      `yaml:"name"  json:"name"`
	Color color.Color `yaml:"color" json:"color"`
	Type  DeviceType  `yaml:"type"  json:"type"`
}

func (info *Info) ToPBUF() *pbuf.DeviceInfo {
	return &pbuf.DeviceInfo{
		Id:    info.ID,
		Name:  info.Name,
		Type:  info.Type.ToPBUF(),
		Color: info.Color.ToPBUF(),
	}
}

type DeviceType string

const (
	MobileDevice      DeviceType = "mobile"
	DesktopDevice     DeviceType = "desktop"
	TVDevice          DeviceType = "tv"
	WebDevice         DeviceType = "web"
	ArduinoDevice     DeviceType = "arduino"
	UnspecifiedDevice DeviceType = "unspecified"
)

func (t DeviceType) ToPBUF() pbuf.DeviceType {
	switch t {
	case MobileDevice:
		return pbuf.DeviceType_DEVICE_TYPE_MOBILE
	case DesktopDevice:
		return pbuf.DeviceType_DEVICE_TYPE_DESKTOP
	case TVDevice:
		return pbuf.DeviceType_DEVICE_TYPE_TV
	default:
		return pbuf.DeviceType_DEVICE_TYPE_UNSPECIFIED
	}
}

func PbufToColor(pbufColor *pbuf.Color) color.Color {
	return color.FromPBUF(pbufColor)
}

func PbufToDType(pbType pbuf.DeviceType) DeviceType {
	switch pbType {
	case pbuf.DeviceType_DEVICE_TYPE_DESKTOP:
		return DesktopDevice
	case pbuf.DeviceType_DEVICE_TYPE_MOBILE:
		return MobileDevice
	case pbuf.DeviceType_DEVICE_TYPE_TV:
		return TVDevice
	default:
		return UnspecifiedDevice
	}
}

func PbufToInfo(pbufInfo *pbuf.DeviceInfo) *Info {
	return &Info{
		ID:    pbufInfo.Id,
		Name:  pbufInfo.Name,
		Color: PbufToColor(pbufInfo.GetColor()),
		Type:  PbufToDType(pbufInfo.GetType()),
	}
}
