package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"time"
)

func main() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			img.Set(x, y, cfuncT(time.Now()))
		}
	}
	f, _ := os.Create("pixset2.png")
	png.Encode(f, img)
}

func cfunc(x, y int) color.Color {
	r := math.Abs(math.Cos(float64(x))) * 50
	g := math.Abs(math.Sin(float64(y))) * 200
	b := math.Abs(math.Cos(float64(x*10-y*10))) * 50
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

func cfuncT(t time.Time) color.Color {
	b := t.Nanosecond() % 200
	g := t.Nanosecond() % 256
	return color.RGBA{B: uint8(b), G: uint8(g), A: 255}
}

func img2mat2img() {

	file, err := os.Open("gopher.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	matrix := make([][]int, height)
	for y := 0; y < height; y++ {
		row := make([]int, width)
		for x := 0; x < width; x++ {
			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
			// Represent color as a single integer (e.g., 0xRRGGBB)
			row[x] = int(c.R)<<16 | int(c.G)<<8 | int(c.B)
		}
		matrix[y] = row
	}

	// Print the matrix (or process as needed)
	out, _ := os.Create("gopher.jpg.source")
	for _, row := range matrix {
		fmt.Fprintln(out, row)
	}
	out.Close()
	// inSrc, _ := os.Open(out.Name())
	// outImgGoph, _ := os.Create("gopher-out.png")
	// png.Encode(outImgGoph, outImg)
}
