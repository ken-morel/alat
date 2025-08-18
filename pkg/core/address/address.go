// Package address stores address specific cod3
package address

import (
	"fmt"
	"net"
)

type Address struct {
	Port   uint16 `yaml:"port"`
	IP     net.IP `yaml:"ip"`
	Phrase string `yaml:"phrase"`
}

// AlatPort       ALAT
const AlatPort = 52529 // change 8 -> 8

var colors = []string{
	"red", "green", "blue", "orange", "indigo",
	"purple", "yellow", "magenta", "cyan",
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
	first := int(port-AlatPort) - 1
	second := int(ip[2])
	third := int(ip[3])

	color := ""
	if first >= 0 {
		color = colors[first%len(colors)] + "-"
	}
	adjective := adjectives[second%len(adjectives)]
	noun := nouns[third%len(nouns)]

	return fmt.Sprintf("%s%s-%s", color, adjective, noun), nil
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

func (addr *Address) String() string {
	return fmt.Sprintf("%s:%d", addr.IP.String(), addr.Port)
}

func GetThisAddress() (Address, error) {
	return NewAdderss(net.IPv4(192, 168, 1, 192), AlatPort)
}
