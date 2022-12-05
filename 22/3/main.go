package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

const abcLen = 26

func main() {
	readFile, err := os.Open("22/3/input")
	if err != nil {
		panic(err.Error())
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// fmt.Printf("Soution 1: %d\n", solve1(fileScanner))
	fmt.Printf("Soution 2: %d\n", solve2(fileScanner))
}

func solve1(scanner *bufio.Scanner) int {
	var total int
	for scanner.Scan() {
		bytes := scanner.Bytes()
		if len(bytes) != 0 {
			half := len(bytes) / 2
			scanner.Bytes()
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
					break
				}
			}
		}
	}
	return total
}

func priority(b byte) int { // TODO: FIX THIS! scoreMap works, this does not.
	val := b - 64
	if val <= abcLen {
		val = val + abcLen
	} else {
		val = val - 32
	}
	return int(val)
}

/*
var scoreMap = map[string]int{
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
