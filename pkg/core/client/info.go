package client

import (
	"alat/pkg/core/address"
	"alat/pkg/core/pbuf"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func GetDeviceInfo(addr address.Address) (info pbuf.DeviceInfo, err error) {
	var contentLength int
	res, err := http.Get("http://" + addr.String() + "/alat-info")
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		err = fmt.Errorf("invalid status code %d:%s", res.StatusCode, res.Status)
		return
	}
	_, err = fmt.Sscanf(res.Header.Get("Content-Length"), "%d", &contentLength)
	if err != nil {
		return
	} else if contentLength < 10 || contentLength > 500 {
		log.Println("Content length not in range, not parsing data")
		err = fmt.Errorf("invalid content length(%d) when getting info", contentLength)
		return
	}
	data := make([]byte, contentLength)
	_, err = res.Body.Read(data)
	if err != nil {
		return
	}
	err = proto.Unmarshal(data, &info)
	return
}
