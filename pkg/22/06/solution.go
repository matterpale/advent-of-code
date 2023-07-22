package y22d06

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "pkg/22/06/"

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
	windowSize, err := strconv.Atoi(args[0])
	if err != nil {
		log.Panicf("invalid argument: '%s', expected int", args[0])
	}
	fmt.Printf("Solution for window size == %d:\n", windowSize)
	s.solve(windowSize)
}

func (s *Solver) Solve1() {
	fmt.Printf("Solution 01:\n%d\n", s.solve(4))
}

func (s *Solver) Solve2() {
	fmt.Printf("Solution 02:\n%d\n", s.solve(14))
}

func (s *Solver) solve(windowSize int) int {
	var position int
	scanner := s.scanner()
	if scanner.Scan() {
		txt := scanner.Text()
		for i := 0; i < len(txt)-windowSize; i++ {
			window := txt[i : i+windowSize]
			set := make(map[int32]bool, windowSize)
			foundIt := true
			for _, char := range window {
				if set[char] {
					foundIt = false
					break
				}
				set[char] = true
			}
			if foundIt {
				position = i + windowSize
				break
			}
		}
	}
	return position
}
