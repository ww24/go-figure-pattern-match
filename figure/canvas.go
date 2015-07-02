/**
 * canvas structure
 *
 */

package figure

import "math"

// Canvas structure
type Canvas struct {
	Original *Figure
	Figures  []*Figure
}

// NewCanvas is constructor
func NewCanvas(filename string) (canvas *Canvas) {
	canvas = &Canvas{
		Figures:  make([]*Figure, 0, 10),
		Original: NewFigure(filename)}

	figureCopy := canvas.Original.NewFigure(canvas.Original.Size, &Rect{
		MaxX: canvas.Original.Width - 1,
		MaxY: canvas.Original.Height - 1})

	canvas.Original.search(&canvas.Figures, figureCopy)

	return
}

// Search method
func (canvas *Canvas) Search(_figure *Figure) (x int, y int) {
	for _, figure := range canvas.Figures {
		if figure.Compare(_figure) {
			x = figure.X - _figure.OffsetLeft
			y = figure.Y - _figure.OffsetTop
			return
		}
	}
	return
}

// GetMaxSize will be return figure
func (canvas *Canvas) GetMaxSize() (_figure *Figure) {
	if len(canvas.Figures) == 0 {
		return
	}

	_figure = canvas.Figures[0]

	for _, figure := range canvas.Figures {
		if figure.Size > _figure.Size {
			_figure = figure
		}
	}
	return
}

func (rect *Rect) update(x int, y int) {
	rect.MinX = int(math.Min(float64(x), float64(rect.MinX)))
	rect.MinY = int(math.Min(float64(y), float64(rect.MinY)))
	rect.MaxX = int(math.Max(float64(x), float64(rect.MaxX)))
	rect.MaxY = int(math.Max(float64(y), float64(rect.MaxY)))
}

// Search figure method
func (figure *Figure) search(figures *[]*Figure, figureCopy *Figure) {
	x, y := func() (x int, y int) {
		for y, line := range figureCopy.Pixels {
			for x, col := range line {
				if col {
					return x, y
				}
			}
		}
		return
	}()

	if x == 0 && y == 0 {
		return
	}

	rect := figureCopy.NewRect()
	size := figureCopy.trace(x, y, rect)

	*figures = append(*figures, figure.NewFigure(size, rect))

	figure.search(figures, figureCopy)
}

func (figure *Figure) trace(x int, y int, rect *Rect) (count int) {
	if figure.Pixels[y][x] == false {
		return
	}

	rect.update(x, y)

	figure.Pixels[y][x] = false
	count = 1

	// right
	if figure.Pixels[y][x+1] {
		count += figure.trace(x+1, y, rect)
	}

	// left
	if figure.Pixels[y][x-1] {
		count += figure.trace(x-1, y, rect)
	}

	// bottom
	if figure.Pixels[y+1][x] {
		count += figure.trace(x, y+1, rect)
	}

	// top
	if figure.Pixels[y-1][x] {
		count += figure.trace(x, y-1, rect)
	}

	return
}
