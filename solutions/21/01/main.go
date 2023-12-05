package y21d01

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "solutions/21/01/"

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
	domain.RunFirstSolution(s.solve1)
}

func (s *Solver) Solve2() {
	domain.RunSecondSolution(s.solve2)
}

func (s *Solver) solve1() {
	scanner := s.scanner()
	var prev int
	incs := -1
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("atoi")
		}
		if n > prev {
			incs++
		}
		prev = n
	}
	fmt.Println(incs)
}

func (s *Solver) solve2() {
	scanner := s.scanner()
	win := [4]int{}
	var i, j, incs int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("atoi")
		}
		win[i] = n
		if j > 2 {
			old, nex := comp(win, i)
			if old < nex {
				incs++
			}
		}
		i = (i + 1) % 4
		j++
	}
	fmt.Println(incs)
}

func comp(win [4]int, i int) (int, int) {
	return win[(i+1)%4] + win[(i+2)%4] + win[(i+3)%4], win[i] + win[(i+2)%4] + win[(i+3)%4]
}
