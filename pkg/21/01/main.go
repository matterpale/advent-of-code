package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input = "21/02/input"

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
