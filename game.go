package termloop

import (
	"fmt"
	"time"

	tb "github.com/gdamore/tcell/v2/termbox"
)

// Represents a top-level Termloop application.
type Game struct {
	screen *Screen
	input  *input
	debug  bool
	logs   []string
}

// NewGame creates a new Game, along with a Screen and input handler.
// Returns a pointer to the new Game.
func NewGame() *Game {
	g := Game{
		screen: NewScreen(),
		input:  newInput(),
		logs:   make([]string, 0),
	}
	return &g
}

// Screen returns the current Screen of a Game.
func (g *Game) Screen() *Screen {
	return g.screen
}

// SetScreen sets the current Screen of a Game.
func (g *Game) SetScreen(s *Screen) {
	g.screen = s
	g.screen.resize(tb.Size())
}

// DebugOn returns a bool showing whether or not debug mode is on.
func (g *Game) DebugOn() bool {
	return g.debug
}

// SetDebugOn sets debug mode's on status to be debugOn.
func (g *Game) SetDebugOn(debugOn bool) {
	g.debug = debugOn
}

// Log takes a log string and additional parameters, which can be substituted
// into the string using standard fmt.Printf rules.
// The formatted string is added to Game g's logs. If debug mode is on, the log will
// be printed to the terminal when Termloop exits.
func (g *Game) Log(log string, items ...interface{}) {
	toLog := "[" + time.Now().Format(time.StampMilli) + "] " +
		fmt.Sprintf(log, items...)
	g.logs = append(g.logs, toLog)
}

func (g *Game) dumpLogs() {
	if g.debug {
		fmt.Println("=== Logs: ===")
		for _, l := range g.logs {
			fmt.Println(l)
		}
		fmt.Println("=============")
	}
}

// SetEndKey sets the Key used to end the game. Default is KeyCtrlC.
// If you don't want an end key, set it to KeyEsc, as this key
// isn't supported and will do nothing.
// (We recommend always having an end key for development/testing.)
func (g *Game) SetEndKey(key tb.Key) {
	g.input.endKey = tb.Key(key)
}

// Start starts a Game running. This should be the last thing called in your
// main function. By default, the escape key exits.
func (g *Game) Start() {
	// Init Termbox
	err := tb.Init()
	tb.SetOutputMode(tb.Output256)
	tb.SetInputMode(tb.InputAlt | tb.InputMouse)
	if err != nil {
		panic(err)
	}
	defer g.dumpLogs()
	defer tb.Close()
	g.screen.resize(tb.Size())

	// Init input
	g.input.start()
	defer g.input.stop()
	clock := time.Now()

mainloop:
	for {
		update := time.Now()
		g.screen.delta = update.Sub(clock).Seconds()
		clock = update

		select {
		case ev := <-g.input.eventQ:
			if ev.Key == g.input.endKey {
				break mainloop
			} else if tb.EventType(ev.Type) == tb.EventResize {
				g.screen.resize(ev.Width, ev.Height)
			} else if tb.EventType(ev.Type) == tb.EventError {
				g.Log(ev.Err.Error())
			}
			g.screen.Tick(ev)
		default:
			g.screen.Tick(tb.Event{Type: tb.EventNone})
		}

		g.screen.Draw()
		// If g.screen.fps is zero (the default), then 1000.0/g.screen.fps -> +Inf -> time.Duration(+Inf), which
		// is a negative number, and so time.Sleep returns immediately.
		time.Sleep(time.Duration((update.Sub(time.Now()).Seconds()*1000.0)+1000.0/g.screen.fps) * time.Millisecond)
	}
}
