package y22d10

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/matterpale/advent-of-code/domain"
)

const (
	dir       = "pkg/22/10/"
	crtWidth  = 40
	crtHeight = 6
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
	cycle, signalSum, x := 0, 0, 1
	scanner := s.scanner()
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			split := strings.Split(txt, " ")
			switch split[0] {
			case "noop":
				step(&cycle, &signalSum, x)
			case "addx":
				step(&cycle, &signalSum, x)
				step(&cycle, &signalSum, x)
				add, _ := strconv.Atoi(split[1])
				x = x + add
			}
		}
	}
	fmt.Printf("Solution 1:\n%d\n", signalSum)
}

func step(cycle, sum *int, x int) {
	*cycle++
	if (*cycle-20)%40 == 0 {
		*sum = *sum + (*cycle)*x
	}
}

func (s *Solver) Solve2() {
	fmt.Println("Solution 2:")
	channel := make(chan bool, crtHeight*crtWidth)
	defer close(channel)
	device := &Device{
		cpu: &CPU{
			scanner: s.scanner(),
			outChan: channel,
			x:       1,
		},
		crt: &CRT{
			inChan: channel,
		},
	}
	device.Boot()
}
