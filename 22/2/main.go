package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input = "22/2/input"

var scoreSlice = [3][3]int{
	{4, 8, 3},
	{1, 5, 9},
	{7, 2, 6},
}

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

func solve(scanner *bufio.Scanner, smartStrategy bool) int {
	var totalScore int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			st := strings.Split(txt, " ")
			var opp, me int
			if !smartStrategy {
				opp, me = score1(st[0], st[1])
			} else {
				opp, me = score2(st[0], st[1])
			}
			totalScore += scoreSlice[opp][me]
		}
	}
	return totalScore
}

func score1(opp, me string) (x, y int) {
	switch opp {
	case "A":
		x = 0
	case "B":
		x = 1
	case "C":
		x = 2
	}
	switch me {
	case "X":
		y = 0
	case "Y":
		y = 1
	case "Z":
		y = 2
	}
	return
}

func score2(opp, me string) (x, y int) {
	switch opp {
	case "A":
		x = 0
		switch me {
		case "X":
			y = 2
		case "Y":
			y = 0
		case "Z":
			y = 1
		}
	case "B":
		x = 1
		switch me {
		case "X":
			y = 0
		case "Y":
			y = 1
		case "Z":
			y = 2
		}
	case "C":
		x = 2
		switch me {
		case "X":
			y = 1
		case "Y":
			y = 2
		case "Z":
			y = 0
		}
	}
	return
}
