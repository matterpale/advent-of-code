package y22d8

import (
	"bufio"
	"fmt"

	"adventofcode/domain"
)

const (
	dir      = "pkg/22/08/"
	gridSize = 99
)

type Solver struct {
	scanner func() *bufio.Scanner
}

func (s *Solver) Init() domain.Solver {
	s.scanner = domain.InputScannerGen(dir)
	return s
}

func (s *Solver) Solve(_ ...string) {
	fmt.Println("Puzzle not abstracted for a general solution.")
}

func (s *Solver) Solve1() {
	fmt.Printf("Solution 01:\n%d\n", s.solve(false))
}

func (s *Solver) Solve2() {
	fmt.Printf("Solution 02:\n%d\n", s.solve(true))
}

func (s *Solver) solve(second bool) int {
	grid := [gridSize][gridSize]byte{}
	scanner := s.scanner()
	for i := 0; scanner.Scan(); i++ {
		bytes := scanner.Bytes()
		if len(bytes) != 0 {
			for j, b := range bytes {
				grid[i][j] = b - 47
			}
		}
	}

	if !second {
		visibleGrid := visibleFromOutside(grid)
		var visibleCount int
		for _, row := range visibleGrid {
			for _, val := range row {
				if val {
					visibleCount++
				}
			}
		}
		return visibleCount
	}

	scoreGrid := scenicScoreGrid(grid)
	var max int
	for _, row := range scoreGrid {
		for _, score := range row {
			if score > max {
				max = score
			}
		}
	}
	return max
}
