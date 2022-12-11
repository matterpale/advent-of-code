package y22d3

import (
	"bufio"
	"fmt"

	"github.com/matterpale/advent-of-code/domain"

	"golang.org/x/exp/slices"
)

const (
	dir    = "pkg/22/03/"
	abcLen = 26
)

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
	var total int
	scanner := s.scanner()
	for scanner.Scan() {
		bytes := scanner.Bytes()
		if len(bytes) != 0 {
			half := len(bytes) / 2
			f, s := bytes[:half], bytes[half:]
			for _, char := range f {
				if slices.Contains(s, char) {
					total += priority(char)
					break
				}
			}
		}
	}
	fmt.Printf("Solution 01:\n%d\n", total)
}

func (s *Solver) Solve2() {
	var total int
	var triad [][]byte
	scanner := s.scanner()
	for scanner.Scan() {
		triad = append(triad, scanner.Bytes())
		if len(triad) < 3 {
			continue
		}
		for _, char := range triad[0] {
			if slices.Contains(triad[1], char) && slices.Contains(triad[2], char) {
				total += priority(char)
				break
			}
		}
		triad = nil
	}
	fmt.Printf("Solution 02:\n%d\n", total)
}

func priority(b byte) int {
	val := b - []byte("A")[0] + 1
	if val <= abcLen {
		val = val + abcLen
	} else {
		val = val - 32
	}
	return int(val)
}
