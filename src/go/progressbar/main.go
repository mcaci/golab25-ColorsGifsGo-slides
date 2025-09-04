package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"log"
	"os"
)

func main() {

	r := image.Rect(0, 0, 1000, 100)

	pbar := image.NewPaletted(r, color.Palette{color.Black, color.White})
	draw.Draw(pbar, r, image.White, image.Point{}, draw.Src)
	var frames []*image.Paletted
	var delay []int
	frames = append(frames, pbar)
	delay = append(delay, 50)
	for i := 5; i <= 1000; i += 5 {
		pbar := image.NewPaletted(r, color.Palette{color.Black, color.White})
		draw.Draw(pbar, r, image.White, image.Point{}, draw.Src)
		draw.Draw(pbar, image.Rect(0, 0, i, 100), image.Black, image.Point{}, draw.Src)
		frames = append(frames, pbar)
		if i == 1000 {
			delay = append(delay, 500)
			continue
		}
		delay = append(delay, 2)
	}
	g := &gif.GIF{
		Image: frames,
		Delay: delay,
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
