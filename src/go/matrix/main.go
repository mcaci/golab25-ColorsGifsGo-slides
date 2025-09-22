package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"

	"github.com/golang/freetype"
)

func main() {
	f, err := os.Open("what-if-matrix-template.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	var dst draw.Image
	switch outImg := img.(type) {
	case draw.Image:
		dst = outImg
	case *image.YCbCr:
		rgbaImg := image.NewRGBA(outImg.Bounds())
		for x := range rgbaImg.Rect.Dx() {
			for y := range rgbaImg.Rect.Dy() {
				rgbaImg.Set(x, y, outImg.At(x, y))
			}
		}
		dst = rgbaImg
	default:
		log.Fatalf("img %q is not of type draw.Image but %T", f.Name(), img)
	}

	// Do something with the decoded image
	fCtx := freetype.NewContext()
	fontBytes, err := os.ReadFile("./Ubuntu-R.ttf")
	if err != nil {
		log.Fatal(err)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	fCtx.SetDPI(72)
	fCtx.SetFont(font)
	const fontSize = 72
	fCtx.SetFontSize(fontSize)
	fCtx.SetClip(dst.Bounds())
	fCtx.SetDst(dst)
	fCtx.SetSrc(image.NewUniform(color.RGBA{255, 255, 255, 255}))
	textXOffset := 35
	// to have it in the bottom of the slide, use:
	textYOffset := 500 + int(fCtx.PointToFixed(fontSize)>>6) // Note shift/truncate 6 bits first
	// to have it at the bottom of the image, use:
	// textYOffset := 600 + int(fCtx.PointToFixed(fontSize)>>6) // Note shift/truncate 6 bits first

	pt := freetype.Pt(textXOffset, textYOffset)
	fCtx.DrawString("Welcome to the Matrix", pt)

	outFile, err := os.Create("welcome-to-the-matrix.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, dst, &jpeg.Options{Quality: 95})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote output to %s", outFile.Name())
}
