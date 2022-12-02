package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("2/input")
	if err != nil {
		panic(err.Error())
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	scoreSlice := [3][3]int{
		{4, 8, 3},
		{1, 5, 9},
		{7, 2, 6},
	}
	var totalScore1, totalScore2 int
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if txt != "" {
			st := strings.Split(txt, " ")
			opp, me := score1(st[0], st[1])
			totalScore1 += scoreSlice[opp][me]
			opp, me = score2(st[0], st[1])
			totalScore2 += scoreSlice[opp][me]
		}
	}
	fmt.Printf("Total score 1: %d\n", totalScore1)
	fmt.Printf("Total score 2: %d", totalScore2)
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
