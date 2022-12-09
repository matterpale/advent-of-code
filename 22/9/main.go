package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"adventofcode/22/9/ropes"
)

const (
	input         = "22/9/input"
	longRopeKnots = 10
)

func main() {
	fmt.Printf("Solution 1: %d\nSolution 2: %d",
		solve(inputScanner(), false),
		solve(inputScanner(), true))
}

func solve(scanner *bufio.Scanner, long bool) int {
	tailVisited := make(map[ropes.Position]bool)

	if !long {
		rope := ropes.NewRope()
		for scanner.Scan() {
			txt := scanner.Text()
			if txt != "" {
				split := strings.Split(txt, " ")
				direction := split[0]
				steps, _ := strconv.Atoi(split[1])
				for i := 0; i < steps; i++ {
					rope.MoveHead(ropes.Direction(direction))
					tailVisited[rope.TailPosition()] = true
				}
			}
		}
		return len(tailVisited)
	}

	longRope := ropes.NewLongRope(longRopeKnots)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			split := strings.Split(txt, " ")
			d := split[0]
			steps, _ := strconv.Atoi(split[1])
			for i := 0; i < steps; i++ {
				longRope.MoveHead(ropes.Direction(d))
				tailVisited[longRope.TailPosition()] = true
			}
		}
	}
	return len(tailVisited)
}

func inputScanner() *bufio.Scanner {
	readFile, _ := os.Open(input)
	info, _ := readFile.Stat()
	scanner := bufio.NewScanner(readFile)
	scanner.Buffer(make([]byte, 0, info.Size()), 0)
	return scanner
}
