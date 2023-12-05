package y22d14

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "solutions/22/14/"

const size = 1024

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
	domain.RunFirstSolution()
	s.solve(false)
}

func (s *Solver) Solve2() {
	domain.RunSecondSolution()
	s.solve(true)
}

func (s *Solver) solve(_ bool) {
	scanner := s.scanner()
	cave := newCave()
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			split1 := strings.Split(txt, " -> ")
			var prev string
			for _, st := range split1 {
				if prev == "" {
					prev = st
					continue
				}
				x1, y1 := getCoords(prev)
				x2, y2 := getCoords(st)
				if x1 == x2 {
					for i := y1; i <= y2; i++ {
						cave[x1][i] = true
					}
				} else {
					for i := x1; i <= x2; i++ {
						cave[i][y1] = true
					}
				}
				prev = st
			}
		}
	}
	total := 0
	cave.pourSand(&total)

	fmt.Println(total)
}

type Cave [][]bool

func newCave() Cave {
	columns := make(Cave, size)
	for i := range columns {
		columns[i] = make([]bool, size)
	}
	return columns
}

func getCoords(str string) (int, int) {
	spl1 := strings.Split(str, ",")
	if len(spl1) != 2 {
		panic("oops, invalid coords")
	}
	x, _ := strconv.Atoi(spl1[0])
	y, _ := strconv.Atoi(spl1[1])
	return x, y
}

func (c *Cave) pourSand(total *int) {
	var compare int
	for {
		for i := 1; i < size; i++ {
			if !(*c)[500][i] {
				continue
			}
			c.placeGrain(500, i-1, total)
		}
		if *total == compare {
			break
		}
		compare = *total
	}
}

func (c *Cave) placeGrain(x, y int, total *int) {
	if (x >= 0 && x <= size-1 && y < size-1) && !(*c)[x][y+1] {
		c.placeGrain(x, y+1, total)
	} else if (x > 1 && y < size-1) && !(*c)[x-1][y+1] {
		c.placeGrain(x-1, y+1, total)
	} else if (x < size-1 && y < size-1) && !(*c)[x+1][y+1] {
		c.placeGrain(x+1, y+1, total)
	}
	(*c)[x][y] = true
	*total++
}
