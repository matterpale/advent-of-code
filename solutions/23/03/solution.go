package y23d03

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "solutions/23/03/"

type Solver struct {
	scanner func() *bufio.Scanner

	sum       int
	window    [3]string
	head, row int
	emptyLine string
}

func (s *Solver) Init() domain.Solver {
	// TODO: register at cmd/aoc/done.go
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
	scanner := s.start()
	for {
		s.processLine()
		if !scanner.Scan() {
			s.nextLine(s.emptyLine)
			s.processLine()
			break
		}
		s.nextLine(scanner.Text())
	}
	fmt.Println(s.sum)
}

func (s *Solver) solve2() {
	scanner := s.start()
	for scanner.Scan() {

		s.nextLine(scanner.Text())
	}
}

func (s *Solver) start() *bufio.Scanner {
	scanner := s.scanner()
	scanner.Scan()
	line1 := scanner.Text()
	var dots []byte
	for i := 0; i < len(line1); i++ {
		dots = append(dots, []byte(".")...)
	}
	scanner.Scan()
	line2 := scanner.Text()

	s.window = [3]string{string(dots), line1, line2}
	s.emptyLine = string(dots)
	s.row = 1
	return scanner
}

func (s *Solver) nextLine(line string) {
	if line == "" {
		line = s.emptyLine
	}
	s.window[s.prevRow()] = line
	s.row = s.nextRow()
}

func (s *Solver) processLine() {
	x, line := 0, s.window[s.row]
	for {
		y := s.findAndProcessNumber(line, x)
		if y == -1 {
			break
		}
		line = line[y:]
		x += y
	}
}

var numberRegex = regexp.MustCompile(`\d+`)
var symbolRegex = regexp.MustCompile(`[^\d.]+`)

func (s *Solver) findAndProcessNumber(line string, offset int) int {
	if !numberRegex.MatchString(line) {
		return -1
	}
	number, err := strconv.Atoi(numberRegex.FindString(line))
	if err != nil {
		panic("BAD BAD NOT GOOD")
	}
	location := numberRegex.FindStringIndex(line)
	if s.isPartNumber(location[0], location[1], offset) {
		s.sum += number
	}
	return location[1]
}

func (s *Solver) isPartNumber(left, right, offset int) bool {
	neighborhood := s.getNeighborhood(left+offset, right+offset)
	for _, n := range neighborhood {
		if isSymbol(n) {
			return true
		}
	}
	return false
}

func (s *Solver) getNeighborhood(left int, right int) []string {
	var neighborhood []string
	if left > 0 {
		neighborhood = append(neighborhood, string([]byte{s.window[0][left-1], s.window[1][left-1], s.window[2][left-1]}))
	}
	if right < len(s.emptyLine) {
		neighborhood = append(neighborhood, string([]byte{s.window[0][right], s.window[1][right], s.window[2][right]}))
	}
	return append(neighborhood, s.window[s.prevRow()][left:right], s.window[s.nextRow()][left:right])
}

func isSymbol(char string) bool {
	return symbolRegex.MatchString(char)
}

func (s *Solver) nextRow() int {
	return (s.row + 1) % 3
}

func (s *Solver) prevRow() int {
	return (s.row + 2) % 3
}
