package main

import (
	"testing"

	"ww24.jp/pattern/figure"
)

const patternFile = "Pattern.txt"

func BenchmarkExec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canvas := figure.NewCanvas(canvasFile)
		pattern := figure.NewFigure(patternFile)
		canvas.Search(pattern)
		canvas.GetMaxSize()
	}
}
