package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

// painting compositionInRedBlueAndYellow from Piet Mondrian
func main() {
	const l, h = 45, 45
	r := image.Rect(0, 0, l, h)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			vlStart1, vlEnd1 := float64(r.Max.X)*0.21, float64(r.Max.X)*0.24
			vlStart2, vlEnd2 := float64(r.Max.X)*0.87, float64(r.Max.X)*0.9
			hlStart1, hlEnd1 := float64(r.Max.Y)*0.28, float64(r.Max.Y)*0.33
			hlStart2, hlEnd2 := float64(r.Max.Y)*0.7, float64(r.Max.Y)*0.73
			hlStart3, hlEnd3 := float64(r.Max.Y)*0.84, float64(r.Max.Y)*0.88
			if x > int(vlEnd1) && y <= int(hlStart2) {
				img.Set(x, y, color.RGBA{R: 255, A: 255})
				continue
			}
			if x <= int(vlStart1) && y > int(hlEnd2) {
				img.Set(x, y, color.RGBA{B: 255, A: 255})
				continue
			}
			if x > int(vlEnd2) && y > int(hlEnd3) {
				img.Set(x, y, color.RGBA{R: 255, G: 255, A: 255})
				continue
			}
			if x > int(vlStart1) && x <= int(vlEnd1) {
				img.Set(x, y, color.Black)
				continue
			}
			if y > int(hlStart1) && y <= int(hlEnd1) && x <= int(vlStart1) {
				img.Set(x, y, color.Black)
				continue
			}
			if y > int(hlStart2) && y <= int(hlEnd2) {
				img.Set(x, y, color.Black)
				continue
			}
			if x > int(vlStart2) && x <= int(vlEnd2) && y > int(hlStart2) {
				img.Set(x, y, color.Black)
				continue
			}
			if y > int(hlStart3) && y <= int(hlEnd3) && x > int(vlStart2) {
				img.Set(x, y, color.Black)
				continue
			}
			img.Set(x, y, color.White)
		}
	}

	dstR := image.Rect(0, 0, l*15, h*15)
	dst := image.NewRGBA(dstR)
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)
	f, _ := os.Create("pietGondrian.png")
	png.Encode(f, dst)
}

func firsttry(img *image.RGBA) {
	r := img.Bounds()
	for x := range r.Max.X {
		for y := range r.Max.Y {
			if (x < 15 && y < 15) || (x >= 30 && y >= 30) {
				img.Set(x, y, image.Black)
			} else if x < 15 && y >= 15 && y < 30 {
				img.Set(x, y, image.White)
			} else if x >= 15 && x < 30 && y < 15 {
				img.Set(x, y, image.White)
			} else if x >= 15 && x < 30 && y >= 15 && y < 30 {
				img.Set(x, y, color.RGBA{R: 255, A: 255})
			} else if x >= 30 && y < 15 {
				img.Set(x, y, color.RGBA{B: 255, A: 255})
			} else if x >= 30 && y >= 15 && y < 30 {
				img.Set(x, y, image.White)
			} else if x < 15 && y >= 30 {
				img.Set(x, y, color.RGBA{R: 255, G: 255, A: 255})
			} else if x >= 15 && x < 30 && y >= 30 {
				img.Set(x, y, image.White)
			}
		}
	}
}
