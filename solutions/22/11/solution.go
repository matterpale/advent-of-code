package y22d11

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"github.com/matterpale/advent-of-code/domain"

	"golang.org/x/exp/slices"
)

const dir = "solutions/22/11/"

type Solver struct {
	scanner func() *bufio.Scanner
}

func (s *Solver) Init() domain.Solver {
	s.scanner = domain.InputScannerGen(dir)
	return s
}

func (s *Solver) Solve(args ...string) {
	if len(args) != 2 {
		log.Panicf("invalid number of arguments: %d", len(args))
	}
	rounds, err := strconv.Atoi(args[0])
	if err != nil {
		log.Panicf("not a number of rounds: %s", args[0])
	}

	relief, err := strconv.ParseBool(args[1])
	if err != nil {
		log.Panicf("not a bool (relief): %s", args[1])
	}
	fmt.Printf("Solution for arg == %s:\n", args[0])
	s.solve(rounds, relief)
}

func (s *Solver) Solve1() {
	domain.RunFirstSolution(func() { s.solve(20, true) })
}

func (s *Solver) Solve2() { // TODO: Fix the solver, somehow only works for 20 rounds :(
	domain.RunSecondSolution(func() { s.solve(10_000, false) })
}

func (s *Solver) solve(rounds int, relief bool) {
	monkeys := initMonkeys(s.scanner(), relief)
	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			m.throwAll()
		}
	}

	var inspections []int
	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}
	slices.Sort(inspections)
	fmt.Println(inspections[len(inspections)-1] * inspections[len(inspections)-2])
}
