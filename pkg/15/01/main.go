package y15d01

import (
	"bufio"
	"fmt"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "pkg/15/01/"

type Solver struct {
	scannerGen func() *bufio.Scanner
}

func (s *Solver) Init() domain.Solver {
	s.scannerGen = domain.InputScannerGen(dir)
	return s
}

func (s *Solver) Solve(_ ...string) {
	fmt.Println("Puzzle not abstracted for a general solution.")
}

func (s *Solver) Solve1() {
	domain.RunFirstSolution()
	s.solve1()
}

func (s *Solver) Solve2() {
	domain.RunSecondSolution()
	s.solve2()
}

func (s *Solver) solve1() {
	scanner := s.scannerGen()
	scanner.Scan()
	input := scanner.Text()
	var f int
	for _, v := range input {
		switch v {
		case 40:
			f++
		case 41:
			f--
		default:
			break
		}
	}
	fmt.Println(f)
}

func (s *Solver) solve2() {
	scanner := s.scannerGen()
	scanner.Scan()
	input := scanner.Text()
	var f, p int
	for i, v := range input {
		switch v {
		case 40:
			f++
		case 41:
			f--
		default:
			break
		}
		if f == -1 {
			p = i + 1
			break
		}
	}
	fmt.Println(p)
}
