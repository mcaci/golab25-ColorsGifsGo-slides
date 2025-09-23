package main

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	pbn2img()
}

func pbn2img() {
	input, err := os.Open("monalisaPBN.source")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	// Read the PNB data and convert it to an image
	var matrix [][]int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var row []int
		line := scanner.Text()
		for num := range strings.FieldsSeq(line) {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, n)
		}
		matrix = append(matrix, row)
	}

	height := len(matrix)
	if height == 0 {
		log.Fatal("Empty matrix")
	}
	width := len(matrix[0])
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y, row := range matrix {
		for x, val := range row {
			r := uint8((val >> 16) & 0xFF)
			g := uint8((val >> 8) & 0xFF)
			b := uint8(val & 0xFF)
			img.Set(x, y, color.NRGBA{R: r, G: g, B: b, A: 255})
		}
	}
	output, err := os.Create("monalisaPBN.png")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	png.Encode(output, img)
}
