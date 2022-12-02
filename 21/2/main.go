package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("21/2/input")
	if err != nil {
		panic(err.Error())
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// fmt.Println(solve1(fileScanner))
	fmt.Println(solve2(fileScanner))
}

func solve1(scanner *bufio.Scanner) int {
	var prev int
	incs := -1
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("atoi")
		}
		if n > prev {
			incs++
		}
		prev = n
	}
	return incs
}

func solve2(scanner *bufio.Scanner) int {
	win := [4]int{}
	var i, j, incs int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("atoi")
		}
		win[i] = n
		if j > 2 {
			old, nex := comp(win, i)
			if old < nex {
				incs++
			}
		}
		i = (i + 1) % 4
		j++
	}
	return incs
}

func comp(win [4]int, i int) (int, int) {
	return win[(i+1)%4] + win[(i+2)%4] + win[(i+3)%4], win[i] + win[(i+2)%4] + win[(i+3)%4]
}
