package controller

import (
	"cli-chess/model"
)

type View interface {
	Render()
	SetBoard(fen model.FenState)
	ShowEvaluation(evaluation string)
	HighlightSquare(x int, y int, color string)
	FontColorOnSquare(x int, y int, color string)
}

type Engine interface {
	Move(fen model.FenState, move string) (string, string, error)
	Evaluate() (string, error)
	NewGame()
	SuggestNextMove(fen model.FenState) string
}

type Controller interface {
	Start()
	HandleInput()
}
