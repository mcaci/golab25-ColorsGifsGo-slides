package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"strings"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func main() {
	base := MustLoad("golab-speakers.png")
	gopherized := MustLoad("McaciGopherizeMe.png")
	gray := image.NewUniform(color.Gray{Y: 150})

	draw.Draw(base, image.Rect(110, 100, base.Bounds().Dx()-110, base.Bounds().Dy()-100), gray, image.Point{110, 100}, draw.Over)
	draw.Draw(base, image.Rect(130, 130, base.Bounds().Dx()-130, base.Bounds().Dy()-130), gopherized, image.Point{130, 130}, draw.Over)
	ftc, err := ftContext(base, "0x0000DE", 94.0)
	if err != nil {
		log.Fatal(err)
	}
	ftc.DrawString("Colors, images and gifs:", fixed.P(1200, 400))
	ftc.DrawString("bring on the fun with Go", fixed.P(1250, 550))
	ftc.DrawString("by Michele Caci", fixed.P(1650, 1150))

	out, _ := os.Create("composition.png")
	png.Encode(out, base)
}

func MustLoad(filename string) draw.Image {
	var decoder func(io.Reader) (image.Image, error) = png.Decode
	if strings.HasSuffix(filename, "jpg") {
		decoder = jpeg.Decode
	}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	img, err := decoder(f)
	if err != nil {
		log.Fatal(err)
	}
	switch outImg := img.(type) {
	case draw.Image:
		return outImg
	case *image.YCbCr:
		rgbaImg := image.NewRGBA(outImg.Bounds())
		for x := range rgbaImg.Rect.Dx() {
			for y := range rgbaImg.Rect.Dy() {
				rgbaImg.Set(x, y, outImg.At(x, y))
			}
		}
		return rgbaImg
	default:
		log.Fatalf("img %q is not of type draw.Image but %T", filename, img)
		return nil
	}
}

func ftContext(bg draw.Image, fgColorHex string, fontSize float64) (*freetype.Context, error) {
	c := freetype.NewContext()
	fontBytes, err := os.ReadFile("Ubuntu-R.ttf")
	if err != nil {
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}
	c.SetFont(f)
	c.SetDPI(72)
	c.SetFontSize(fontSize)
	c.SetClip(bg.Bounds())
	c.SetDst(bg)
	fgColor, err := ParseHexColor(fgColorHex)
	if err != nil {
		return nil, err
	}
	c.SetSrc(image.NewUniform(fgColor))
	c.SetHinting(font.HintingNone)
	return c, nil
}

func ParseHexColor(hex string) (color.RGBA, error) {
	var c color.RGBA
	var err error
	c.A = 0xff
	switch len(hex) {
	case 8:
		_, err = fmt.Sscanf(hex, "0x%02x%02x%02x", &c.R, &c.G, &c.B)
		return c, err
	case 5:
		_, err = fmt.Sscanf(hex, "0x%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
		return c, err
	default:
		return color.RGBA{}, fmt.Errorf("invalid length, must be 8 or 5")
	}
}
