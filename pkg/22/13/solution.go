package y22d13

import (
	"bufio"
	"fmt"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "pkg/22/13/"

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
	domain.RunFirstSolution(func() { s.solve(false) })
}

func (s *Solver) Solve2() {
	domain.RunSecondSolution(func() { s.solve(true) })
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
