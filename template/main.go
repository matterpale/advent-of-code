package main

import (
	"bufio"
	"fmt"
	"os"
)

var input = "year/day/input" // TODO: specify!

func main() {
	readFile1, _ := os.Open(input)
	readFile2, _ := os.Open(input)

	fileScanner1 := bufio.NewScanner(readFile1)
	fileScanner1.Split(bufio.ScanLines)
	fmt.Printf("Solution 1: %d\n", solve1(fileScanner1))

	fileScanner2 := bufio.NewScanner(readFile2)
	fileScanner2.Split(bufio.ScanLines)
	fmt.Printf("Solution 2: %d\n", solve2(fileScanner2))
}

func solve1(scanner *bufio.Scanner) int {
	var total int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
		}
	}
	return total
}

func solve2(scanner *bufio.Scanner) int {
	var total int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
		}
	}
	return total
}
