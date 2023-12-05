package y22d12

import (
	"bufio"
	"fmt"

	"github.com/matterpale/advent-of-code/domain"
	"golang.org/x/exp/slices"
)

const dir = "solutions/22/12/"

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

func (s *Solver) solve(short bool) int {
	scanner := s.scanner()
	var grid Grid
	var start, end [2]int
	for i := 0; scanner.Scan(); i++ {
		bytes := scanner.Bytes()
		if len(bytes) != 0 {
			grid = append(grid, []byte{})
			for j, b := range bytes {
				if b == []byte("S")[0] {
					start = [2]int{i, j}
					grid[i] = append(grid[i], []byte("a")[0])
					continue
				}
				if b == []byte("E")[0] {
					end = [2]int{i, j}
					grid[i] = append(grid[i], []byte("z")[0])
					continue
				}
				grid[i] = append(grid[i], b)
			}
		}
	}

	stop := func(current [2]int) bool {
		return current == start
	}
	if short {
		stop = func(current [2]int) bool {
			return grid[current[0]][current[1]] == []byte("a")[0]
		}
	}

	size := len(grid) * len(grid[0])
	visited, queue := make([][2]int, 0, size), make([]vertex, 0, size)
	visited, queue = append(visited, end), append(queue, vertex{
		position: end,
		steps:    0,
	})

	for ; len(queue) > 0; queue = queue[1:] {
		current := queue[0]
		for _, nei := range grid.neighbours(current.position) {
			steps := current.steps + 1
			if stop(nei) {
				return steps
			}
			if !slices.Contains(visited, nei) {
				visited, queue = append(visited, nei), append(queue, vertex{
					position: nei,
					steps:    steps,
				})
			}
		}
	}
	return 0
}
