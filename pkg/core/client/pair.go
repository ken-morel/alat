package client

import (
	"alat/pkg/core/address"
	"alat/pkg/core/device"
	"alat/pkg/core/pbuf"
	"alat/pkg/core/service"
	"bytes"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"
)

func SendPairRequest(addr address.Address, token string, services []service.Service) (rcv bool, err error) {
	url := "http://" + addr.String() + "/pair/request"
	info, err := device.ThisDeviceInfo.ToPBuf()
	if err != nil {
		return
	}
	var pbservices []*pbuf.Service
	for _, srv := range services {
		pb := srv.ToPBuf()
		pbservices = append(pbservices, &pb)
	}
	data, err := proto.Marshal(&pbuf.PairRequest{
		Token:    token,
		Device:   &info,
		Services: pbservices,
	})
	if err != nil {
		return
	}
	buffer := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		return
	}
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(data)))

	client := http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		rcv = true
	} else {
		rcv = false
	}

	return
}
