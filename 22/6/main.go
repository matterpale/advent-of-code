package main

import (
	"bufio"
	"fmt"
	"os"
)

var input = "22/6/input"

func main() {
	readFile1, _ := os.Open(input)
	readFile2, _ := os.Open(input)

	fileScanner1 := bufio.NewScanner(readFile1)
	fileScanner1.Split(bufio.ScanLines)
	fmt.Printf("Solution 1: %d\n", solve(fileScanner1, 4))

	fileScanner2 := bufio.NewScanner(readFile2)
	fileScanner2.Split(bufio.ScanLines)
	fmt.Printf("Solution 2: %d\n", solve(fileScanner2, 14))
}

func solve(scanner *bufio.Scanner, windowSize int) int {
	var position int
	if scanner.Scan() {
		txt := scanner.Text()
		for i := 0; i < len(txt)-windowSize; i++ {
			window := txt[i : i+windowSize]
			set := make(map[int32]bool, windowSize)
			foundIt := true
			for _, char := range window {
				if set[char] {
					foundIt = false
					break
				}
				set[char] = true
			}
			if foundIt {
				position = i + windowSize
				break
			}
		}
	}
	return position
}
