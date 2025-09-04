package imgconv

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func WebsafePixset() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewPaletted(r, palette.WebSafe)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			b := float64(x) / 1024 * 255
			g := float64(y) / 1024 * 255
			img.Set(x, y, color.RGBA{B: uint8(b), G: uint8(g), A: 255})
		}
	}
	f, _ := os.Create("pixsetWebSafe.png")
	png.Encode(f, img)
}

func NicePixset() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			b := float64(x) / float64(r.Max.X) * 255
			g := float64(y) / float64(r.Max.Y) * 255
			img.Set(x, y, color.RGBA{B: uint8(b), G: uint8(g), A: 255})
		}
	}
	f, _ := os.Create("pixset.png")
	png.Encode(f, img)

}

func Rect() {
	// var img image.RGBA
	r := image.Rect(0, 0, 640, 480)
	img := image.NewRGBA(r)
	for i := range r.Max.X {
		for j := range r.Max.Y {
			img.Set(i, j, color.RGBA{B: 150, A: 255})
		}
	}
	f, _ := os.Create("blue.png")
	// png.Encode(f, img)
	var magic color.Palette = []color.Color{
		color.RGBA{G: 255, A: 255},
		color.RGBA{R: 255, A: 255},
		color.RGBA{B: 255, A: 255},
		color.RGBA{G: 255, R: 255, B: 255, A: 255},
		color.RGBA{A: 255},
	}
	magic = palette.Plan9
	img2 := image.NewPaletted(image.Rect(0, 0, 640, 480), magic)
	for i := range img2.Bounds().Max.X {
		for j := range img2.Bounds().Max.Y {
			img2.Set(i, j, color.Gray{Y: 100})
		}
	}
	png.Encode(f, img)
}

func SpookyGopher() {
	f, err := os.Open("gopher.jpg")
	if err != nil {
		log.Fatal(err)
	}
	gopher, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	outGopher := image.NewRGBA64(gopher.Bounds())
	for i := range gopher.Bounds().Max.X {
		for j := range gopher.Bounds().Max.Y {
			r, g, b, a := gopher.At(i, j).RGBA()
			log.Print(r, g, b, a)
			if r > 60000 && g > 60000 && b > 60000 {
				outGopher.Set(i, j, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 0})
				continue
			}
			outGopher.Set(i, j, gopher.At(i, j))
		}
	}
	out, _ := os.Create("out.png")
	png.Encode(out, outGopher)
}
