package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"math"
	"os"
)

func main() {
	r := image.Rect(0, 0, 300, 400)
	// Image
	base := image.NewRGBA(r)
	bg := image.NewUniform(HexToRGBA("#adffffff"))
	draw.Draw(base, base.Bounds(), bg, bg.Bounds().Min, draw.Over)
	draw.DrawMask(base, image.Rect(37, 40, 100, 150), ears(HexToRGBA("#0e7d99ff")), image.Point{0, 0}, body(HexToRGBA("#00add8ff")), image.Point{0, 0}, draw.Over)
	draw.DrawMask(base, image.Rect(170, 50, 300, 150), ears(HexToRGBA("#0e7d99ff")), image.Point{0, 0}, body(HexToRGBA("#00add8ff")), image.Point{-20, +120}, draw.Over)
	draw.Draw(base, image.Rect(50, 50, 250, 350), body(HexToRGBA("#00add8ff")), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(75, 130, 200, 80), eyes(color.White), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(130, 130, 183, 50), eyes(color.White), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(100, 125, 200, 175), noseNmouth(HexToRGBA("#eeefa0ff")), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(150, 180, 240, 280), paws(HexToRGBA("#eeefd0ff")), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(50, 180, 140, 280), paws(HexToRGBA("#eeefd0ff")), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(50, 280, 140, 380), paws(HexToRGBA("#eeefd0ff")), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(150, 280, 240, 380), paws(HexToRGBA("#eeefd0ff")), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(143, 168, 240, 280), tooth(color.White), image.Point{0, 0}, draw.Over)
	draw.Draw(base, image.Rect(138, 168, 240, 280), tooth(color.White), image.Point{0, 0}, draw.Over)
	out, _ := os.Create("picasso-gopher.png")
	defer out.Close()
	png.Encode(out, base)
	// Party Gopher GIF
	// Frame 1
	frame1 := image.NewPaletted(r, palette.Plan9)
	draw.Draw(frame1, frame1.Bounds(), bg, bg.Bounds().Min, draw.Over)
	draw.DrawMask(frame1, image.Rect(37, 40, 100, 150), ears(HexToRGBA("#0e7d99ff")), image.Point{0, 0}, body(HexToRGBA("#00add8ff")), image.Point{0, 0}, draw.Over)
	draw.DrawMask(frame1, image.Rect(170, 50, 300, 150), ears(HexToRGBA("#0e7d99ff")), image.Point{0, 0}, body(HexToRGBA("#00add8ff")), image.Point{-20, +120}, draw.Over)
	draw.Draw(frame1, image.Rect(50, 50, 250, 350), body(HexToRGBA("#00add8ff")), image.Point{0, 0}, draw.Over)
	draw.Draw(frame1, image.Rect(75, 130, 200, 80), eyes(color.White), image.Point{0, 0}, draw.Over)
	draw.Draw(frame1, image.Rect(130, 130, 183, 50), eyes(color.White), image.Point{0, 0}, draw.Over)
	draw.Draw(frame1, image.Rect(100, 125, 200, 175), noseNmouth(HexToRGBA("#eeefa0ff")), image.Point{0, 0}, draw.Over)
	draw.Draw(frame1, image.Rect(150, 180, 240, 280), paws(HexToRGBA("#eaecaeff")), image.Point{0, 0}, draw.Over)
	draw.Draw(frame1, image.Rect(50, 180, 140, 280), paws(HexToRGBA("#eaecaeff")), image.Point{0, 0}, draw.Over)
	draw.Draw(frame1, image.Rect(50, 280, 140, 380), paws(HexToRGBA("#eaecaeff")), image.Point{0, 0}, draw.Over)
	draw.Draw(frame1, image.Rect(150, 280, 240, 380), paws(HexToRGBA("#eaecaeff")), image.Point{0, 0}, draw.Over)
	draw.Draw(frame1, image.Rect(143, 168, 240, 280), tooth(color.White), image.Point{0, 0}, draw.Over)
	draw.Draw(frame1, image.Rect(138, 168, 240, 280), tooth(color.White), image.Point{0, 0}, draw.Over)
	// Frame 2
	frame2 := image.NewPaletted(r, palette.Plan9)
	draw.Draw(frame2, frame2.Bounds(), bg, bg.Bounds().Min, draw.Over)
	draw.DrawMask(frame2, image.Rect(37, 40, 100, 150), ears(HexToRGBA("#0e7d99ff")), image.Point{0, 0}, body(HexToRGBA("#00add8ff")), image.Point{0, 0}, draw.Over)
	draw.DrawMask(frame2, image.Rect(170, 50, 300, 150), ears(HexToRGBA("#0e7d99ff")), image.Point{0, 0}, body(HexToRGBA("#00add8ff")), image.Point{-20, +120}, draw.Over)
	draw.Draw(frame2, image.Rect(50, 50, 250, 350), body(HexToRGBA("#00add8ff")), image.Point{0, 0}, draw.Over)
	draw.Draw(frame2, image.Rect(75, 130, 200, 80), eyes(color.White), image.Point{0, 0}, draw.Over)
	draw.Draw(frame2, image.Rect(130, 130, 183, 50), eyes(color.White), image.Point{0, 0}, draw.Over)
	draw.Draw(frame2, image.Rect(100, 125, 200, 175), noseNmouth(HexToRGBA("#eeefa0ff")), image.Point{0, 0}, draw.Over)
	draw.Draw(frame2, image.Rect(200, 180, 290, 280), paws(HexToRGBA("#eaecaeff")), image.Point{0, 0}, draw.Over)
	draw.Draw(frame2, image.Rect(100, 180, 190, 280), paws(HexToRGBA("#eaecaeff")), image.Point{0, 0}, draw.Over)
	draw.Draw(frame2, image.Rect(50, 280, 140, 380), paws(HexToRGBA("#eaecaeff")), image.Point{0, 0}, draw.Over)
	draw.Draw(frame2, image.Rect(150, 280, 240, 380), paws(HexToRGBA("#eaecaeff")), image.Point{0, 0}, draw.Over)
	draw.Draw(frame2, image.Rect(143, 168, 240, 280), tooth(color.White), image.Point{0, 0}, draw.Over)
	draw.Draw(frame2, image.Rect(138, 168, 240, 280), tooth(color.White), image.Point{0, 0}, draw.Over)
	g := &gif.GIF{}
	g.Image = append(g.Image, frame1, frame2)
	g.Delay = append(g.Delay, 30, 30)
	outGif, _ := os.Create("picasso-gopher.gif")
	defer outGif.Close()
	gif.EncodeAll(outGif, g)
}

func HexToRGBA(hex string) color.Color {
	var r, g, b, a uint8
	if hex[0] == '#' {
		hex = hex[1:]
	}
	if len(hex) != 8 {
		return color.Opaque
	}
	_, err := fmt.Sscanf(hex, "%02x%02x%02x%02x", &r, &g, &b, &a)
	if err != nil {
		return color.Transparent
	}
	return color.RGBA{r, g, b, a}
}

func body(c color.Color) image.Image {
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
				base.Set(x, y, c)
			case math.Abs(x2/a2+y2/b2) >= 1:
				base.Set(x, y, color.Transparent)
			default:
				base.Set(x, y, color.Black)
			}
		}
	}
	return base
}

var (
	eyes       = roundWithInnerBlackSquare
	ears       = roundWithInnerBlackSquare
	noseNmouth = roundWithInnerBlackSquare
)

func roundWithInnerBlackSquare(c color.Color) image.Image {
	const (
		rx, ry, rds = 200, 200, 20
		dx, dy      = 50, 50
		r2          = rds * rds
	)
	r := image.Rect(0, 0, rx, ry)
	base := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			var x2 = float64((x - dx) * (x - dx))
			var y2 = float64((y - dy) * (y - dy))

			switch {
			case math.Abs(x2+y2) < r2:
				base.Set(x, y, c)
			case math.Abs(x2+y2) >= r2:
				base.Set(x, y, color.Transparent)
			default:
				base.Set(x, y, color.Black)
			}
		}
	}
	iris := image.NewUniform(color.Black)
	draw.Draw(base, image.Rect(40, 40, 60, 60), iris, iris.Bounds().Min, draw.Over)
	return base
}

func paws(c color.Color) image.Image {
	const (
		rx, ry = 100, 100
		dx, dy = rx / 2, ry / 2
		a, b   = 10, 20
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
				base.Set(x, y, c)
			case math.Abs(x2/a2+y2/b2) >= 1:
				base.Set(x, y, color.Transparent)
			default:
				base.Set(x, y, color.Black)
			}
		}
	}
	return base
}

func tooth(c color.Color) image.Image {
	const (
		rx, ry = 20, 20
		dx, dy = rx / 2, ry / 2
		a, b   = 2, 4
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
				base.Set(x, y, c)
			case math.Abs(x2/a2+y2/b2) >= 1:
				base.Set(x, y, color.Transparent)
			default:
				base.Set(x, y, color.Black)
			}
		}
	}
	return base
}
