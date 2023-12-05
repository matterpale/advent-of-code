package y00d00

// TODO: Specify!

import (
	"bufio"
	"fmt"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "solutions/YY/DD/" // TODO: Specify!

type Solver struct {
	scanner func() *bufio.Scanner
}

func (s *Solver) Init() domain.Solver {
	// TODO: Register at cmd/aoc/done.go
	s.scanner = domain.InputScannerGen(dir)
	return s
}

func (s *Solver) Solve(_ ...string) {
	// TODO: Consider providing a more abstract solution.
	fmt.Println("Puzzle not abstracted for a general solution.")
}

func (s *Solver) Solve1() {
	domain.RunFirstSolution(s.solve1)
}

func (s *Solver) Solve2() {
	domain.RunSecondSolution(s.solve2)
}

func (s *Solver) solve1() {
	scanner := s.scanner()
	var n int
	// TODO: Implement first solution.
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			n++
		}
	}
	fmt.Println(n)
}

func (s *Solver) solve2() {
	scanner := s.scanner()
	var n int
	// TODO: Implement second solution.
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			n++
		}
	}
	fmt.Println(n)
}
