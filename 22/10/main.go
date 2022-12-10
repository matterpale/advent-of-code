package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	input     = "22/10/input"
	crtWidth  = 40
	crtHeight = 6
)

func inputScanner() *bufio.Scanner {
	readFile, _ := os.Open(input)
	info, _ := readFile.Stat()
	scanner := bufio.NewScanner(readFile)
	scanner.Buffer(make([]byte, 0, info.Size()), 0)
	return scanner
}

func main() {
	fmt.Printf("Solution 1: %d\nSolution 2:\n\n",
		solve1(inputScanner()))
	solve2(inputScanner())
}

func solve1(scanner *bufio.Scanner) int {
	cycle, signalSum, x := 0, 0, 1
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			split := strings.Split(txt, " ")
			switch split[0] {
			case "noop":
				step(&cycle, &signalSum, x)
			case "addx":
				step(&cycle, &signalSum, x)
				step(&cycle, &signalSum, x)
				add, _ := strconv.Atoi(split[1])
				x = x + add
			}
		}
	}
	return signalSum
}

func step(cycle, sum *int, x int) {
	*cycle++
	if (*cycle-20)%40 == 0 {
		*sum = *sum + (*cycle)*x
	}
}

func solve2(scanner *bufio.Scanner) {
	channel := make(chan bool, crtHeight*crtWidth)
	defer close(channel)
	device := &Device{
		cpu: &CPU{
			scanner: scanner,
			outChan: channel,
			x:       1,
		},
		crt: &CRT{
			inChan: channel,
		},
	}
	device.Boot()
}
