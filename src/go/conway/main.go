package conway

import (
	"image"
	"image/color"
	"image/gif"
	"math/rand"
	"os"
)

func Run() {
	// in := [][]int{
	// 	{0, 0, 1, 0, 0},
	// 	{1, 0, 1, 0, 0},
	// 	{0, 1, 1, 0, 0},
	// 	{0, 0, 0, 1, 0},
	// 	{0, 0, 0, 0, 1},
	// }
	const size = 100
	in := make([][]int, size)
	for i := range in {
		in[i] = make([]int, size)
		for j := range in[i] {
			if rand.Float64() < 0.1 { // ~20% chance cell is alive
				in[i][j] = 1
			}
		}
	}
	var outImgs []*image.Paletted
	for range 800 {
		// time.Sleep(1 * time.Second)
		// cmd := exec.Command("clear")
		// cmd.Stdout = os.Stdout
		// cmd.Run()
		out := GameOfLife(in)
		// for _, row := range out {
		// 	for _, cell := range row {
		// 		if cell == 1 {
		// 			fmt.Print("O ")
		// 		} else {
		// 			fmt.Print(". ")
		// 		}
		// 	}
		// 	fmt.Println()
		// }
		img := image.NewPaletted(image.Rect(0, 0, len(out[0]), len(out)), color.Palette{color.Black, color.White})
		for y, row := range out {
			for x, cell := range row {
				if cell == 1 {
					img.SetColorIndex(x, y, 1) // Set to white for live cells
				} else {
					img.SetColorIndex(x, y, 0) // Set to black for dead cells
				}
			}
		}
		outImgs = append(outImgs, img)
		in = out
	}
	g := gif.GIF{
		Image: outImgs,
	}
	for range g.Image {
		g.Delay = append(g.Delay, 15) // Set delay for each frame
	}
	f, _ := os.Create("game_of_life.gif")
	gif.EncodeAll(f, &g)
	f.Close()
}
