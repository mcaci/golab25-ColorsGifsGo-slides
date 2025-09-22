package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
)

func main() {
	frmBounds := image.Rect(0, 0, 1024, 768)
	frm1 := image.NewPaletted(frmBounds, palette.Plan9)
	draw.Draw(frm1, frmBounds, image.NewUniform(color.RGBA{G: 150, A: 255}), image.Pt(0, 0), draw.Over)
	frm2 := image.NewPaletted(frmBounds, palette.Plan9)
	draw.Draw(frm2, frmBounds, image.NewUniform(color.RGBA{G: 150, R: 150, A: 255}), image.Pt(0, 0), draw.Over)
	f, _ := os.Create("myFirst.gif")
	g := &gif.GIF{
		Image: []*image.Paletted{frm1, frm2},
		Delay: []int{100, 100},
	}
	gif.EncodeAll(f, g)
}
