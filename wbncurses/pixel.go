package wbncurses

import "fmt"

const Reset = "\033[0m"

const Black = "\033[30m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const White = "\033[37m"
const Default = "\033[39m"

const BgBlack = "\033[40m"
const BgRed = "\033[41m"
const BgGreen = "\033[42m"
const BgYellow = "\033[43m"
const BgBlue = "\033[44m"
const BgMagenta = "\033[45m"
const BgCyan = "\033[46m"
const BgWhite = "\033[47m"
const BgDefault = "\033[49m"

type Pixel struct {
	X       int
	Y       int
	Char    string
	Color   string
	BgColor string
}

func (px *Pixel) render() {
	fmt.Printf("%s%s%s%s", px.Color, px.BgColor, px.Char, Reset)
}
