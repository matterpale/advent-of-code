package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = "22/4/input"

func main() {
	readFile1, _ := os.Open(input)
	readFile2, _ := os.Open(input)

	fileScanner1 := bufio.NewScanner(readFile1)
	fileScanner1.Split(bufio.ScanLines)
	fmt.Printf("Solution 1: %d\n", solve(fileScanner1, false))

	fileScanner2 := bufio.NewScanner(readFile2)
	fileScanner2.Split(bufio.ScanLines)
	fmt.Printf("Solution 2: %d\n", solve(fileScanner2, true))
}

func solve(scanner *bufio.Scanner, partial bool) int {
	var total int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			split := strings.Split(txt, ",")
			iv1 := interval(split[0])
			iv2 := interval(split[1])
			if iv1[0] == iv2[0] || iv1[1] == iv2[1] {
				total++
			} else if iv1[0] < iv2[0] {
				if partial && iv1[1] >= iv2[0] {
					total++
				} else if iv1[1] > iv2[1] {
					total++
				}
			} else if partial && iv1[0] <= iv2[1] {
				total++
			} else if iv1[1] < iv2[1] {
				total++
			}
		}
	}
	return total
}

func interval(rg string) []int {
	s := strings.Split(rg, "-")
	i1, err := strconv.Atoi(s[0])
	if err != nil {
		return nil
	}
	i2, err := strconv.Atoi(s[1])
	if err != nil {
		return nil
	}
	return []int{i1, i2}
}
