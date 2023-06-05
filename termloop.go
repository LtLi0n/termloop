package termloop

import (
	"strings"

	tb "github.com/gdamore/tcell/v2/termbox"
)

// A Canvas is a 2D array of Cells, used for drawing.
// The structure of a Canvas is an array of columns.
// This is so it can be addressed canvas[x][y].
type Canvas [][]tb.Cell

// NewCanvas returns a new Canvas, with
// width and height defined by arguments.
func NewCanvas(width, height int) Canvas {
	canvas := make(Canvas, width)
	for i := range canvas {
		canvas[i] = make([]tb.Cell, height)
	}
	return canvas
}

func (canvas *Canvas) equals(oldCanvas *Canvas) bool {
	c := *canvas
	c2 := *oldCanvas
	if c2 == nil {
		return false
	}
	sz_c := len(c)
	sz_c2 := len(c2)
	if sz_c != sz_c2 {
		return false
	}
	// both arrays might be empty.
	if sz_c == 0 {
		return true
	}
	if len(c[0]) != len(c2[0]) {
		return false
	}
	for i := range c {
		for j := range c[i] {
			equal := CellsEqual(&c[i][j], &c2[i][j])
			if !equal {
				return false
			}
		}
	}
	return true
}

// CanvasFromString returns a new Canvas, built from
// the characters in the string str. Newline characters in
// the string are interpreted as a new Canvas row.
func CanvasFromString(str string) Canvas {
	lines := strings.Split(str, "\n")
	runes := make([][]rune, len(lines))
	width := 0
	for i := range lines {
		runes[i] = []rune(lines[i])
		width = max(width, len(runes[i]))
	}
	height := len(runes)
	canvas := make(Canvas, width)
	for i := 0; i < width; i++ {
		canvas[i] = make([]tb.Cell, height)
		for j := 0; j < height; j++ {
			if i < len(runes[j]) {
				canvas[i][j] = tb.Cell{Ch: runes[j][i]}
			}
		}
	}
	return canvas
}

// Drawable represents something that can be drawn, and placed in a Level.
type Drawable interface {
	Tick(tb.Event) // Method for processing events, e.g. input
	Draw(*Screen)  // Method for drawing to the screen
}

// Physical represents something that can collide with another
// Physical, but cannot process its own collisions.
// Optional addition to Drawable.
type Physical interface {
	Position() (int, int) // Return position, x and y
	Size() (int, int)     // Return width and height
}

// DynamicPhysical represents something that can process its own collisions.
// Implementing this is an optional addition to Drawable.
type DynamicPhysical interface {
	Position() (int, int) // Return position, x and y
	Size() (int, int)     // Return width and height
	Collide(Physical)     // Handle collisions with another Physical
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Abstract Termbox stuff for convenience - users
// should only need Termloop imported

func CellsEqual(c1 *tb.Cell, c2 *tb.Cell) bool {
	return c1.Fg == c2.Fg &&
		c1.Bg == c2.Bg &&
		c1.Ch == c2.Ch
}
