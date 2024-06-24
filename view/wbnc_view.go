package view

import (
	"cli-chess/model"
	"cli-chess/wbncurses"
	"strings"
)

type WannaBeNCursesView struct {
	wannaBeNCurses *wbncurses.WannaBeNcurses
	chessBoard     *wbncurses.Display
}

func NewWannaBeNCursesView() *WannaBeNCursesView {
	nCurses := wbncurses.NewWannaBeNcurses()
	chessBoard := nCurses.CreateDisplay(9, 9)
	return &WannaBeNCursesView{wannaBeNCurses: nCurses, chessBoard: chessBoard}
}

func (vbncv *WannaBeNCursesView) Render() {
	vbncv.wannaBeNCurses.Render()
}

func (vbncv *WannaBeNCursesView) SetBoard(fen model.FenState) {
	vbncv.fenToBoard(fen.GetFen())
}

func (vbncv *WannaBeNCursesView) HighlightSquare(x int, y int, color string) {
	vbncv.chessBoard.SetPixel(x, y, vbncv.chessBoard.GetPixel(x, y).Char, wbncurses.White, color)
}

func (vbncv *WannaBeNCursesView) FontColorOnSquare(x int, y int, color string) {
	pixel := vbncv.chessBoard.GetPixel(x, y)
	vbncv.chessBoard.SetPixel(x, y, pixel.Char, color, pixel.BgColor)
}

func (vbncv *WannaBeNCursesView) ShowEvaluation(evaluation string) {
}

func (vbncv *WannaBeNCursesView) fenToBoard(fen string) {
	board := strings.Split(fen, " ")[0]
	rows := strings.Split(board, "/")

	vbncv.chessBoard.SetPixel(0, 0, " ", wbncurses.White, wbncurses.BgDefault)

	// uncoment to show letters on top
	// for i := 0; i < 8; i++ {
	// 	vbncv.chessBoard.SetPixel(i+1, 0, string(rune(i+97)), wbncurses.White, wbncurses.BgDefault)
	// }

	for i, row := range rows {
		// uncoment to show numbers on the left
		// vbncv.chessBoard.SetPixel(0, i+1, strconv.Itoa(8-i), wbncurses.White, wbncurses.BgDefault)
		x := 0
		for j := 0; j < len(row); j++ {
			x += vbncv.fenToPixel(string(row[j]), x, i)
		}
	}

}

func (vbncv *WannaBeNCursesView) fenToPixel(char string, x int, y int) int {
	switch char {
	case "P":
		vbncv.chessBoard.SetPixel(x, y, "♟", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "p":
		vbncv.chessBoard.SetPixel(x, y, "♙", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "r":
		vbncv.chessBoard.SetPixel(x, y, "♖", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "R":
		vbncv.chessBoard.SetPixel(x, y, "♜", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "n":
		vbncv.chessBoard.SetPixel(x, y, "♘", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "N":
		vbncv.chessBoard.SetPixel(x, y, "♞", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "b":
		vbncv.chessBoard.SetPixel(x, y, "♗", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "B":
		vbncv.chessBoard.SetPixel(x, y, "♝", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "q":
		vbncv.chessBoard.SetPixel(x, y, "♕", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "Q":
		vbncv.chessBoard.SetPixel(x, y, "♛", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "k":
		vbncv.chessBoard.SetPixel(x, y, "♔", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "K":
		vbncv.chessBoard.SetPixel(x, y, "♚", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "1":
		vbncv.chessBoard.SetPixel(x, y, " ", wbncurses.White, vbncv.getPixelBackground(x, y))
	case "2":
		for i := 0; i < 2; i++ {
			vbncv.chessBoard.SetPixel(x+i, y, " ", wbncurses.White, vbncv.getPixelBackground(x+i, y))
		}
		return 2
	case "3":
		for i := 0; i < 3; i++ {
			vbncv.chessBoard.SetPixel(x+i, y, " ", wbncurses.White, vbncv.getPixelBackground(x+i, y))
		}
		return 3
	case "4":
		for i := 0; i < 4; i++ {
			vbncv.chessBoard.SetPixel(x+i, y, " ", wbncurses.White, vbncv.getPixelBackground(x+i, y))
		}
		return 4
	case "5":
		for i := 0; i < 5; i++ {
			vbncv.chessBoard.SetPixel(x+i, y, " ", wbncurses.White, vbncv.getPixelBackground(x+i, y))
		}
		return 5
	case "6":
		for i := 0; i < 6; i++ {
			vbncv.chessBoard.SetPixel(x+i, y, " ", wbncurses.White, vbncv.getPixelBackground(x+i, y))
		}
		return 6
	case "7":
		for i := 0; i < 7; i++ {
			vbncv.chessBoard.SetPixel(x+i, y, " ", wbncurses.White, vbncv.getPixelBackground(x+i, y))
		}
		return 7
	case "8":
		for i := 0; i < 8; i++ {
			vbncv.chessBoard.SetPixel(x+i, y, " ", wbncurses.White, vbncv.getPixelBackground(x+i, y))
		}
		return 8
	}

	return 1
}

func (vbncv *WannaBeNCursesView) getPixelBackground(x int, y int) string {
	if (x%2 == 0 && y%2 == 0) || (x%2 == 1 && y%2 == 1) {
		return wbncurses.BgDefault
	}
	return wbncurses.BgBlack
}
