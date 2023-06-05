package main

import (
	"io/ioutil"

	tl "github.com/LtLi0n/termloop"
	tb "github.com/gdamore/tcell/v2/termbox"
)

type Player struct {
	*tl.Entity
}

func (p *Player) Tick(ev tb.Event) {
	// Enable arrow key movement
	if ev.Type == tb.EventKey {
		x, y := p.Position()
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
		p.SetPosition(x, y)
	}
}

// Here we define a parse function for reading a Player out of JSON.
func parsePlayer(data map[string]interface{}) tl.Drawable {
	e := tl.NewEntity(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		1, 1,
	)
	e.SetCell(0, 0, &tb.Cell{
		Ch: []rune(data["ch"].(string))[0],
		Fg: tb.Attribute(data["color"].(float64)),
	})
	return &Player{e}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	g := tl.NewGame()
	g.Screen().SetFps(30)
	l := tl.NewBaseLevel(tb.Cell{Bg: 76, Fg: 1})
	lmap, err := ioutil.ReadFile("level.json")
	checkErr(err)
	parsers := make(map[string]tl.EntityParser)
	parsers["Player"] = parsePlayer
	err = tl.LoadLevelFromMap(string(lmap), parsers, l)
	checkErr(err)
	g.Screen().SetLevel(l)
	g.Start()
}
