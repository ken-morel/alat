// Package client stores methods to query a server
package client

import (
	"alat/pkg/core/address"
	"alat/pkg/core/device"
	"fmt"
	"net"
)

func SearchDevices(channel chan<- device.DeviceInfo) error {
	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}
	for _, iface := range interfaces {
		// Skip loopback and down interfaces
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error getting addresses from interface", err.Error())
			continue
		}
		for _, rawAddr := range addrs {
			var ip net.IP

			switch v := rawAddr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue // Skip IPv6 addresses
			}

			if !ip.IsPrivate() {
				continue
			}
			for offset := range 10 {
				addr, err := address.NewAdderss(ip, uint16(offset+address.AlatPort))
				if err != nil {
					continue
				}
				found := false
				if !found {
					info, err := GetDeviceInfo(addr)
					if err != nil {
						fmt.Println("Error during info", err)
					} else {
						channel <- device.NewDeviceInfo(addr, &info)
					}
				}
			}

		}
	}
	close(channel)
	return nil
}
