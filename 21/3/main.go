package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var input = "21/3/input"

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
	totals := make(map[int]int, 16)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			for i, char := range txt {
				if char == 49 {
					_, ok := totals[len(txt)-1-i]
					if !ok {
						totals[len(txt)-1-i] = 0
					}
					totals[len(txt)-1-i]++
				}
			}
		}

	}
	var g, e float64
	for i, t := range totals {
		if t > 500 {
			g = g + math.Pow(2, float64(i))
		} else {
			e = e + math.Pow(2, float64(i))
		}
	}
	return int(g) * int(e)
}

func solve2(scanner *bufio.Scanner) int {
	totals := make(map[int]int, 16)
	report := make([][16]bool, 1024)
	var i int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			for j, char := range txt {
				if char == 49 {
					report[i][j] = true
					_, ok := totals[j]
					if !ok {
						totals[j] = 0
					}
					totals[j]++
				}
			}
		}
		i++
	}

	var as, bs []int
	for j, col := range report {
		if totals[j] > 500 {
			for k, b := range col {
				if b {
					as = append(as, k)
				} else {
					bs = append(bs, k)
				}
			}
		}
	}
	return 0
}
