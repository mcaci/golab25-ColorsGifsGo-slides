package main

import (
	"bufio"
	"image"
	"image/color"

	"image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
)

func main() {
	heatmap()
}

func heatmap() {

	in, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	scanner := bufio.NewScanner(in)
	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()
		for i := range line {
			// map characters from 0 to 9 with a color.RGBA from red to black typical of heatmaps
			switch line[i] {
			case '0':
				img.Set(r, i, color.RGBA{R: 0, G: 0, B: 255, A: 255})
			case '1':
				img.Set(r, i, color.RGBA{R: 0, G: 128, B: 255, A: 255})
			case '2':
				img.Set(r, i, color.RGBA{R: 0, G: 255, B: 255, A: 255})
			case '3':
				img.Set(r, i, color.RGBA{R: 0, G: 255, B: 128, A: 255})
			case '4':
				img.Set(r, i, color.RGBA{R: 0, G: 255, B: 0, A: 255})
			case '5':
				img.Set(r, i, color.RGBA{R: 128, G: 255, B: 0, A: 255})
			case '6':
				img.Set(r, i, color.RGBA{R: 255, G: 255, B: 0, A: 255})
			case '7':
				img.Set(r, i, color.RGBA{R: 255, G: 128, B: 0, A: 255})
			case '8':
				img.Set(r, i, color.RGBA{R: 255, G: 0, B: 0, A: 255})
			case '9':
				img.Set(r, i, color.RGBA{R: 128, G: 0, B: 0, A: 255})
			}
		}
	}
	dst := image.NewRGBA(r)
	draw.NearestNeighbor.Scale(dst, r, img, img.Bounds(), draw.Over, nil)
	f, _ := os.Create("heatmap.png")
	png.Encode(f, img)
}

func matrixToColor(matrix [][]int, x, y, w, h int) color.Color {
	mx := x * w / 1024
	my := y * h / 768
	if mx >= w {
		mx = w - 1
	}
	if my >= h {
		my = h - 1
	}
	val := matrix[my][mx]
	// Map value (0-100) to color gradient from blue (low) to red (high)
	r := uint8(val * 255 / 100)
	g := uint8(0)
	b := uint8(255 - r)
	return color.RGBA{R: r, G: g, B: b, A: 255}
}
