package y22d4

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "pkg/22/04/"

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
	partial, err := strconv.ParseBool(args[0])
	if err != nil {
		log.Panicf("invalid argument: '%s', expected bool", args[0])
	}
	fmt.Printf("Solution for partial == %t:\n", partial)
	s.solve(partial)
}

func (s *Solver) Solve1() {
	fmt.Printf("Solution 01:\n%d\n", s.solve(false))
}

func (s *Solver) Solve2() {
	fmt.Printf("Solution 02:\n%d\n", s.solve(true))
}

func (s *Solver) solve(partial bool) int {
	var total int
	scanner := s.scanner()
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			split := strings.Split(txt, ",")
			iv1 := interval(split[0])
			iv2 := interval(split[1])
			if iv1[0] == iv2[0] || iv1[1] == iv2[1] {
				total++
			} else if iv1[0] < iv2[0] {
				if partial && iv1[1] >= iv2[0] {
					total++
				} else if iv1[1] > iv2[1] {
					total++
				}
			} else if partial && iv1[0] <= iv2[1] {
				total++
			} else if iv1[1] < iv2[1] {
				total++
			}
		}
	}
	return total
}

func interval(rg string) []int {
	s := strings.Split(rg, "-")
	i1, err := strconv.Atoi(s[0])
	if err != nil {
		return nil
	}
	i2, err := strconv.Atoi(s[1])
	if err != nil {
		return nil
	}
	return []int{i1, i2}
}
