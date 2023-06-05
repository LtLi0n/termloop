package main

import (
	tl "github.com/LtLi0n/termloop"
	tb "github.com/gdamore/tcell/v2/termbox"
)

type CollRec struct {
	*tl.Rectangle
	move bool
	px   int
	py   int
}

func NewCollRec(x, y, w, h int, color tb.Attribute, move bool) *CollRec {
	return &CollRec{
		Rectangle: tl.NewRectangle(x, y, w, h, color),
		move:      move,
	}
}

func (r *CollRec) Tick(ev tb.Event) {
	// Enable arrow key movement
	if ev.Type == tb.EventKey && r.move {
		r.px, r.py = r.Position()
		switch ev.Key {
		case tb.KeyArrowRight:
			r.SetPosition(r.px+1, r.py)
		case tb.KeyArrowLeft:
			r.SetPosition(r.px-1, r.py)
		case tb.KeyArrowUp:
			r.SetPosition(r.px, r.py-1)
		case tb.KeyArrowDown:
			r.SetPosition(r.px, r.py+1)
		}
	}
}

func (r *CollRec) Collide(p tl.Physical) {
	// Check if it's a CollRec we're colliding with
	if _, ok := p.(*CollRec); ok && r.move {
		r.SetColor(tb.ColorBlue)
		r.SetPosition(r.px, r.py)
	}
}

func main() {
	g := tl.NewGame()
	g.Screen().SetFps(60)
	l := tl.NewBaseLevel(tb.Cell{
		Bg: tb.ColorWhite,
	})
	l.AddEntity(NewCollRec(3, 3, 3, 3, tb.ColorRed, true))
	l.AddEntity(NewCollRec(7, 4, 3, 3, tb.ColorGreen, false))
	g.Screen().SetLevel(l)
	g.Screen().AddEntity(tl.NewFpsText(0, 0, tb.ColorRed, tb.ColorDefault, 0.5))
	g.Start()
}
