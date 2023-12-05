package y23d02

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "solutions/23/02/"

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
	c := &cubeCountChecker{
		r: 12,
		g: 13,
		b: 14,
	}
	var sum, n int
	for scanner.Scan() {
		n++
		txt := scanner.Text()
		if txt == "" {
			break
		}
		game := strings.Split(txt, ":")[1]
		possible := true
		grabs := strings.Split(game, ";")
		for _, grab := range grabs {
			colors := strings.Split(grab, ",")
			for _, color := range colors {
				str := strings.Split(color, " ")
				nb, err := strconv.Atoi(str[1])
				if err != nil {
					fmt.Println("BAD BAD NOT GOOD")
				}
				if !c.ok(str[2], nb) {
					possible = false
					break
				}
			}
			if !possible {
				break
			}
		}
		if possible {
			sum += n
		}
	}
	fmt.Println(sum)
}

func (s *Solver) solve2() {
	scanner := s.scanner()
	mcp := &minCubePower{}
	var sum, n int
	for scanner.Scan() {
		n++
		txt := scanner.Text()
		if txt == "" {
			break
		}
		game := strings.Split(txt, ":")[1]
		grabs := strings.Split(game, ";")
		for _, grab := range grabs {
			colors := strings.Split(grab, ",")
			for _, color := range colors {
				str := strings.Split(color, " ")
				nb, err := strconv.Atoi(str[1])
				if err != nil {
					fmt.Println("BAD BAD NOT GOOD")
				}
				mcp.observe(str[2], nb)
			}
		}
		sum += mcp.power()
	}
	fmt.Println(sum)
}

type cubeCountChecker struct {
	r, g, b int
}

func (c *cubeCountChecker) ok(color string, number int) bool {
	switch color {
	case red:
		return number <= c.r
	case green:
		return number <= c.g
	case blue:
		return number <= c.b
	}
	return false
}

const (
	red   = "red"
	green = "green"
	blue  = "blue"
)

type minCubePower struct {
	r, g, b int
}

func (p *minCubePower) observe(color string, number int) {
	switch color {
	case red:
		if number > p.r {
			p.r = number
		}
	case green:
		if number > p.g {
			p.g = number
		}
	case blue:
		if number > p.b {
			p.b = number
		}
	}
}

func (p *minCubePower) power() int {
	power := p.r * p.g * p.b
	p.r, p.g, p.b = 0, 0, 0
	return power
}
