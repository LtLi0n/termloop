package main

import (
	"fmt"
	"os"

	tl "github.com/LtLi0n/termloop"
	tb "github.com/gdamore/tcell/v2/termbox"
)

type MovingText struct {
	*tl.Text
}

func (m *MovingText) Tick(ev tb.Event) {
	// Enable arrow key movement
	if ev.Type == tb.EventKey {
		x, y := m.Position()
		switch ev.Key {
		case tb.KeyArrowRight:
			x += 1
		case tb.KeyArrowLeft:
			x -= 1
		case tb.KeyArrowUp:
			y -= 1
		case tb.KeyArrowDown:
			y += 1
		}
		m.SetPosition(x, y)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a string as first argument")
		return
	}
	g := tl.NewGame()
	g.Screen().SetFps(30)
	g.Screen().AddEntity(&MovingText{tl.NewText(0, 0, os.Args[1], tb.ColorWhite, tb.ColorBlue)})
	g.Start()
}
