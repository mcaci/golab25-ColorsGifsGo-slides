package conway

func GameOfLife(matrix [][]int) [][]int {
	height := len(matrix)
	if height == 0 {
		return nil
	}
	width := len(matrix[0])
	newMatrix := make([][]int, height)
	for i := range newMatrix {
		newMatrix[i] = make([]int, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			liveNeighbors := countLiveNeighbors(matrix, x, y, width, height)
			if matrix[y][x] == 1 { // Cell is alive
				if liveNeighbors < 2 || liveNeighbors > 3 {
					newMatrix[y][x] = 0 // Cell dies
				} else {
					newMatrix[y][x] = 1 // Cell lives
				}
			} else { // Cell is dead
				if liveNeighbors == 3 {
					newMatrix[y][x] = 1 // Cell becomes alive
				} else {
					newMatrix[y][x] = 0 // Cell remains dead
				}
			}
		}
	}

	return newMatrix
}

func countLiveNeighbors(matrix [][]int, x, y, width, height int) int {
	liveCount := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue // Skip the cell itself
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < width && ny >= 0 && ny < height && matrix[ny][nx] == 1 {
				liveCount++
			}
		}
	}
	return liveCount
}