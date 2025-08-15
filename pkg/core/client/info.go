package client

import (
	"alat/pkg/core/address"
	"alat/pkg/core/pbuf"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func GetDeviceInfo(addr address.Address) (info pbuf.DeviceInfo, err error) {
	res, err := http.Get("http://" + addr.String() + "/alat-info")
	if err != nil {
		return
	}
	data := make([]byte, 1024) // 1kb, that's too much
	length, err := res.Body.Read(data)
	if err != nil {
		return
	}
	err = proto.Unmarshal(data[0:length], &info) // slice it, not copy!
	return
}
