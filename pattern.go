/**
 * figure pattern match
 *
 */

package main

import (
	"flag"
	"fmt"
	"os"

	"ww24.jp/pattern/figure"
)

const canvasFile = "Canvas.txt"

func main() {
	flag.Parse()

	filepath := flag.Arg(0)

	if filepath == "" {
		fmt.Fprintln(os.Stderr, "usage: pattern pattern_file")
		return
	}

	canvas := figure.NewCanvas(canvasFile)
	pattern := figure.NewFigure(filepath)
	resX, resY := canvas.Search(pattern)
	fmt.Printf("(%d, %d), ", resX, resY)
	fmt.Printf("(%d, %d), ", resX+pattern.Width, resY)
	fmt.Printf("(%d, %d), ", resX, resY+pattern.Height)
	fmt.Printf("(%d, %d)\n", resX+pattern.Width, resY+pattern.Height)

	figure := canvas.GetMaxSize()
	fmt.Printf("%dpx ", figure.Size)
	fmt.Printf("(%d, %d)\n", figure.X, figure.Y)
	// canvas.Figures[7].Print()
}
