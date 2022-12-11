package main

import (
	"bufio"
	"fmt"
	"os"
)

const input = "15/01/input"

func main() {
	fmt.Println(solve1(inputScanner().Text()))
	fmt.Println(solve2(inputScanner().Text()))
}

func solve1(input string) int {
	var f int
	for _, v := range input {
		switch v {
		case 40:
			f++
		case 41:
			f--
		default:
			break
		}
	}
	return f
}

func solve2(input string) int {
	var f, p int
	for i, v := range input {
		switch v {
		case 40:
			f++
		case 41:
			f--
		default:
			break
		}
		if f == -1 {
			p = i + 1
			break
		}
	}
	return p
}

func inputScanner() *bufio.Scanner {
	readFile, _ := os.Open(input)
	info, _ := readFile.Stat()
	scanner := bufio.NewScanner(readFile)
	scanner.Buffer(make([]byte, 0, info.Size()), 0)
	return scanner
}
