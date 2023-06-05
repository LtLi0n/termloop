package termloop

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	tb "github.com/gdamore/tcell/v2/termbox"
)

// Image processing

func colorGridFromFile(filename string) *[][]tb.Attribute {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	// Pull pixel colour data out of image
	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y
	colors := make([][]tb.Attribute, w)
	for i := range colors {
		colors[i] = make([]tb.Attribute, h)
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			if a < 1 {
				continue
			}
			R := int(r >> 8)
			G := int(g >> 8)
			B := int(b >> 8)
			colors[x-bounds.Min.X][y-bounds.Min.Y] = RgbTo256Color(R, G, B)
		}
	}
	return &colors
}

// BackgroundCanvasFromFile takes a path to an image file,
// and creates a canvas of background-only Cells representing
// the image. This can be applied to an Entity with ApplyCanvas.
func BackgroundCanvasFromFile(filename string) *Canvas {
	colors := colorGridFromFile(filename)
	c := make(Canvas, len(*colors))
	for i := range c {
		c[i] = make([]tb.Cell, len((*colors)[i]))
		for j := range c[i] {
			c[i][j] = tb.Cell{Bg: (*colors)[i][j]}
		}
	}
	return &c
}

// ForegroundCanvasFromFile takes a path to an image file,
// and creates a canvas of foreground-only Cells representing
// the image. This can be applied to an Entity with ApplyCanvas.
func ForegroundCanvasFromFile(filename string) *Canvas {
	colors := colorGridFromFile(filename)
	c := make(Canvas, len(*colors))
	for i := range c {
		c[i] = make([]tb.Cell, len((*colors)[i]))
		for j := range c[i] {
			c[i][j] = tb.Cell{Fg: (*colors)[i][j]}
		}
	}
	return &c
}
