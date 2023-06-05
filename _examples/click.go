package main

import (
	"fmt"

	tl "github.com/LtLi0n/termloop"
	tb "github.com/gdamore/tcell/v2/termbox"
)

type EventInfo struct {
	*tl.Text
}

func NewEventInfo(x, y int) *EventInfo {
	return &EventInfo{tl.NewText(x, y, "Click somewhere", tb.ColorWhite, tb.ColorBlack)}
}

func (info *EventInfo) Tick(ev tb.Event) {
	if ev.Type != tb.EventMouse {
		return
	}
	var name string
	switch ev.Key {
	case tb.MouseLeft:
		name = "Mouse Left"
	case tb.MouseMiddle:
		name = "Mouse Middle"
	case tb.MouseRight:
		name = "Mouse Right"
	case tb.MouseWheelUp:
		name = "Mouse Wheel Up"
	case tb.MouseWheelDown:
		name = "Mouse Wheel Down"
	case tb.MouseRelease:
		name = "Mouse Release"
	default:
		name = fmt.Sprintf("Unknown Key (%#x)", ev.Key)
	}
	info.SetText(fmt.Sprintf("%s @ [%d, %d]", name, ev.MouseX, ev.MouseY))
}

type Clickable struct {
	*tl.Rectangle
}

func NewClickable(x, y, w, h int, col tb.Attribute) *Clickable {
	return &Clickable{tl.NewRectangle(x, y, w, h, col)}
}

func (c *Clickable) Tick(ev tb.Event) {
	x, y := c.Position()
	if ev.Type == tb.EventNone || ev.Type == tb.EventResize {
		return
	}
	if ev.Type != tb.EventMouse {
		return
	}
	if ev.Type == tb.EventMouse && ev.MouseX == x && ev.MouseY == y {
		if c.Color() == tb.ColorWhite {
			c.SetColor(tb.ColorBlack)
		} else {
			c.SetColor(tb.ColorWhite)
		}
	}
}

func main() {
	g := tl.NewGame()
	g.Screen().SetFps(60)
	g.Screen().AddEntity(NewEventInfo(0, 0))
	for i := 0; i < 40; i++ {
		for j := 1; j < 20; j++ {
			g.Screen().AddEntity(NewClickable(i, j, 1, 1, tb.ColorWhite))
		}
	}

	g.Start()
}
