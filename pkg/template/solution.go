package y00d0

import (
	"bufio"
	"fmt"

	"adventofcode/domain"
)

const dir = "pkg/year/day/" // TODO: Specify

type Solver struct {
	scanner func() *bufio.Scanner
}

func (s *Solver) Init() domain.Solver {
	// TODO: Register at pkg/implemented
	s.scanner = domain.InputScannerGen(dir)
	return s
}

func (s *Solver) Solve(_ ...string) {
	fmt.Println("Puzzle not abstracted for a general solution.")
}

func (s *Solver) Solve1() {
	fmt.Println("Solution 01:")
	s.solve(false)
}

func (s *Solver) Solve2() {
	fmt.Println("Solution 02:")
	s.solve(true)
}

func (s *Solver) solve(_ bool) {
	scanner := s.scanner()
	var n int
	// TODO: Implement solution
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			n++
		}
	}
	fmt.Println(n)
}