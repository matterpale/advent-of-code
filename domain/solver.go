package domain

import (
	"bufio"
	"fmt"
	"os"
)

type Solver interface {
	Init() Solver
	Solve(args ...string)
	Solve1()
	Solve2()
}

func InputScannerGen(dir string) func() *bufio.Scanner {
	return func() *bufio.Scanner {
		readFile, _ := os.Open(fmt.Sprintf("%s/input", dir))
		info, _ := readFile.Stat()
		scanner := bufio.NewScanner(readFile)
		scanner.Buffer(make([]byte, 0, info.Size()), int(info.Size())+1)
		return scanner
	}
}

func RunFirstSolution(solution func()) {
	fmt.Println("Solution 01:")
	solution()
}

func RunSecondSolution(solution func()) {
	fmt.Println("Solution 02:")
	solution()
}
