package banny

import "image/color"

var Config = Banny{
	Rows: 5,
	Foreground: []color.NRGBA{
		rgb(0, 0, 255),		// blue
		rgb(160,32,240),	// purple
		rgb(154, 205,50),	// yellowgreen
		rgb(238, 130,238),	// violet
		rgb(255,99,71),		// tomato
		rgb(0,128,128),		// teal
		rgb(0,255,127),		// springgreen
		rgb(160,90,205),	// slateblue
		rgb(112, 128,144),	// slategray
		rgb(160,82,45),		// sienna
		rgb(128,128,0),		// olive
		rgb(210,105,30),	// chocolate
		rgb(220,20,60),		// crimson
	},
	Background: rgb(240, 240, 240),
}

func rgb(r, g, b uint8) color.NRGBA { return color.NRGBA{R: r, G: g, B: b, A: 255} }
