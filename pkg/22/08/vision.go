package y22d8

func visibleFromOutside(grid [gridSize][gridSize]byte) (visibleGrid [gridSize][gridSize]bool) {
	for i, row := range grid {
		var top byte
		for j, spot := range row {
			if spot > top {
				visibleGrid[i][j] = true
				top = spot
			}
		}
		top = 0
		for j := len(row) - 1; j >= 0; j-- {
			if grid[i][j] > top {
				visibleGrid[i][j] = true
				top = grid[i][j]
			}
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		var top byte
		for j := 0; j < len(grid); j++ {
			if grid[j][i] > top {
				visibleGrid[j][i] = true
				top = grid[j][i]
			}
		}
		top = 0
		for j := len(grid) - 1; j >= 0; j-- {
			if grid[j][i] > top {
				visibleGrid[j][i] = true
				top = grid[j][i]
			}
		}
	}
	return
}

func scenicScoreGrid(grid [gridSize][gridSize]byte) (scoreGrid [gridSize][gridSize]int) {
	for i := range grid {
		for j := range grid[i] {
			scoreGrid[i][j] = treeScenicScore(grid, i, j)
		}
	}
	return
}

func treeScenicScore(grid [gridSize][gridSize]byte, i, j int) int {
	height := grid[i][j]
	var u, d, l, r int
	up, down, left, right := i == 0, i == gridSize-1, j == 0, j == gridSize-1
	for n := 1; ; n++ {
		if !up {
			u++
			if grid[i-n][j] >= height || i-n == 0 {
				up = true
			}
		}
		if !down {
			d++
			if grid[i+n][j] >= height || i+n == gridSize-1 {
				down = true
			}
		}
		if !left {
			l++
			if grid[i][j-n] >= height || j-n == 0 {
				left = true
			}
		}
		if !right {
			r++
			if grid[i][j+n] >= height || j+n == gridSize-1 {
				right = true
			}
		}
		if up && down && left && right {
			break
		}
	}
	return u * d * l * r
}
