package banny

import (
	"image/color"
	"image"
	"net"
	"log"
	"strings"
	"crypto/sha1"
	"regexp"
	"fmt"
)

type Banny struct {
	Rows       int
	Foreground []color.NRGBA
	Background color.NRGBA
}

// Generate generates the identicon, default is .png.
func (banny *Banny) Generate(width int, s string) image.Image {
	var data []byte = banny.getRowData(s)
	// fmt.Println("data:", data)
	// Output:
	// data: [25 245 123 220 244 141 97 121 185 57 61 203 122 18 94 255 2 172 106 178]
	fg, bg := banny.colors(data[0])
	palette := color.Palette{bg, fg}
	img := image.NewPaletted(image.Rect(0, 0, width, width), palette)

	for _, rect := range banny.blocks(width, data[1:]) {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			for y := rect.Min.Y; y < rect.Max.Y; y++ {
				img.Pix[y*img.Stride+x] = 1
			}
		}
	}
	return img
}

func (banny *Banny) blocks(width int, data []byte) []image.Rectangle {
	blockWidth := width / (banny.Rows + 1)
	padding := blockWidth / 2
	num := banny.Rows * (banny.Rows/2 + banny.Rows%2)
	res := make([]image.Rectangle, 0, banny.Rows*banny.Rows)

	for i := 0; i < num; i++ {
		if !banny.fill(i, data) {
			continue
		}

		column := i / banny.Rows
		row := i % banny.Rows
		// fmt.Println("col and row:", column, row)
		// Output:
		// col and row: 0 0
		// ...

		pt := image.Pt(padding+column*blockWidth, padding+row*blockWidth)
		rect := image.Rectangle{pt, image.Pt(pt.X+blockWidth, pt.Y+blockWidth)}
		res = append(res, rect)

		if column < banny.Rows/2+banny.Rows%2-1 {
			rect.Min.X = padding + (banny.Rows-column-1)*blockWidth
			rect.Max.X = rect.Min.X + blockWidth
			res = append(res, rect)
		}
	}
	return res
}

func (banny *Banny) fill(block int, data []byte) bool {
	// NOTE: This method is quoted from sigil https://github.com/cupcake/sigil
	// needed be replaced in the future!
	if data[block/8]>>uint(8-((block%8)+1))&1 == 0 {
		return false
	}
	return true
}

// getRowData gets []byte which is hashed from ip address or email address.
func (banny *Banny) getRowData(s string) []byte {
	if b, _ := regexp.Match("^(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|[1-9])\\."+
		"(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)\\."+
		"(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)\\."+
		"(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)$", []byte(s)); b {

		fmt.Println("This is an ip address:", s)

		// hash ip address
		h := sha1.New()
		h.Write([]byte(s))
		return h.Sum(nil)
	} else if b, _ := regexp.Match("^([a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+)$", []byte(s)); b {

		fmt.Println("This is an email address:", s)

		// hash email address
		h := sha1.New()
		h.Write([]byte(s))
		return h.Sum(nil)
	} else {
		fmt.Println("Input is neither an ip address nor email address, use local ip address instead.")
		conn, err := net.Dial("udp", "google.com:80")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		// get ip like xxx.xxx.xxx.xxx
		ip := strings.Split(conn.LocalAddr().String(), ":")[0]

		// hash ip address
		h := sha1.New()
		h.Write([]byte(ip))
		return h.Sum(nil)
	}
}

// colors sets the foreground and background of the identicon.
func (banny *Banny) colors(b byte) (color.NRGBA, color.NRGBA) {
	fg := banny.Foreground[int(b)%len(banny.Foreground)]
	return fg, banny.Background
}
