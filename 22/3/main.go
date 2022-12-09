package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

const (
	abcLen = 26
	input  = "22/3/input"
)

func inputScanner() *bufio.Scanner {
	readFile, _ := os.Open(input)
	info, _ := readFile.Stat()
	scanner := bufio.NewScanner(readFile)
	scanner.Buffer(make([]byte, 0, info.Size()), 1)
	return scanner
}

func main() {
	fmt.Printf("Solution 1: %d\n", solve1(inputScanner()))
	fmt.Printf("Solution 2: %d\n", solve2(inputScanner()))
}

func solve1(scanner *bufio.Scanner) int {
	var total int
	for scanner.Scan() {
		bytes := scanner.Bytes()
		if len(bytes) != 0 {
			half := len(bytes) / 2
			f, s := bytes[:half], bytes[half:]
			for _, char := range f {
				if slices.Contains(s, char) {
					total += priority(char)
					break
				}
			}
		}
	}
	return total
}

func solve2(scanner *bufio.Scanner) int {
	var total int
	var group [][]byte
	for scanner.Scan() {
		group = append(group, scanner.Bytes())
		if len(group) < 3 {
			continue
		}
		for _, char := range group[0] {
			if slices.Contains(group[1], char) && slices.Contains(group[2], char) {
				total += priority(char)
				break
			}
		}
		group = nil
	}
	return total
}

func priority(b byte) int {
	val := b - []byte("A")[0] + 1
	if val <= abcLen {
		val = val + abcLen
	} else {
		val = val - 32
	}
	return int(val)
}
