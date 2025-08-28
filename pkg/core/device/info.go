// Package device: holds peer related things
package device

import "alat/pkg/pbuf"

type Info struct {
	ID    string
	Name  string
	Color Color
	Type  DeviceType
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
	MobileDevice  DeviceType = "mobile"
	DesktopDevice DeviceType = "desktop"
	TVDevice      DeviceType = "tv"
	UnknownDevice DeviceType = "unknown"
)

func (t DeviceType) ToPBUF() pbuf.DeviceType {
	switch t {
	case MobileDevice:
		return pbuf.DeviceType_Mobile
	case DesktopDevice:
		return pbuf.DeviceType_Desktop
	case TVDevice:
		return pbuf.DeviceType_TV
	default:
		return pbuf.DeviceType_Unknown
	}
}

func PbufToColor(pbufColor *pbuf.Color) *Color {
	return &Color{
		uint8(pbufColor.GetR()), uint8(pbufColor.GetG()), uint8(pbufColor.GetB()),
	}
}

func PbufToDType(pbType *pbuf.DeviceType) DeviceType {
	switch pbType {
	case pbuf.DeviceType_Desktop.Enum():
		return DesktopDevice
	case pbuf.DeviceType_Mobile.Enum():
		return MobileDevice
	case pbuf.DeviceType_TV.Enum():
		return TVDevice
	default:
		return UnknownDevice
	}
}

func PbufToInfo(pbufInfo *pbuf.DeviceInfo) *Info {
	return &Info{
		ID:    pbufInfo.Id,
		Name:  pbufInfo.Name,
		Color: *PbufToColor(pbufInfo.GetColor()),
		Type:  PbufToDType(pbufInfo.GetType().Enum()),
	}
}
