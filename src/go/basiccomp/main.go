package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func MakeLayer(r image.Rectangle, c color.Color) *image.RGBA {
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			img.Set(x, y, c)
		}
	}
	return img
}

func main() {
	dst := MakeLayer(image.Rect(0, 0, 1024, 768), color.RGBA{G: 150, A: 255})
	src := MakeLayer(image.Rect(0, 0, 800, 600), color.White)
	draw.Draw(dst, image.Rect(224, 168, dst.Rect.Max.X, dst.Rect.Max.Y), src, image.Point{224, 168}, draw.Over)
	f, _ := os.Create("comp-white-on-green.png")
	png.Encode(f, src)
	dstBounds()
	srcPoint()
	srcClippedArea()
}

func dstBounds() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	const xMinBound, yMinBound = 224, 168
	for x := range r.Max.X {
		for y := range r.Max.Y {
			if (x < xMinBound+20 && x >= xMinBound-10) &&
				(y < yMinBound+20 && y >= yMinBound-10) {
				img.Set(x, y, color.RGBA{R: 255, A: 255})
				continue
			}
			if (x < r.Max.X && x >= r.Max.X-30) &&
				(y < r.Max.Y && y >= r.Max.Y-30) {
				img.Set(x, y, color.RGBA{R: 255, A: 255})
				continue
			}
			if x < xMinBound+10 && x >= xMinBound && y >= yMinBound {
				img.Set(x, y, color.Black)
				continue
			}
			if y < yMinBound+10 && y >= yMinBound && x >= xMinBound {
				img.Set(x, y, color.Black)
				continue
			}
			if x >= xMinBound && (y > r.Max.Y-10) {
				img.Set(x, y, color.Black)
				continue
			}
			if y >= yMinBound && (x > r.Max.X-10) {
				img.Set(x, y, color.Black)
				continue
			}
			if x > r.Max.X-2 || y > r.Max.Y-2 || y < 2 || x < 2 {
				img.Set(x, y, color.Black)
				continue
			}
			img.Set(x, y, color.RGBA{G: 150, A: 255})
		}
	}
	f, _ := os.Create("comp-bounds.png")
	png.Encode(f, img)
}

func srcPoint() {
	r := image.Rect(0, 0, 800, 600)
	img := image.NewRGBA(r)
	const xPoint, yPoint = 224, 168
	for x := range r.Max.X {
		for y := range r.Max.Y {
			if (x < xPoint+20 && x >= xPoint-20) &&
				(y < yPoint+20 && y >= yPoint-20) {
				img.Set(x, y, color.RGBA{R: 255, A: 255})
				continue
			}
			if x < xPoint+10 && x >= xPoint-10 && y >= yPoint-10 {
				img.Set(x, y, color.Black)
				continue
			}
			if y < yPoint+10 && y >= yPoint-10 && x >= xPoint-10 {
				img.Set(x, y, color.Black)
				continue
			}
			if x > r.Max.X-2 || y > r.Max.Y-2 || y < 2 || x < 2 {
				img.Set(x, y, color.Black)
				continue
			}
			img.Set(x, y, color.White)
		}
	}
	f, _ := os.Create("comp-src-point.png")
	png.Encode(f, img)
}

func srcClippedArea() {
	const xPoint, yPoint = 224, 168
	r := image.Rect(0, 0, 800-xPoint, 600-yPoint)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			if x < 20 && y < 20 {
				img.Set(x, y, color.RGBA{R: 255, A: 255})
				continue
			}
			if x > r.Max.X-2 || y > r.Max.Y-2 || y < 2 || x < 2 {
				img.Set(x, y, color.Black)
				continue
			}
			img.Set(x, y, color.White)
		}
	}
	f, _ := os.Create("comp-src-clipped-area.png")
	png.Encode(f, img)
}
