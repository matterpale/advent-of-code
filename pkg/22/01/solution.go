package y22d1

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"github.com/matterpale/advent-of-code/domain"

	"golang.org/x/exp/slices"
)

const dir = "pkg/22/01/"

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
	topLimit, err := strconv.Atoi(args[0])
	if err != nil {
		log.Panicf("invalid argument: '%s', expected int", args[0])
	}
	fmt.Printf("Solution for topLimit == %d:\n", topLimit)
	s.solve(topLimit)
}

func (s *Solver) Solve1() {
	fmt.Println("Solution 01:")
	s.solve(1)
}

func (s *Solver) Solve2() {
	fmt.Println("Solution 02:")
	s.solve(3)
}

func (s *Solver) solve(topLimit int) {
	elves := sortElves(s.scanner())
	var sum int
	for i := 1; i <= topLimit; i++ {
		sum = sum + elves[len(elves)-i]
	}
	fmt.Println(sum)
}

func sortElves(scanner *bufio.Scanner) []int {
	elves := []int{0}
	var i int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			i++
			elves = append(elves, 0)
			continue
		}
		n, _ := strconv.Atoi(scanner.Text())
		elves[i] += n
	}

	slices.Sort(elves)
	return elves
}
