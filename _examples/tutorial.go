package main

import (
	tl "github.com/LtLi0n/termloop"
	tb "github.com/gdamore/tcell/v2/termbox"
)

type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.Entity.Draw(screen)
}

func (player *Player) Tick(event tb.Event) {
	if event.Type == tb.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tb.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tb.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tb.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tb.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
	}
}

func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}

func main() {
	game := tl.NewGame()
	game.Screen().SetFps(30)
	level := tl.NewBaseLevel(tb.Cell{
		Bg: tb.ColorGreen,
		Fg: tb.ColorBlack,
		Ch: 'v',
	})
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tb.ColorBlue))
	player := Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}
	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tb.Cell{Fg: tb.ColorRed, Ch: '옷'})
	level.AddEntity(&player)
	game.Screen().SetLevel(level)
	game.Start()
}
