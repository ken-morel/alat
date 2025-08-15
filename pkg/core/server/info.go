package server

import (
	"alat/pkg/core/pbuf"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func handleInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Alat-Device", "true")
	w.WriteHeader(200)
	msg, err := proto.Marshal(&pbuf.DeviceInfo{
		Code: config.DeviceCode,
		Name: config.DeviceName,
		Type: pbuf.DeviceInfo_DeviceType(config.DeviceType),
		Color: &pbuf.DeviceColor{
			R: uint32(config.DeviceColor.R),
			G: uint32(config.DeviceColor.G),
			B: uint32(config.DeviceColor.B),
		},
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		w.Write(msg)
	}
}
