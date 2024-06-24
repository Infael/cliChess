package main

import (
	"cli-chess/controller"
	"cli-chess/engineconnector"
	"cli-chess/view"

	tb "github.com/nsf/termbox-go"
)

func main() {

	engine := engineconnector.NewEngineConnector("stockfish")
	engine.Start()
	view := view.NewWannaBeNCursesView()
	controller := controller.NewTerminalGuiController(view, engine)

	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()

	controller.Start()
	for controller.IsRunning {

		ev := tb.PollEvent()

		if ev.Ch == 'b' {
			controller.BackMove()
		} else if ev.Ch == 'n' {
			controller.NextMove()
		} else if ev.Ch == 'h' {
			controller.Hint()
		} else if ev.Ch == 'r' {
			controller.NewGame()
		} else {
			switch ev.Type {
			case tb.EventKey:
				switch ev.Key {
				case tb.KeyArrowUp:
					controller.CursorUp()
				case tb.KeyArrowDown:
					controller.CursorDown()
				case tb.KeyArrowLeft:
					controller.CursorLeft()
				case tb.KeyArrowRight:
					controller.CursorRight()
				case tb.KeySpace:
					controller.Select()
				case tb.KeyEsc:
					controller.Escape()
				}
			}

		}
	}
}
