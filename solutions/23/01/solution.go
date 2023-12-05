package y23d01

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/matterpale/advent-of-code/domain"
	"github.com/porfirion/trie"
)

const dir = "solutions/23/01/"

type Solver struct {
	scanner func() *bufio.Scanner
}

func (s *Solver) Init() domain.Solver {
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
	var first, last, calibrationSum int
	var firstSet bool
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		firstSet = false
		for _, char := range txt {
			if i, err := strconv.Atoi(string(char)); err == nil {
				if !firstSet {
					first = i
					firstSet = true
				}
				last = i
			}
		}
		calibrationValue := 10*first + last
		calibrationSum += calibrationValue
	}
	fmt.Println(calibrationSum)
}

func (s *Solver) solve2() {
	scanner := s.scanner()
	var first, last, calibrationSum int
	var firstSet bool
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		firstSet = false
		for i, char := range txt {
			if digit, err := strconv.Atoi(string(char)); err == nil {
				if !firstSet {
					first = digit
					firstSet = true
				}
				last = digit
			}
			rest := txt[i:]
			if digit, _, ok := digits.SearchPrefixInString(rest); ok {
				if !firstSet {
					first = digit
					firstSet = true
				}
				last = digit
			}
		}
		calibrationValue := 10*first + last
		calibrationSum += calibrationValue
	}
	fmt.Println(calibrationSum)
}

var digits = trie.BuildFromMap[int](map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
})
