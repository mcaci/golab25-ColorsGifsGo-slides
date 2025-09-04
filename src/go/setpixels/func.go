package setpixels

import (
	"image/color"
	"math"
)

func Color(x, y int) color.Color {
	r := math.Abs(math.Sin(float64(x))) * 255
	g := math.Abs(math.Cos(float64(y))) * 255
	return color.RGBA{R: uint8(r), G: uint8(g), A: 255}
}
