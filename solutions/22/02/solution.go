package y22d02

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "solutions/22/02/"

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
	smart, err := strconv.ParseBool(args[0])
	if err != nil {
		log.Panicf("invalid argument: '%s', expected bool", args[0])
	}
	fmt.Printf("Solution for smart == %t:\n", smart)
	s.solve(smart)
}

func (s *Solver) Solve1() {
	domain.RunFirstSolution(func() { s.solve(false) })
}

func (s *Solver) Solve2() {
	domain.RunSecondSolution(func() { s.solve(true) })
}

func (s *Solver) solve(smartStrategy bool) {
	var scoreSlice = [3][3]int{
		{4, 8, 3},
		{1, 5, 9},
		{7, 2, 6},
	}
	var totalScore int
	scanner := s.scanner()
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			st := strings.Split(txt, " ")
			var opp, me int
			if !smartStrategy {
				opp, me = score1(st[0], st[1])
			} else {
				opp, me = score2(st[0], st[1])
			}
			totalScore += scoreSlice[opp][me]
		}
	}
	fmt.Println(totalScore)
}

func score1(opp, me string) (x, y int) {
	switch opp {
	case "A":
		x = 0
	case "B":
		x = 1
	case "C":
		x = 2
	}
	switch me {
	case "X":
		y = 0
	case "Y":
		y = 1
	case "Z":
		y = 2
	}
	return
}

func score2(opp, me string) (x, y int) {
	switch opp {
	case "A":
		x = 0
		switch me {
		case "X":
			y = 2
		case "Y":
			y = 0
		case "Z":
			y = 1
		}
	case "B":
		x = 1
		switch me {
		case "X":
			y = 0
		case "Y":
			y = 1
		case "Z":
			y = 2
		}
	case "C":
		x = 2
		switch me {
		case "X":
			y = 1
		case "Y":
			y = 2
		case "Z":
			y = 0
		}
	}
	return
}
