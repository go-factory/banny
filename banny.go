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
func (banny *Banny) Generate(width int) image.Image {
	var data []byte = banny.getRowData()
	bg, fg := banny.colors(data[0])
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
	blockWidth, padding := width/(banny.Rows+1), width/(banny.Rows+1)/2
	num := banny.Rows * (banny.Rows/2 + banny.Rows%2)
	res := make([]image.Rectangle, 0, banny.Rows*banny.Rows)

	for i := 0; i < num; i++ {

		column := i / banny.Rows
		row := i % banny.Rows

		if !fill(row, column, data) {
			continue
		}

		pt := image.Pt(padding+row*blockWidth, padding+column*blockWidth)
		rect := image.Rectangle{pt, image.Pt(pt.X+blockWidth, pt.Y+blockWidth)}
		res = append(res, rect)

		if column < banny.Rows/2+banny.Rows%2 {
			rect.Min.X = padding + (banny.Rows-column-1)*blockWidth
			rect.Max.X = rect.Min.X + blockWidth
			res = append(res, rect)
		}
	}
	return res
}

func fill(row, column int, data []byte) bool {
	if data[row/len(data)]>>uint(column%len(data))&1 == 0 {
		return true
	}
	return false
}

// getRowData gets []byte which is hashed from ip address.
func (banny *Banny) getRowData() []byte {
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
func (banny *Banny) colors(b byte) (color.NRGBA, color.NRGBA) {
	fg := banny.Foreground[int(b)%len(banny.Foreground)]
	return fg, banny.Background
}
