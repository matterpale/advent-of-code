package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("15/1/input")
	if err != nil {
		panic(err.Error())
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	fmt.Println(solve1(fileScanner.Text()))
	fmt.Println(solve2(fileScanner.Text()))
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
