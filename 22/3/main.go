package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

const abcLen = 26

var input = "22/3/input"

func main() {
	readFile1, _ := os.Open(input)
	readFile2, _ := os.Open(input)

	fileScanner1 := bufio.NewScanner(readFile1)
	fileScanner1.Split(bufio.ScanLines)
	fmt.Printf("Solution 1: %d\n", solve1(fileScanner1))

	fileScanner2 := bufio.NewScanner(readFile2)
	fileScanner2.Split(bufio.ScanLines)
	fmt.Printf("Solution 2: %d\n", solve2(fileScanner2))
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

func solve2(scanner *bufio.Scanner) int { // TODO: 2615 instead of 2609???
	var total, n int
	for scanner.Scan() {
		bytes1 := scanner.Bytes()
		scanner.Scan()
		bytes2 := scanner.Bytes()
		scanner.Scan()
		bytes3 := scanner.Bytes()
		if len(bytes1) != 0 {
			for _, char := range bytes1 {
				if slices.Contains(bytes2, char) && slices.Contains(bytes3, char) {
					total += priority(char)
					n++
					break
				}
			}
		}
	}
	fmt.Println(n)
	return total
}

func priority(b byte) int {
	val := b - 64
	if val <= abcLen {
		val = val + abcLen
	} else {
		val = val - 32
	}
	return int(val)
}

/*
func priority2(b byte) int {
	return priorityMap[string(b)]
}

var priorityMap = map[string]int{
	"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7,
	"h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14,
	"o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21,
	"v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
	"A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33,
	"H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40,
	"O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47,
	"V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52,
}
*/
