// Package pair: converts between internal and pbuf types
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/pbuf"
)

func pbufToDeviceInfo(pbufInfo *pbuf.DeviceInfo) *device.DeviceInfo {
	return &device.DeviceInfo{
		ID:   pbufInfo.Id,
		Name: pbufInfo.Name,
		OS:   pbufInfo.Os,
	}
}

func deviceInfoToPbuf(info *device.DeviceInfo) *pbuf.DeviceInfo {
	return &pbuf.DeviceInfo{
		Id:   info.ID,
		Name: info.Name,
		Os:   info.OS,
	}
}

func pbufToPairingSession(pbufSession *pbuf.PairingSession) *PairingSession {
	var token security.PairToken
	copy(token[:], pbufSession.Token)
	return &PairingSession{
		Initiator: *pbufToDeviceInfo(pbufSession.Initiator),
		Responder: *pbufToDeviceInfo(pbufSession.Responder),
		Token:     token,
	}
}

func pairingSessionToPbuf(session *PairingSession) *pbuf.PairingSession {
	return &pbuf.PairingSession{
		Initiator: deviceInfoToPbuf(&session.Initiator),
		Responder: deviceInfoToPbuf(&session.Responder),
		Token:     session.Token[:],
	}
}

func pbufToPairedDevice(pbufDevice *pbuf.PairedDevice) *device.PairedDevice {
	var cert security.Certificate
	copy(cert[:], pbufDevice.Certificate)
	var token security.PairToken
	copy(token[:], pbufDevice.Token)
	return &device.PairedDevice{
		Certificate: cert,
		Token:       token,
	}
}

func pairedDeviceToPbuf(d *device.PairedDevice) *pbuf.PairedDevice {
	return &pbuf.PairedDevice{
		Certificate: d.Certificate[:],
		Token:       d.Token[:],
	}
}
