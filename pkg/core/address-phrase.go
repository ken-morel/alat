package core

import (
	"fmt"
	"net"
	"net/http"
)

type Address struct {
	Port   uint16
	IP     net.IP
	Phrase string
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

func AddressPhrase(ip net.IP, port uint16) (string, error) {
	ip = ip.To4()
	if ip == nil {
		return "", fmt.Errorf("address is not a valid IPv4 address")
	}

	// TODO: Take more than justs the last two bytes
	first := int(port - AlatPort)
	second := int(ip[2])
	third := int(ip[3])

	color := colors[first%len(colors)]
	adjective := adjectives[second%len(adjectives)]
	noun := nouns[third%len(nouns)]

	return fmt.Sprintf("%s-%s-%s", color, adjective, noun), nil
}

func NewAdderss(ip net.IP, port uint16) (Address, error) {
	phrase, err := AddressPhrase(ip, port)
	if err != nil {
		return Address{}, err
	}
	return Address{
		Port:   port,
		IP:     ip,
		Phrase: phrase,
	}, nil
}

func (addr *Address) Ping() (bool, error) {
	http.NewRequest("GET", "http://localhost", nil)
	return true, nil
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
				addr, err := NewAdderss(ip, uint16(offset+AlatPort))
				if err != nil {
					continue
				}
				found := false
				for _, a := range addresses {
					if a.Phrase == addr.Phrase {
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
