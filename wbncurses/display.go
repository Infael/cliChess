package wbncurses

import "fmt"

type Display struct {
	width  int
	height int
	pixels [][]Pixel
}

func newDisplay(width int, height int) *Display {
	display := Display{
		width:  width,
		height: height,
		pixels: make([][]Pixel, height),
	}
	for i := range display.pixels {
		display.pixels[i] = make([]Pixel, width)
	}
	return &display
}

func (d *Display) SetPixel(x int, y int, char string, color string, bgColor string) {
	d.pixels[y][x] = Pixel{
		X:       x,
		Y:       y,
		Char:    char,
		Color:   color,
		BgColor: bgColor,
	}
}

func (d *Display) GetPixel(x int, y int) Pixel {
	return d.pixels[y][x]
}

func (d *Display) Clear() {
	for i := range d.pixels {
		for j := range d.pixels[i] {
			d.pixels[i][j] = Pixel{}
		}
	}
}

func (d *Display) SetAll(char string, color string, bgColor string) {
	for i := range d.pixels {
		for j := range d.pixels[i] {
			d.SetPixel(j, i, char, color, bgColor)
		}
	}
}

func (d *Display) render() {
	for i := range d.pixels {
		for j := range d.pixels[i] {
			d.pixels[i][j].render()
		}
		fmt.Println()
	}
}
