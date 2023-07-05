package ssvg

import (
	"testing"
)

func TestWriteSvg(t *testing.T) {
	svg := new(Svg)

	// Front face of cube
	svg.Add(&Line{X1: 50, Y1: 50, X2: 150, Y2: 50})   // top edge
	svg.Add(&Line{X1: 50, Y1: 150, X2: 150, Y2: 150}) // bottom edge
	svg.Add(&Line{X1: 50, Y1: 50, X2: 50, Y2: 150})   // left edge
	svg.Add(&Line{X1: 150, Y1: 50, X2: 150, Y2: 150}) // right edge

	// Back face of cube
	svg.Add(&Line{X1: 0, Y1: 0, X2: 100, Y2: 0}) // top edge, offset -50, -50 and shrunken
	svg.Add(&Line{X1: 0, Y1: 0, X2: 0, Y2: 100}) // left edge, offset -50, -50 and shrunken

	// Connect front face with back face
	svg.Add(&Line{X1: 0, Y1: 0, X2: 50, Y2: 50})    // from back top left to front top left
	svg.Add(&Line{X1: 100, Y1: 0, X2: 150, Y2: 50}) // from back top right to front top right
	svg.Add(&Line{X1: 0, Y1: 100, X2: 50, Y2: 150}) // from back bottom left to front bottom left

	svg.Add(&Rect{X: 50, Y: 50, W: 100, H: 100})
	svg.Add(&Rect{X: 45, Y: 45, W: 100, H: 100})
	svg.Add(&Rect{X: 40, Y: 40, W: 100, H: 100})
	svg.Add(&Rect{X: 35, Y: 35, W: 100, H: 100})
	svg.Add(&Rect{X: 30, Y: 30, W: 100, H: 100})
	svg.Add(&Rect{X: 25, Y: 25, W: 100, H: 100})
	svg.Add(&Rect{X: 20, Y: 20, W: 100, H: 100})
	svg.Add(&Rect{X: 15, Y: 15, W: 100, H: 100})
	svg.Add(&Rect{X: 10, Y: 10, W: 100, H: 100})
	svg.Add(&Rect{X: 5, Y: 5, W: 100, H: 100})
	// Text
	svg.Add(&Text{X: 0, Y: -50, Text: "http://firsh.me"})
	svg.WriteFile("test.svg", 220)
}
