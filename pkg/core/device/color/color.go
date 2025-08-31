// Package color stores easy identifiable colors
package color

import "alat/pkg/pbuf"

type Color struct {
	Name string
	Hex  string
	R    uint8
	G    uint8
	B    uint8
}

var Colors = []Color{
	{"Red", "#FF0000", 255, 0, 0},
	{"Blue", "#0000FF", 0, 0, 255},
	{"Green", "#008000", 0, 128, 0},
	{"Yellow", "#FFFF00", 255, 255, 0},
	{"Purple", "#800080", 128, 0, 128},
	{"Orange", "#FFA500", 255, 165, 0},
	{"Pink", "#FFC0CB", 255, 192, 203},
	{"Teal", "#008080", 0, 128, 128},
	{"Cyan", "#00FFFF", 0, 255, 255},
	{"Magenta", "#FF00FF", 255, 0, 255},
	{"Lime", "#00FF00", 0, 255, 0},
	{"Maroon", "#800000", 128, 0, 0},
	{"Navy", "#000080", 0, 0, 128},
	{"Olive", "#808000", 128, 128, 0},
	{"Gray", "#808080", 128, 128, 128},
	{"Silver", "#C0C0C0", 192, 192, 192},
	{"Black", "#000000", 0, 0, 0},
	{"White", "#FFFFFF", 255, 255, 255},
	{"Brown", "#A52A2A", 165, 42, 42},
	{"Gold", "#FFD700", 255, 215, 0},
}

var DefaultColor = &Colors[0]

func (c *Color) ToPBUF() *pbuf.Color {
	return &pbuf.Color{
		Name: c.Name,
	}
}

func FromPBUF(pb *pbuf.Color) *Color {
	for _, c := range Colors {
		if c.Name == pb.Name {
			return &c
		}
	}
	return DefaultColor
}

func FromString(nameOrHex string) *Color {
	for _, col := range Colors {
		if col.Name == nameOrHex || col.Hex == nameOrHex {
			return &col
		}
	}
	return DefaultColor
}
