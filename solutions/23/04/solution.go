package y23d04

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/matterpale/advent-of-code/domain"
	"golang.org/x/exp/slices"
)

const dir = "solutions/23/04/" // TODO: Specify!

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
	var n int
	for scanner.Scan() {
		txt := scanner.Text()
		card := strings.Split(strings.Split(txt, ":")[1], "|")
		winners, drawn := strings.Split(card[0], " "), strings.Split(card[1], " ")
		cardScore := 1
		var win []int
		for _, w := range winners {
			if w == "" {
				continue
			}
			x, _ := strconv.Atoi(w)
			win = append(win, x)
		}
		for _, d := range drawn {
			if d == "" {
				continue
			}
			x, _ := strconv.Atoi(d)
			if slices.Contains(win, x) {
				cardScore <<= 1
			}
		}
		if cardScore != 1 {
			n += cardScore >> 1
		}
	}
	fmt.Println(n)
}

func (s *Solver) solve2() {
	scanner := s.scanner()
	sc := &scratchcards{}
	sc.add(1, 1)
	var limit int
	for scanner.Scan() {
		limit++
		txt := scanner.Text()
		card := strings.Split(strings.Split(txt, ":")[1], "|")
		winners, drawn := strings.Split(card[0], " "), strings.Split(card[1], " ")
		var win []int
		for _, w := range winners {
			if w == "" {
				continue
			}
			x, _ := strconv.Atoi(w)
			win = append(win, x)
		}
		var wins int
		for _, d := range drawn {
			if d == "" {
				continue
			}
			x, _ := strconv.Atoi(d)
			if slices.Contains(win, x) {
				wins++
			}
		}
		sc.add(wins, sc.copies[sc.head])
		sc.add(1, 1)
		sc.head++
	}
	var sum int
	for i := 0; i < limit; i++ {
		sum += sc.copies[i]
	}
	fmt.Println(sum)
}

type scratchcards struct {
	copies []int
	head   int
}

func (s *scratchcards) add(n, k int) {
	for i := 1; i <= n; i++ {
		if s.head+i >= len(s.copies) {
			s.copies = append(s.copies, k)
		} else {
			s.copies[s.head+i] += k
		}
	}
}
