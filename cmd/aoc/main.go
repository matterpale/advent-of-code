package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/matterpale/advent-of-code/pkg/implemented"
)

// TODO: perhaps make it a CLI?

func main() {
	args := os.Args
	solvers := implemented.GetSolvers()

	switch len(args) {
	case 1:
		allSolutions(solvers)
		return
	case 2:
		log.Panicf("none or min. 02 args expected (year, day); got: %s", args[1])
	default:
		puzzle, err := getPuzzleString(args[1], args[2])
		if err != nil {
			log.Panic(err.Error())
		}
		solver, ok := solvers[puzzle]
		if !ok {
			log.Panicf("unimplemented solver: %s", puzzle)
		}

		fmt.Printf("\tPuzzle %s\n", puzzle)
		// Solutions to the 02 predefined puzzles
		if len(args) == 3 {
			solver.Solve1()
			solver.Solve2()
			return
		}

		// Parameterized solution, if applicable
		solver.Solve(args[3:]...)
	}
}

func allSolutions(solvers implemented.Solvers) {
	for _, puzzle := range sortedKeys(solvers) {
		fmt.Printf("\tPuzzle %s:\n", puzzle)
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
