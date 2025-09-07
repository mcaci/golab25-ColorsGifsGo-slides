package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

func main() {
	r := image.Rect(0, 0, 600, 800)
	base := image.NewRGBA(r)
	bg := image.NewUniform(color.Gray{Y: 200})

	draw.Draw(base, base.Bounds(), bg, bg.Bounds().Min, draw.Over)
	draw.Draw(base, image.Rect(200, 250, 400, 550), body(), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(250, 150, 350, 250), eyes(), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(275, 175, 300, 200), eyes(), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(150, 200, 250, 300), ears(), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(350, 200, 450, 300), ears(), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(275, 275, 325, 325), nose(), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(250, 325, 350, 375), mouth(), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(175, 500, 275, 600), paws(), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(325, 500, 425, 600), paws(), image.Point{0, 0}, draw.Over)

	out, _ := os.Create("picasso-gopher.png")
	png.Encode(out, base)
}

func body() image.Image {
	const (
		rx, ry = 200, 300
		dx, dy = rx - 100, ry - 150
		a, b   = rx / 2, ry / 2
		a2, b2 = a * a, b * b
	)
	r := image.Rect(0, 0, rx, ry)
	base := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			var x2 = float64((x - dx) * (x - dx))
			var y2 = float64((y - dy) * (y - dy))

			switch {
			case math.Abs(x2/a2+y2/b2) < 1:
				base.Set(x, y, color.RGBA{R: 0, G: 0, B: 255, A: 255})
			case math.Abs(x2/a2+y2/b2) > 1:
				base.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 0})
			default:
				base.Set(x, y, color.Black)
			}
		}
	}
	return base
}

func eyes() image.Image {
	const (
		rx, ry = 200, 100
		dx, dy = rx - 100, ry - 50
		a, b   = rx / 2, ry / 2
		a2, b2 = a * a, b * b
	)
	r := image.Rect(0, 0, rx, ry)
	base := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			var x2 = float64((x - dx) * (x - dx))
			var y2 = float64((y - dy) * (y - dy))

			switch {
			case math.Abs(x2/a2+y2/b2) < 1:
				base.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			case math.Abs(x2/a2+y2/b2) > 1:
				base.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 0})
			default:
				base.Set(x, y, color.Black)
			}
		}
	}
	return base
}

func ears() image.Image {
	const (
		rx, ry = 100, 100
		dx, dy = rx - 50, ry - 50
		a, b   = rx / 2, ry / 2
		a2, b2 = a * a, b * b
	)
	r := image.Rect(0, 0, rx, ry)
	base := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			var x2 = float64((x - dx) * (x - dx))
			var y2 = float64((y - dy) * (y - dy))

			switch {
			case math.Abs(x2/a2+y2/b2) < 1:
				base.Set(x, y, color.RGBA{R: 255, G: 200, B: 0, A: 255})
			case math.Abs(x2/a2+y2/b2) > 1:
				base.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 0})
			default:
				base.Set(x, y, color.Black)
			}
		}
	}
	return base
}

func mouth() image.Image {
	const (
		rx, ry = 100, 50
		dx, dy = rx - 50, ry - 25
		a, b   = rx / 2, ry / 2
		a2, b2 = a * a, b * b
	)
	r := image.Rect(0, 0, rx, ry)
	base := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			var x2 = float64((x - dx) * (x - dx))
			var y2 = float64((y - dy) * (y - dy))

			switch {
			case math.Abs(x2/a2+y2/b2) < 1:
				base.Set(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})
			case math.Abs(x2/a2+y2/b2) > 1:
				base.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 0})
			default:
				base.Set(x, y, color.Black)
			}
		}
	}
	return base
}

func nose() image.Image {
	const (
		rx, ry = 50, 50
		dx, dy = rx - 25, ry - 25
		a, b   = rx / 2, ry / 2
		a2, b2 = a * a, b * b
	)
	r := image.Rect(0, 0, rx, ry)
	base := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			var x2 = float64((x - dx) * (x - dx))
			var y2 = float64((y - dy) * (y - dy))

			switch {
			case math.Abs(x2/a2+y2/b2) < 1:
				base.Set(x, y, color.RGBA{R: 0, G: 0, B: 0, A: 255})
			case math.Abs(x2/a2+y2/b2) > 1:
				base.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 0})
			default:
				base.Set(x, y, color.Black)
			}
		}
	}
	return base
}

func paws() image.Image {
	const (
		rx, ry = 100, 100
		dx, dy = rx - 50, ry - 50
		a, b   = rx / 2, ry / 2
		a2, b2 = a * a, b * b
	)
	r := image.Rect(0, 0, rx, ry)
	base := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			var x2 = float64((x - dx) * (x - dx))
			var y2 = float64((y - dy) * (y - dy))

			switch {
			case math.Abs(x2/a2+y2/b2) < 1:
				base.Set(x, y, color.RGBA{R: 255, G: 200, B: 0, A: 255})
			case math.Abs(x2/a2+y2/b2) > 1:
				base.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 0})
			default:
				base.Set(x, y, color.Black)
			}
		}
	}
	return base
}
