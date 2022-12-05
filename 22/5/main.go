package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = "22/5/input"

func main() {
	readFile1, _ := os.Open(input)
	readFile2, _ := os.Open(input)

	fileScanner1 := bufio.NewScanner(readFile1)
	fileScanner1.Split(bufio.ScanLines)
	fmt.Printf("Solution 1: %s\n", solve1(fileScanner1))

	fileScanner2 := bufio.NewScanner(readFile2)
	fileScanner2.Split(bufio.ScanLines)
	fmt.Printf("Solution 2: %s\n", solve2(fileScanner2))
}

func solve1(scanner *bufio.Scanner) string {
	crg := prep(scanner)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			split := strings.Split(txt, " ")
			load, _ := strconv.Atoi(split[1])
			src, _ := strconv.Atoi(split[3])
			dst, _ := strconv.Atoi(split[5])
			lifted := lift1(&crg[src-1], load)
			crg[dst-1] = append(crg[dst-1], lifted...)
		}
	}
	return topWord(crg)
}

func solve2(scanner *bufio.Scanner) string {
	crg := prep(scanner)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			split := strings.Split(txt, " ")
			load, _ := strconv.Atoi(split[1])
			src, _ := strconv.Atoi(split[3])
			dst, _ := strconv.Atoi(split[5])
			lifted := lift2(&crg[src-1], load)
			crg[dst-1] = append(crg[dst-1], lifted...)
		}
	}
	return topWord(crg)
}

func prep(scanner *bufio.Scanner) (crg [][]byte) {
	for scanner.Scan() {
		bytes := scanner.Bytes()
		if len(bytes) == 0 {
			break
		}
		slot, ix := 0, 1
		if bytes[ix] == 49 {
			continue
		}
		for ; ix < len(bytes); slot, ix = slot+1, ix+4 {
			if slot >= len(crg) {
				crg = append(crg, []byte{})
			}
			if b := bytes[ix]; b != 32 {
				crg[slot] = append(crg[slot], bytes[ix])
			}
		}
	}

	for _, s := range crg {
		reverse(&s)
	}
	return
}

func topWord(c [][]byte) string {
	var word []byte
	for _, s := range c {
		word = append(word, s[len(s)-1])
	}
	return string(word)
}

func lift1(slot *[]byte, load int) []byte {
	lifted := (*slot)[len(*slot)-load:]
	*slot = (*slot)[:len(*slot)-load]
	reverse(&lifted)
	return lifted
}

func lift2(slot *[]byte, load int) []byte {
	lifted := (*slot)[len(*slot)-load:]
	*slot = (*slot)[:len(*slot)-load]
	return lifted
}

func reverse(s *[]byte) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}