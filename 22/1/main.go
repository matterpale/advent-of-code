package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var input = "22/1/input"

func main() {
	readFile, _ := os.Open(input)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	top, top3 := solve(fileScanner)
	fmt.Printf("Solution 1: %d\n", top)
	fmt.Printf("Solution 2: %d\n", top3)
}

func solve(scanner *bufio.Scanner) (top, top3 int) {
	elves := []int{0}
	var i int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			i++
			elves = append(elves, 0)
			continue
		}
		n, _ := strconv.Atoi(scanner.Text())
		elves[i] += n
	}

	var max1, max2, max3 int
	for _, v := range elves {
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}
	return max1, max1 + max2 + max3
}
