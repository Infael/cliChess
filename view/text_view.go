package view

import (
	"cli-chess/model"
	"fmt"
	"strings"
)

type View struct {
}

func NewView() *View {
	return &View{}
}

// RenderBoard renders the board to the screen
func (v *View) RenderBoard(fen model.FenState) {
	v.fenToBoard(fen.GetFen())
}

func (v *View) ShowEvaluation(evaluation string) {
	fmt.Println("Evaluation: ", evaluation)
}

func (v *View) fenToBoard(fen string) {
	board := strings.Split(fen, " ")[0]
	rows := strings.Split(board, "/")
	fmt.Println()
	fmt.Println("  +-+-+-+-+-+-+-+-+")
	for i, row := range rows {
		fmt.Print(8-i, " ")
		for _, cell := range row {
			fmt.Print("|")
			v.cellToPiece(string(cell))
			fmt.Print("")
		}
		fmt.Print("|")
		fmt.Println()
		fmt.Println("  +-+-+-+-+-+-+-+-+")
	}
	fmt.Println("   a b c d e f g h")
	fmt.Println()
}

func (v *View) cellToPiece(cell string) {
	switch cell {
	case "p":
		fmt.Print("♟")
	case "P":
		fmt.Print("♙")
	case "r":
		fmt.Print("♜")
	case "R":
		fmt.Print("♖")
	case "n":
		fmt.Print("♞")
	case "N":
		fmt.Print("♘")
	case "b":
		fmt.Print("♝")
	case "B":
		fmt.Print("♗")
	case "q":
		fmt.Print("♛")
	case "Q":
		fmt.Print("♕")
	case "k":
		fmt.Print("♚")
	case "K":
		fmt.Print("♔")
	case "1":
		fmt.Print(" ")
	case "2":
		fmt.Print(" | ")
	case "3":
		fmt.Print(" | | ")
	case "4":
		fmt.Print(" | | | ")
	case "5":
		fmt.Print(" | | | | ")
	case "6":
		fmt.Print(" | | | | | ")
	case "7":
		fmt.Print(" | | | | | | ")
	case "8":
		fmt.Print(" | | | | | | | ")
	}
}
