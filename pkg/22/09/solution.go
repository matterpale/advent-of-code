package y22d9

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"adventofcode/domain"
	"adventofcode/pkg/22/09/ropes"
)

const dir = "pkg/22/09/"

type Solver struct {
	scanner func() *bufio.Scanner
}

func (s *Solver) Init() domain.Solver {
	s.scanner = domain.InputScannerGen(dir)
	return s
}

func (s *Solver) Solve(args ...string) {
	if len(args) != 1 {
		log.Panicf("invalid number of arguments: %d", len(args))
	}
	knots, err := strconv.Atoi(args[0])
	if err != nil {
		log.Panicf("invalid argument: '%s', expected int", args[0])
	}
	if knots < 2 {
		log.Panicf("invalid argument: at least 02 knots needed; got %d", knots)
	}
	fmt.Printf("Solution for %d knots:\n", knots)
	s.solve(uint(knots))
}

func (s *Solver) Solve1() {
	fmt.Printf("Solution 01:\n%d\n", s.solve(2))
}

func (s *Solver) Solve2() {
	fmt.Printf("Solution 02:\n%d\n", s.solve(10))
}

func (s *Solver) solve(knots uint) int {
	scanner := s.scanner()
	tailVisited := make(map[ropes.Position]bool)
	rope := ropes.NewRope(knots)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			split := strings.Split(txt, " ")
			d := split[0]
			steps, _ := strconv.Atoi(split[1])
			for i := 0; i < steps; i++ {
				rope.MoveHead(ropes.Direction(d))
				tailVisited[rope.TailPosition()] = true
			}
		}
	}
	return len(tailVisited)
}
