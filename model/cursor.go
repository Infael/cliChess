package model

import (
	"fmt"
	"strconv"
)

type Cursor struct {
	X     int
	Y     int
	State string
}

func NewCursor() *Cursor {
	return &Cursor{X: 0, Y: 0, State: "normal"}
}

func CopyCursor(c *Cursor) *Cursor {
	return &Cursor{X: c.X, Y: c.Y, State: c.State}
}

func (c *Cursor) Up() {
	if c.Y > 0 {
		c.Y--
	}
}

func (c *Cursor) Down() {
	if c.Y < 7 {
		c.Y++
	}
}

func (c *Cursor) Left() {
	if c.X > 0 {
		c.X--
	}
}

func (c *Cursor) Right() {
	if c.X < 7 {
		c.X++
	}
}

func (c *Cursor) PositionToChessNotation() string {

	letter := string(rune(c.X + 97))
	number := strconv.Itoa(8 - c.Y)

	fmt.Println("letter: ", letter)
	fmt.Println("number: ", number)
	return letter + number
}

func ChessNotationToPosition(chessNotation string) (int, int) {
	x := int(chessNotation[0]) - 97
	y := 8 - int(chessNotation[1]-48)
	return x, y
}
