package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"log"
	"os"
)

func main() {
	bounds := image.Rect(0, 0, 680, 420)
	var g gif.GIF
	for i := range 420 {
		log.Printf("%d/420", i+1)
		g.Image = append(g.Image, frame(i, bounds))
		g.Delay = append(g.Delay, 5)
	}
	f, _ := os.Create("gradients.gif")
	gif.EncodeAll(f, &g)
}

func frame(i int, bounds image.Rectangle) *image.Paletted {
	img := image.NewPaletted(bounds, palette.Plan9)
	for x := range bounds.Max.X {
		for y := range bounds.Max.Y {
			b := uint8((float64(x+i) + 1.62) / float64(bounds.Max.X) * 255)
			g := uint8(float64(y+i) / float64(bounds.Max.Y) * 255)
			img.Set(x, y, color.RGBA{B: b, G: g, A: 255})
		}
	}
	return img
}
