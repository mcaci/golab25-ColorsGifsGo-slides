package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"image/png"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	green()
	blue()
	bgGradient()
	bwChessboard()
	rbgGradient()
	rgCosSinStripes()
	rgbNanoseconds()
	rgbNanosecondsGif()
}

func blue() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			img.Set(x, y, color.RGBA{B: 150, A: 255})
		}
	}
	f, _ := os.Create("blue.png")
	png.Encode(f, img)
}

func green() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			img.Set(x, y, color.RGBA{G: 150, A: 255})
		}
	}
	f, _ := os.Create("green.png")
	png.Encode(f, img)
}

func bgGradient() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			b := float64(x) / float64(r.Max.X) * 255
			g := float64(y) / float64(r.Max.Y) * 255
			img.Set(x, y, color.RGBA{B: uint8(b), G: uint8(g), A: 255})
		}
	}
	f, _ := os.Create("bgGradient.png")
	png.Encode(f, img)
}

func bwChessboard() {
	r := image.Rect(0, 0, 1024, 1024)
	const l = 1024 / 8
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			switch {
			case x < 10 || y < 10 || x > r.Max.X-10 || y > r.Max.Y-10:
				img.Set(x, y, color.Black)
			case (x/128+y/128)%2 == 0:
				img.Set(x, y, color.White)
			default:
				img.Set(x, y, color.Black)
			}
		}
	}
	f, _ := os.Create("bwChessboard.png")
	png.Encode(f, img)
}

func rbgGradient() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			img.Set(x, y, color.RGBA{R: uint8(x / 4), G: uint8(y / 3), B: 150, A: 255})
		}
	}
	f, _ := os.Create("rbgGradient.png")
	png.Encode(f, img)
}

func rgCosSinStripes() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			r := (1 + math.Cos(float64(x)/10)) * 255
			g := (1 + math.Sin(float64(y)/10)) * 255
			img.Set(x, y, color.RGBA{R: uint8(r), G: uint8(g), B: 0, A: 255})
		}
	}
	f, _ := os.Create("rgCosSinStripes.png")
	png.Encode(f, img)
}

func rgbNanoseconds() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			nano := uint32(time.Now().UnixNano())
			r := uint8((nano >> 16) & 0xFF)
			g := uint8((nano >> 8) & 0xFF)
			b := uint8(nano & 0xFF)
			img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}
	f, _ := os.Create("rgbNanoseconds.png")
	png.Encode(f, img)
}

func rgbNanosecondsGif() {
	var frames []*image.Paletted
	var delays []int
	for i := range 100 {
		log.Print(i)
		r := image.Rect(0, 0, 1024, 768)
		img := image.NewPaletted(r, palette.Plan9)
		for x := range r.Max.X {
			for y := range r.Max.Y {
				nano := uint32(time.Now().UnixNano())
				r := uint8(nano & 0xFF)
				g := uint8((nano >> 8) & 0xFF)
				b := uint8((nano >> 16) & 0xFF)
				img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
			}
		}
		frames = append(frames, img)
		delays = append(delays, 10)
	}
	f, _ := os.Create("nanosecRGB.gif")
	g := &gif.GIF{
		Image: frames,
		Delay: delays,
	}
	gif.EncodeAll(f, g)
}
