package wbncurses

import (
	"fmt"
)

type WannaBeNcurses struct {
	displays []*Display
}

func NewWannaBeNcurses() *WannaBeNcurses {
	return &WannaBeNcurses{}
}

func (wbnc *WannaBeNcurses) Render() {
	wbnc.clearTerminal()
	for _, display := range wbnc.displays {
		display.render()
	}
}

func (wbnc *WannaBeNcurses) CreateDisplay(width int, height int) *Display {
	display := newDisplay(width, height)
	wbnc.displays = append(wbnc.displays, display)
	return display
}

func (wbnc *WannaBeNcurses) clearTerminal() {
	fmt.Print("\033[H\033[2J")
}
