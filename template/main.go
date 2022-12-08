package main

import (
	"bufio"
	"fmt"
	"os"
)

var input = "year/day/input" // TODO: specify!

func main() {
	fmt.Printf("Solution 1: %d\nSolution 2: %d",
		solve(inputScanner(), false),
		solve(inputScanner(), true))
}

func solve(scanner *bufio.Scanner, second bool) int {
	var total int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			if !second {
				total++
				continue
			}
			total--
		}
	}
	return total
}

func inputScanner() *bufio.Scanner {
	readFile, _ := os.Open(input)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}
