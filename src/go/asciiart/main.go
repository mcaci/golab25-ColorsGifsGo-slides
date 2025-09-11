package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"
	"strings"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

func main() {
	var figlet string
	flag.StringVar(&figlet, "f", "standard", "Figlet font to use")
	flag.Parse()
	text := "Welcome to GoLab!"
	if len(flag.Args()) > 0 {
		text = flag.Args()[0]
	}
	fig := figure.NewFigure(text, figlet, true)
	fig.Print()
	asciiArtLines := prepareText(text, figlet)
	const (
		lImg, hImg = 5600, 600
		fontSize   = 64.0
		fontPath   = "./Ubuntu-R.ttf"
		xPtFactor  = 0.6
		yPtFactor  = 1.2
		bgColorHex = "0xf4ff42" // or "0xfe7dd0" or "0x7AA1FF"
		fgColorHex = "0xf4ff42"
	)
	// For a solid background:
	// image := makePng(asciiArtLines, lImg, hImg, bgColorHex, fgColorHex, fontPath, fontSize, xPtFactor, yPtFactor, nil)
	// For a dynamic background (paint providing the function parameter):
	image := makePng(asciiArtLines, lImg, hImg, bgColorHex, fgColorHex, fontPath, fontSize, xPtFactor, yPtFactor,
		func(l, h int, img *image.Paletted) {
			for x := range l {
				for y := range h {
					var r int
					b := (1 + math.Cos(float64(x)/10)) * 40
					g := (1 + math.Sin(float64(y)/10)) * 40
					img.Set(x, y, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 100})
				}
			}
		})
	writePng(image, "asciiart.png")
}

func prepareText(text, figlet string) []string {
	fig := figure.NewFigure(text, figlet, true)
	asciiArtLines := fig.Slicify()
	var maxLineLen int
	for i := range asciiArtLines {
		if maxLineLen >= len(asciiArtLines[i]) {
			continue
		}
		maxLineLen = len(asciiArtLines[i])
	}
	for i := range asciiArtLines {
		asciiArtLines[i] += strings.Repeat(" ", maxLineLen-len(asciiArtLines[i]))
	}
	return asciiArtLines
}

func makePng(asciiArtLines []string, l, h int, bgColorHex, fgColorHex string, fontPath string, fontSize, xPtFactor, yPtFactor float64, paint func(int, int, *image.Paletted)) *image.Paletted {
	var img *image.Paletted
	var err error
	switch paint {
	case nil:
		img, err = setupBG(bgColorHex, l, h)
	default:
		img, err = setupBGFunc(bgColorHex, l, h, paint)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = drawFG(asciiArtLines, 0, 0, img, fgColorHex, fontPath, fontSize, xPtFactor, yPtFactor)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func setupBG(bgHex string, l, h int) (*image.Paletted, error) {
	c, err := parseHexColor(bgHex)
	if err != nil {
		return nil, err
	}
	return setupBGFunc(bgHex, l, h, func(l, h int, bg *image.Paletted) {
		draw.Draw(bg, bg.Bounds(), image.NewUniform(c), image.Pt(0, 0), draw.Src)
	})

}

func setupBGFunc(bgHex string, l, h int, paint func(int, int, *image.Paletted)) (*image.Paletted, error) {
	bg := image.NewPaletted(image.Rect(0, 0, l, h), palette.Plan9)
	paint(l, h, bg)
	return bg, nil
}

func drawFG(lines []string, s, e int, bg draw.Image, fgHex, fontPath string, fontSize, xPtFactor, yPtFactor float64) error {
	c, err := fgContext(bg, fgHex, fontPath, fontSize)
	if err != nil {
		return err
	}
	textXOffset := 10
	textYOffset := 30 + int(c.PointToFixed(fontSize)>>6) // Note shift/truncate 6 bits first

	pt := freetype.Pt(textXOffset, textYOffset)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		switch {
		case s < e && e < len(line):
			line = line[s:e]
		case s < e && e >= len(line):
			line = line[s:]
		}
		// log.Print(line)
		startX := pt.X
		for _, char := range line {
			_, err := c.DrawString(string(char), pt)
			if err != nil {
				return err
			}
			pt.X += c.PointToFixed(fontSize * xPtFactor)
		}
		pt.X = startX
		pt.Y += c.PointToFixed(fontSize * yPtFactor)
	}
	return nil
}

func fgContext(bg draw.Image, fgColorHex, fontPath string, fontSize float64) (*freetype.Context, error) {
	c := freetype.NewContext()
	fontBytes, err := os.ReadFile(fontPath)
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
	fgColor, err := parseHexColor(fgColorHex)
	if err != nil {
		return nil, err
	}
	c.SetSrc(image.NewUniform(fgColor))
	c.SetHinting(font.HintingNone)
	return c, nil
}

func parseHexColor(hex string) (color.RGBA, error) {
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

func writePng(image *image.Paletted, path string) {
	f := mustFile(path)
	defer f.Close()
	err := png.Encode(f, image)
	if err != nil {
		log.Fatal(err)
	}
}

func mustFile(name string) *os.File {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
