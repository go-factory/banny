package banny

import (
	"image/color"
	"net"
	"log"
	"strings"
	"crypto/sha1"
	"image"
)

type Banny struct {
	Rows       int
	Foreground []color.NRGBA
	Background color.NRGBA
}

// Generate generates the identicon, default is .png.
func (b *Banny) Generate(width int) image.Image {
	var data []byte = b.getRowData()
	bg, fg := b.colors(data[0])
	palette := color.Palette{bg, fg}
	img := image.NewPaletted(image.Rect(0, 0, width, width), palette)

	// TODO: divide indenticon into small cells and decide whether cells will be filled or not.

	return img
}

// getRowData gets []byte which is hashed from ip address.
func (b *Banny) getRowData() []byte {
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

// colors sets the foreground and background of the identicon.
func (b *Banny) colors(by byte) (color.NRGBA, color.NRGBA) {
	fg := b.Foreground[int(by)%len(b.Foreground)]
	return fg, b.Background
}
