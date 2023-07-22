package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/matterpale/advent-of-code/domain"
	"github.com/matterpale/advent-of-code/internal/implemented"
)

// TODO: perhaps make it a CLI?

func main() {
	args := os.Args
	argCount := len(args) - 1
	switch argCount {
	case 0:
		allSolutions(implemented.InitAllSolvers())
		return
	case 1:
		log.Panicf("none or at least 2 args expected (year, day); got: %s", args[1])
	default:
		solver := puzzleSolver(args[1], args[2])
		// Solutions to the 2 predefined puzzles
		if argCount == 2 {
			solver.Solve1()
			solver.Solve2()
			return
		}
		// Parameterized solution, if applicable
		solver.Solve(args[3:]...)
	}
}

func puzzleSolver(year, day string) domain.Solver {
	puzzle, err := getPuzzleString(year, day)
	if err != nil {
		log.Panic(err.Error())
	}
	solver, ok := implemented.AllSolvers[puzzle]
	if !ok {
		log.Panicf("unimplemented solver: %s", puzzle)
	}
	printPuzzleHead(puzzle)
	return solver.Init()
}

func allSolutions(solvers implemented.Solvers) {
	for _, puzzle := range sortedKeys(solvers) {
		printPuzzleHead(puzzle)
		solvers[puzzle].Solve1()
		solvers[puzzle].Solve2()
		fmt.Println()
	}
}

func sortedKeys(s implemented.Solvers) []string {
	keys := make([]string, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

func printPuzzleHead(puzzle string) {
	fmt.Printf("\tPuzzle %s:\n", puzzle)
}
