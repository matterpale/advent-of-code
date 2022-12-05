package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("22/4/input")
	if err != nil {
		panic(err.Error())
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	//	fmt.Printf("Solution 1: %d\n", solve1(fileScanner))
	fmt.Printf("Solution 2: %d\n", solve2(fileScanner))
}

func solve1(scanner *bufio.Scanner) int {
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
				if iv1[1] > iv2[0] {
					total++
				}
			} else if iv1[1] < iv2[1] {
				total++
			}
		}
	}
	return total
}

func solve2(scanner *bufio.Scanner) int {
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
				if iv1[1] >= iv2[0] {
					total++
				}
			} else if iv1[0] <= iv2[1] {
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
