package y22d07

import (
	"bufio"
	"fmt"

	"github.com/matterpale/advent-of-code/domain"
)

const (
	dir      = "solutions/22/07/"
	total    = 70_000_000
	required = 30_000_000
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
	fmt.Printf("Solution 1:\n%d\n", s.solve(false))
}

func (s *Solver) Solve2() {
	fmt.Printf("Solution 2:\n%d\n", s.solve(true))
}

func (s *Solver) solve(task2 bool) int {
	scanner := s.scanner()
	root := prepFileSystem(scanner)
	used := root.fetchSize()
	if !task2 {
		return root.sumSmallDirSizesWithDuplicities()
	}

	free := total - used
	need := required - free
	var get int
	if need > 0 {
		get = root.findSmallestGarbageDirSize(used, need)
	}
	return get
}
