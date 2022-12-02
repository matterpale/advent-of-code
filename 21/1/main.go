package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var h, d int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		split := strings.Split(txt, " ")
		n, err := strconv.Atoi(split[1])
		if err != nil {
			panic("atoi")
		}
		switch split[0] {
		case "forward":
			h = h + n
		case "down":
			d = d + n
		case "up":
			d = d - n
		}
	}
	return h * d
}

func solve2(scanner *bufio.Scanner) int {
	var h, d, a int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		split := strings.Split(txt, " ")
		n, err := strconv.Atoi(split[1])
		if err != nil {
			panic("atoi")
		}
		switch split[0] {
		case "forward":
			h = h + n
			d = d + (a * n)
		case "down":
			a = a + n
		case "up":
			a = a - n
		}
	}
	return h * d
}
