package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"log"
	"os"
)

func main() {

	r := image.Rect(0, 0, 1000, 100)

	pbar := image.NewPaletted(r, palette.Plan9)
	draw.Draw(pbar, r, image.White, image.Point{}, draw.Src)
	drawBorder(pbar, r, color.Black)
	var frames []*image.Paletted
	var delay []int
	frames = append(frames, pbar)
	delay = append(delay, 100)
	for i := 5; i <= 1000; i += 5 {
		pbar := image.NewPaletted(r, palette.Plan9)
		draw.Draw(pbar, r, image.White, image.Point{}, draw.Src)
		draw.Draw(pbar, image.Rect(0, 0, i, 100), image.NewUniform(color.Gray{Y: 150}), image.Point{}, draw.Src)
		drawBorder(pbar, r, color.Black)
		frames = append(frames, pbar)
		if i == 1000 {
			delay = append(delay, 100)
			continue
		}
		delay = append(delay, 8)
	}
	g := &gif.GIF{
		Image:     frames,
		Delay:     delay,
		LoopCount: -1,
	}
	f, err := os.Create("progressbar.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = gif.EncodeAll(f, g)
	if err != nil {
		log.Fatal(err)
	}
}

func drawBorder(img draw.Image, r image.Rectangle, c color.Color) {
	for x := range r.Max.X {
		for y := range r.Max.Y {
			switch {
			case x < 10,
				x > r.Max.X-10,
				y < 10,
				y > r.Max.Y-10:
				img.Set(x, y, c)
			default:
				continue
			}
		}
	}
}

func pngProgressBar() {
	r := image.Rect(0, 0, 1000, 100)
	pbar := image.NewRGBA(r)
	draw.Draw(pbar, r, image.White, image.Point{}, draw.Src)
	draw.Draw(pbar, image.Rect(0, 0, 100, 100), image.Black, image.Point{}, draw.Src)
	f, err := os.Create("progressbar.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = png.Encode(f, pbar)
	if err != nil {
		log.Fatal(err)
	}
}
