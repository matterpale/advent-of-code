package y00d0 // Package y00d0 // TODO: Specify!

import (
	"bufio"
	"fmt"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "pkg/{year}/{day}/" // TODO: Specify!

type Solver struct {
	scanner func() *bufio.Scanner
}

func (s *Solver) Init() domain.Solver {
	// TODO: Once solved, register at internal/implemented!
	s.scanner = domain.InputScannerGen(dir)
	return s
}

func (s *Solver) Solve(_ ...string) {
	// TODO: Consider providing a more abstract solution.
	fmt.Println("Puzzle not abstracted for a general solution.")
}

func (s *Solver) Solve1() {
	fmt.Println("Solution 01:")
	s.solve1()
}

func (s *Solver) Solve2() {
	fmt.Println("Solution 02:")
	s.solve2()
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
