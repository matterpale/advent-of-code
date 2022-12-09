package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const input = "21/2/input"

func inputScanner() *bufio.Scanner {
	readFile, _ := os.Open(input)
	info, _ := readFile.Stat()
	scanner := bufio.NewScanner(readFile)
	scanner.Buffer(make([]byte, 0, info.Size()), 0)
	return scanner
}

func main() {
	fmt.Println(solve1(inputScanner()))
	fmt.Println(solve2(inputScanner()))
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
