package sysinfo

import (
	"alat/pkg/core/address"
	"alat/pkg/core/pbuf"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func GetInfo(address address.Address, token string) (*pbuf.SysInfo, error) {
	req, err := http.NewRequest("GET", "http://"+address.String()+"/sysinfo/get", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Pair-Token", token)
	req.Header.Set("Content-Type", "application/protobuf")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-ok status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var info pbuf.SysInfo
	if err := proto.Unmarshal(data, &info); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &info, nil
}
