package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			if x < 50 && y < 50 {
				img.Set(x, y, color.RGBA{R: 255, A: 255})
				continue
			}
			if x > r.Max.X-50 && y > r.Max.Y-50 {
				img.Set(x, y, color.RGBA{R: 255, A: 255})
				continue
			}
			if x < 10 || x > r.Max.X-10 {
				img.Set(x, y, color.Black)
				continue
			}
			if y < 10 || y > r.Max.Y-10 {
				img.Set(x, y, color.Black)
				continue
			}
			img.Set(x, y, color.White)
		}
	}
	f, _ := os.Create("bounds.png")
	png.Encode(f, img)
}
