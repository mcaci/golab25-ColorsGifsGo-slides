package imgconv

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	"golang.org/x/image/bmp"
)

func ConvCurrDir() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for i := range entries {
		fname := entries[i].Name()
		log.Print("Reading " + fname)
		var img image.Image
		switch {
		case strings.HasSuffix(fname, "jpeg"):
			outName := fname[:strings.LastIndex(fname, ".")]
			os.Rename(fname, outName+".jpg")
			continue
		case strings.HasSuffix(fname, "png"):
			inFile, err := os.Open(fname)
			if err != nil {
				log.Fatal(err)
			}
			defer inFile.Close()
			img, err = png.Decode(inFile)
			if err != nil {
				log.Fatal(err)
			}
		case strings.HasSuffix(fname, "bmp"):
			inFile, err := os.Open(fname)
			if err != nil {
				log.Fatal(err)
			}
			defer inFile.Close()
			img, err = bmp.Decode(inFile)
			if err != nil {
				log.Fatal(err)
			}
		// case strings.HasSuffix(fname, "pdf"):
		// 	log.Print(fname + " not supported. Skipping")
		// 	continue
		default:
			log.Print(fname + " not supported. Skipping")
			continue
		}
		outName := fname[:strings.LastIndex(fname, ".")]
		out, err := os.Create(outName + ".jpg")
		if err != nil {
			log.Fatal(err)
		}
		err = jpeg.Encode(out, img, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
