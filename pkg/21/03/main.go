package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var input = "21/03/input"

func inputScanner() *bufio.Scanner {
	readFile, _ := os.Open(input)
	info, _ := readFile.Stat()
	scanner := bufio.NewScanner(readFile)
	scanner.Buffer(make([]byte, 0, info.Size()), 0)
	return scanner
}

func main() {
	fmt.Printf("Solution 01: %d\n", solve1(inputScanner()))
	fmt.Printf("Solution 02: %d\n", solve2(inputScanner()))
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
	report := make([][]bool, 1024)
	for i := 0; scanner.Scan(); i++ {
		txt := scanner.Text()
		if txt != "" {
			for _, char := range txt {
				report[i] = append(report[i], char == 49)
			}
		}
	}

	oxyRep := make([][]bool, 1024)
	co2Rep := make([][]bool, 1024)
	copy(oxyRep, report)
	copy(co2Rep, report)

	oxyInd := handleReport(oxyRep, false)
	co2Ind := handleReport(co2Rep, true)

	var oxy, co2 float64
	for i := range report[oxyInd] {
		if report[oxyInd][len(report[oxyInd])-1-i] {
			oxy = oxy + math.Pow(2, float64(i))
		}
	}

	for i := range report[co2Ind] {
		if report[co2Ind][len(report[co2Ind])-1-i] {
			co2 = co2 + math.Pow(2, float64(i))
		}
	}
	return int(oxy * co2)
}

func handleReport(report [][]bool, co2 bool) int {
	var index int
	for k := 0; k < 12; k++ {
		mostly1, last2 := isColumnMostlyOnes(report)
		for i, row := range report {
			if row != nil {
				if (mostly1 == row[0]) != co2 {
					report[i] = row[1:]
					if last2 {
						index = i
						break
					}
					continue
				}
				report[i] = nil
			}
		}
	}
	return index
}

func isColumnMostlyOnes(report [][]bool) (bool, bool) {
	var yes, no, n int
	for _, row := range report {
		if row != nil {
			if row[0] {
				yes++
			} else {
				no++
			}
			n++
		}
	}
	return yes >= no, n == 2
}
