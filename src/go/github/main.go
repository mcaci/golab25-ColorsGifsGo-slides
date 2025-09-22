package main

import (
	"encoding/xml"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/image/draw"
)

func main() {
	readFromXml()
	tryReadFromGithub()
}

func readFromXml() {
	// Open XML data file
	f, err := os.Open("data.xml")
	if err != nil {
		log.Fatalf("failed to open data.xml: %v", err)
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	const l, h = 52, 7
	var y int
	p := color.Palette{
		MustHexToRGBA("EFF2F5"),
		MustHexToRGBA("ACEEBB"),
		MustHexToRGBA("4AC26B"),
		MustHexToRGBA("2DA44E"),
		MustHexToRGBA("116329"),
	}
	for i := range p {
		fmt.Printf("Palette[%d]: %+v\n", i, p[i])
	}
	r := image.Rect(0, 0, l, h)
	img := image.NewPaletted(r, p)

	for {
		tok, err := decoder.Token()
		if err != nil {
			break // EOF or error
		}
		se, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		if len(se.Attr) < 3 {
			continue // skip if not enough attributes
		}
		if se.Attr[0].Value == "0" {
			y++
		}
		x, err1 := strconv.Atoi(se.Attr[0].Value)
		n, err2 := strconv.Atoi(se.Attr[2].Value)
		if err1 != nil || err2 != nil || n < 0 || n >= len(p) {
			continue // skip invalid data
		}
		img.Set(x, y-1, p[n])
	}

	dstR := image.Rect(0, 0, l*20, h*20)
	dst := image.NewRGBA(dstR)
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	out, err := os.Create("contribCount.png")
	if err != nil {
		log.Fatalf("failed to create output image: %v", err)
	}
	defer out.Close()
	if err := png.Encode(out, dst); err != nil {
		log.Fatalf("failed to encode PNG: %v", err)
	}
}

// MustHexToRGBA converts a hex color string to color.RGBA, panics on error.
func MustHexToRGBA(hex string) color.RGBA {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		panic(fmt.Errorf("invalid hex color length: %q", hex))
	}
	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		panic(fmt.Errorf("invalid red value: %w", err))
	}
	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		panic(fmt.Errorf("invalid green value: %w", err))
	}
	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		panic(fmt.Errorf("invalid blue value: %w", err))
	}
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

func tryReadFromGithub() {

	// get content of this URL https://github.com/mcaci?action=show&controller=profiles&tab=contributions&user_id=mcaci
	resp, err := http.Get("https://github.com/mcaci?action=show&controller=profiles&tab=contributions&user_id=mcaci")
	if err != nil {
		log.Fatalf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	const l, h = 52, 7
	var y int
	p := color.Palette{
		MustHexToRGBA("EFF2F5"),
		MustHexToRGBA("ACEEBB"),
		MustHexToRGBA("4AC26B"),
		MustHexToRGBA("2DA44E"),
		MustHexToRGBA("116329"),
	}
	r := image.Rect(0, 0, l, h)
	img := image.NewPaletted(r, p)

	decoder := xml.NewDecoder(resp.Body)
	for {
		tok, err := decoder.Token()
		if err != nil {
			log.Print(err)
			break // EOF or error
		}
		se, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		if se.Name.Local != "table" {
			continue
		}
		if len(se.Attr) < 8 {
			continue // skip if not enough attributes
		}
		if se.Attr[0].Name.Local != "tabindex" {
			continue
		}
		if se.Attr[1].Value == "0" {
			y++
		}
		x, err1 := strconv.Atoi(se.Attr[1].Value)
		n, err2 := strconv.Atoi(se.Attr[7].Value)
		if err1 != nil || err2 != nil || n < 0 || n >= len(p) {
			log.Println("Skipping invalid data:", se.Attr[1].Value, se.Attr[7].Value)
			continue // skip invalid data
		}
		log.Println("Setting pixel:", x, y-1, "to color index", n)
		img.Set(x, y-1, p[n])
	}

	dstR := image.Rect(0, 0, l*20, h*20)
	dst := image.NewRGBA(dstR)
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	out, err := os.Create("contribCount1.png")
	if err != nil {
		log.Fatalf("failed to create output image: %v", err)
	}
	defer out.Close()
	if err := png.Encode(out, dst); err != nil {
		log.Fatalf("failed to encode PNG: %v", err)
	}
}
