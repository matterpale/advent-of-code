package y21d02

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "pkg/21/02/"

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
	domain.RunFirstSolution(s.solve1)
}

func (s *Solver) Solve2() {
	domain.RunSecondSolution(s.solve2)
}

func (s *Solver) solve1() {
	var h, d int
	scanner := s.scanner()
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		split := strings.Split(txt, " ")
		n, err := strconv.Atoi(split[1])
		if err != nil {
			panic("atoi")
		}
		switch split[0] {
		case "forward":
			h = h + n
		case "down":
			d = d + n
		case "up":
			d = d - n
		}
	}
	fmt.Println(h * d)
}

func (s *Solver) solve2() {
	var h, d, a int
	scanner := s.scanner()
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		split := strings.Split(txt, " ")
		n, err := strconv.Atoi(split[1])
		if err != nil {
			panic("atoi")
		}
		switch split[0] {
		case "forward":
			h = h + n
			d = d + (a * n)
		case "down":
			a = a + n
		case "up":
			a = a - n
		}
	}
	fmt.Println(h * d)
}
