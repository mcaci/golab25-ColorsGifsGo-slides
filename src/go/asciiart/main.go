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
	text := flag.Args()[0]
	fig := figure.NewFigure(text, figlet, true)
	fig.Print()
	asciiArtLines := prepareText(text, figlet)
	const (
		lImg, hImg = 4500, 600
		fontSize   = 24.0
		fontPath   = "./Ubuntu-R.ttf"
		xPtFactor  = 0.6
		yPtFactor  = 1.2
		bgColorHex = "0x7AA1FF"
		fgColorHex = "0x002300"
	)
	image := makePng(asciiArtLines, lImg, hImg, bgColorHex, fgColorHex, fontPath, fontSize, xPtFactor, yPtFactor)
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

func makePng(asciiArtLines []string, l, h int, bgColorHex, fgColorHex string, fontPath string, fontSize, xPtFactor, yPtFactor float64) *image.Paletted {
	img, err := setupBG(bgColorHex, l, h)
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
	bg := image.NewPaletted(image.Rect(0, 0, l, h), palette.Plan9)
	draw.Draw(bg, bg.Bounds(), image.NewUniform(c), image.Pt(0, 0), draw.Src)
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
