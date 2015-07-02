/**
 * figure structure
 *
 */

package figure

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Figure structure
type Figure struct {
	Pixels       [][]bool
	OffsetTop    int
	OffsetLeft   int
	OffsetBottom int
	OffsetRight  int
	Width        int
	Height       int
	Size         int
	X, Y         int
}

// Rect structure
type Rect struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
}

// NewFigure is constructor
func NewFigure(filename string) (figure *Figure) {
	figure = new(Figure)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = fmt.Fscanln(file, &figure.Height, &figure.Width)
	if err != nil {
		panic(err)
	}

	minX := figure.Width - 1
	maxX := 0

	figure.Pixels = make([][]bool, figure.Height)

	for y := 0; y < figure.Height; y++ {
		// この行に1が存在すれば true になる
		oneFlag := false

		figure.Pixels[y] = make([]bool, figure.Width)

		var buff string
		_, err = fmt.Fscanln(file, &buff)
		if err != nil {
			panic(err)
		}

		strs := strings.SplitN(buff, "", figure.Width)
		for x := 0; x < figure.Width; x++ {
			num, err := strconv.Atoi(strs[x])
			if err != nil {
				panic(err)
			}

			figure.Pixels[y][x] = num == 1

			if figure.Pixels[y][x] {
				figure.Size++
				oneFlag = true

				minX = int(math.Min(float64(x), float64(minX)))
				maxX = int(math.Max(float64(x), float64(maxX)))
			}
		}

		if oneFlag {
			if figure.OffsetTop == 0 {
				figure.OffsetTop = y
			}
		} else {
			if figure.OffsetTop > 0 && figure.OffsetBottom == 0 {
				figure.OffsetBottom = figure.Height - y
			}
		}
	}

	figure.OffsetLeft = minX
	figure.OffsetRight = figure.Width - maxX - 1

	return
}

// NewFigure is constructor
func (figure *Figure) NewFigure(size int, rect *Rect) (fig *Figure) {
	fig = &Figure{
		Width:  rect.MaxX - rect.MinX + 1,
		Height: rect.MaxY - rect.MinY + 1,
		Size:   size,
		X:      rect.MinX,
		Y:      rect.MinY}

	fig.Pixels = make([][]bool, fig.Height)

	for y := rect.MinY; y <= rect.MaxY; y++ {
		fig.Pixels[y-rect.MinY] = make([]bool, fig.Width)

		for x := rect.MinX; x <= rect.MaxX; x++ {
			fig.Pixels[y-rect.MinY][x-rect.MinX] = figure.Pixels[y][x]
		}
	}

	return
}

// NewRect is constructor
func (figure *Figure) NewRect() (rect *Rect) {
	rect = &Rect{
		MinX: figure.Width - 1,
		MinY: figure.Height - 1}
	return
}

func (figure *Figure) getInnerWidth() (width int) {
	width = figure.Width - figure.OffsetLeft - figure.OffsetRight
	return
}

func (figure *Figure) getInnerHeight() (height int) {
	height = figure.Height - figure.OffsetTop - figure.OffsetTop
	return
}

// Compare method
func (figure *Figure) Compare(_figure *Figure) bool {
	// if figure.getInnerWidth() != _figure.getInnerWidth() ||
	// 	figure.getInnerHeight() != _figure.getInnerHeight() ||
	// 	figure.Size != figure.Size {
	// 	return false
	// }

	for y, line := range figure.Pixels {
		for x, col := range line {
			if col != _figure.Pixels[y+_figure.OffsetTop][x+_figure.OffsetLeft] {
				return false
			}
		}
	}

	return true
}

// Print method for debug
func (figure *Figure) Print() {
	for _, line := range figure.Pixels {
		for _, col := range line {
			if col {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println("")
	}
}
