package controller

import (
	"cli-chess/model"
	"cli-chess/wbncurses"
	"fmt"

	tb "github.com/nsf/termbox-go"
)

type TerminalGuiController struct {
	IsRunning        bool
	engine           Engine
	view             View
	cursor           *model.Cursor
	moveCursor       *model.Cursor
	currentFen       *model.FenState
	recentFens       []*model.FenState
	historyIndex     int
	lastOpponentMove string
}

func NewTerminalGuiController(view View, engine Engine) *TerminalGuiController {
	startingFen := model.NewGame()

	return &TerminalGuiController{
		IsRunning:        false,
		view:             view,
		engine:           engine,
		cursor:           model.NewCursor(),
		moveCursor:       nil,
		currentFen:       startingFen,
		recentFens:       []*model.FenState{startingFen},
		historyIndex:     0,
		lastOpponentMove: "",
	}
}

func (tgc *TerminalGuiController) Start() {
	err := tb.Init()
	if err != nil {
		panic(err)
	}
	tgc.IsRunning = true
	tgc.view.SetBoard(*tgc.currentFen)
	tgc.view.HighlightSquare(tgc.cursor.X, tgc.cursor.Y, wbncurses.BgGreen)
	tgc.view.Render()
}

func (tgc *TerminalGuiController) UpdateView() {
	tgc.view.SetBoard(*tgc.currentFen)
	tgc.view.HighlightSquare(tgc.cursor.X, tgc.cursor.Y, wbncurses.BgGreen)
	if tgc.moveCursor != nil {
		tgc.view.HighlightSquare(tgc.moveCursor.X, tgc.moveCursor.Y, wbncurses.BgYellow)
	}

	if tgc.lastOpponentMove != "" {
		x, y := model.ChessNotationToPosition(tgc.lastOpponentMove[2:4])
		tgc.view.FontColorOnSquare(x, y, wbncurses.Red)
	}
	tgc.view.Render()
}

func (tgc *TerminalGuiController) Escape() {
	if tgc.cursor.State == "move" {
		tgc.cursor.State = "normal"
		tgc.moveCursor = nil
		tgc.UpdateView()
	} else {
		tgc.IsRunning = false
	}
}

func (tgc *TerminalGuiController) NewGame() {
	tgc.currentFen = model.NewGame()
	tgc.recentFens = []*model.FenState{tgc.currentFen}
	tgc.historyIndex = 0
	tgc.cursor = model.NewCursor()
	tgc.moveCursor = nil
	tgc.lastOpponentMove = ""
	tgc.UpdateView()
}

func (tgc *TerminalGuiController) CursorUp() {
	if tgc.cursor.State == "normal" {
		tgc.cursor.Up()
	} else {
		tgc.moveCursor.Up()
	}
	tgc.UpdateView()
}

func (tgc *TerminalGuiController) CursorDown() {
	if tgc.cursor.State == "normal" {
		tgc.cursor.Down()
	} else {
		tgc.moveCursor.Down()
	}
	tgc.UpdateView()
}

func (tgc *TerminalGuiController) CursorLeft() {
	if tgc.cursor.State == "normal" {
		tgc.cursor.Left()
	} else {
		tgc.moveCursor.Left()
	}
	tgc.UpdateView()
}

func (tgc *TerminalGuiController) CursorRight() {
	if tgc.cursor.State == "normal" {
		tgc.cursor.Right()
	} else {
		tgc.moveCursor.Right()
	}
	tgc.UpdateView()
}

func (tgc *TerminalGuiController) Select() {
	if tgc.cursor.State == "normal" {
		tgc.moveCursor = model.CopyCursor(tgc.cursor)
		tgc.cursor.State = "move"
		tgc.UpdateView()
	} else {
		// move the piece
		move := tgc.cursor.PositionToChessNotation() + tgc.moveCursor.PositionToChessNotation()
		newFen, engineMove, err := tgc.engine.Move(*tgc.currentFen, move)
		tgc.lastOpponentMove = engineMove

		if err != nil {
			fmt.Println(err)
		} else {
			if tgc.historyIndex < len(tgc.recentFens)-1 {
				tgc.recentFens = tgc.recentFens[:tgc.historyIndex+1]
			}
			tgc.currentFen = model.NewFenState(newFen)
			tgc.recentFens = append(tgc.recentFens, tgc.currentFen)
			tgc.historyIndex = len(tgc.recentFens) - 1
		}

		// update cursor
		tgc.cursor.State = "normal"
		tgc.cursor.X = tgc.moveCursor.X
		tgc.cursor.Y = tgc.moveCursor.Y
		tgc.moveCursor = nil
		tgc.UpdateView()
	}
}

func (tgc *TerminalGuiController) Hint() {
	suggestion := tgc.engine.SuggestNextMove(*tgc.currentFen)
	xFrom, yFrom := model.ChessNotationToPosition(suggestion[:2])
	xTo, yTo := model.ChessNotationToPosition(suggestion[2:4])

	tgc.view.HighlightSquare(xFrom, yFrom, wbncurses.BgYellow)
	tgc.view.HighlightSquare(xTo, yTo, wbncurses.BgYellow)
	tgc.view.Render()
}

func (tgc *TerminalGuiController) BackMove() {
	if tgc.historyIndex == 0 {
		return
	}

	tgc.historyIndex--
	tgc.currentFen = tgc.recentFens[tgc.historyIndex]
	tgc.lastOpponentMove = ""
	tgc.UpdateView()
}

func (tgc *TerminalGuiController) NextMove() {
	if tgc.historyIndex == len(tgc.recentFens)-1 {
		return
	}

	tgc.historyIndex++
	tgc.currentFen = tgc.recentFens[tgc.historyIndex]
	tgc.UpdateView()
}
