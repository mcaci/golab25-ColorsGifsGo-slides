package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"time"
)

func main() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			img.Set(x, y, cfuncT(time.Now()))
		}
	}
	f, _ := os.Create("pixset2.png")
	png.Encode(f, img)
}

func cfunc(x, y int) color.Color {
	r := math.Abs(math.Cos(float64(x))) * 50
	g := math.Abs(math.Sin(float64(y))) * 200
	b := math.Abs(math.Cos(float64(x*10-y*10))) * 50
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

func cfuncT(t time.Time) color.Color {
	b := t.Nanosecond() % 200
	g := t.Nanosecond() % 256
	return color.RGBA{B: uint8(b), G: uint8(g), A: 255}
}
