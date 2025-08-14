package core

import (
	"fmt"
	"net"
)

type Address struct {
	Port uint16
	IP   net.IP
}

const AlatPort = 60000

var colors = []string{
	"red", "green", "blue", "orange", "indigo",
	"purple", "vermillion", "yellow", "magenta", "cyan",
}

var adjectives = []string{
	"happy", "silly", "bouncy", "quick", "quiet", "sunny", "windy", "jolly",
	"breezy", "clever", "daring", "fancy", "gentle", "lucky", "proud", "witty",
	"brave", "calm", "eager", "fiery", "grand", "humble", "lively", "merry",
	"noble", "placid", "regal", "sharp", "trusty", "vivid", "zesty", "amber",
}

var nouns = []string{
	"fox", "cat", "dog", "bird", "lion", "tiger", "bear", "wolf",
	"river", "ocean", "mountain", "forest", "meadow", "stream", "comet", "star",
	"planet", "moon", "galaxy", "nebula", "canyon", "valley", "island", "harbor",
	"castle", "tower", "bridge", "beacon", "shield", "sword", "arrow", "quill",
}

func (addr *Address) Phrase() (string, error) {
	ip := addr.IP.To4()
	if ip == nil {
		return "", fmt.Errorf("address is not a valid IPv4 address")
	}

	// An IPv4 address is 4 bytes. We use the last two, which are most likely
	// to be unique on a local network.
	first := int(addr.Port - AlatPort)
	second := int(ip[2])
	third := int(ip[3])

	color := colors[first%len(colors)]
	adjective := adjectives[second%len(adjectives)]
	noun := nouns[third%len(nouns)]

	return fmt.Sprintf("%s-%s-%s", color, adjective, noun), nil
}

func NewAdderss(ip net.IP, port uint16) Address {
	return Address{
		Port: port,
		IP:   ip,
	}
}

func GetLocalAddresses() ([]Address, error) {
	var addresses []Address
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range interfaces {
		// Skip loopback and down interfaces
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error getting interfaces addresses", err.Error())
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
				addr := NewAdderss(ip, uint16(offset+AlatPort))
				phrase, err := addr.Phrase()
				if err != nil {
					continue
				}
				found := false
				for _, a := range addresses {
					if p, e := a.Phrase(); e != nil || p == phrase {
						found = true
						break
					}
				}
				if !found {
					addresses = append(addresses, addr)
				}
			}

		}
	}

	if len(addresses) == 0 {
		return nil, fmt.Errorf("no suitable local network addresses found")
	}
	return addresses, nil
}
