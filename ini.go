package banny

import "image/color"

var Config = Banny{
	Rows: 5,
	Foreground: []color.NRGBA{
		rgb(45, 79, 255),
		rgb(141, 69, 170),
		rgb(49, 203, 115),
	},
	Background: rgb(240, 240, 240),
}

func rgb(r, g, b uint8) color.NRGBA { return color.NRGBA{R: r, G: g, B: b, A: 255} }
