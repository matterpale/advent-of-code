package y22d05

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/matterpale/advent-of-code/domain"
)

const dir = "solutions/22/05/"

type Solver struct {
	scanner func() *bufio.Scanner
}

func (s *Solver) Init() domain.Solver {
	s.scanner = domain.InputScannerGen(dir)
	return s
}

func (s *Solver) Solve(args ...string) {
	if len(args) != 1 {
		log.Panicf("invalid number of arguments: %d", len(args))
	}
	modern, err := strconv.ParseBool(args[0])
	if err != nil {
		log.Panicf("invalid argument: '%s', expected bool", args[0])
	}
	fmt.Printf("Solution for modern == %t:\n", modern)
	s.solve(modern)
}

func (s *Solver) Solve1() {
	fmt.Printf("Solution 01:\n%s\n", s.solve(false))
}

func (s *Solver) Solve2() {
	fmt.Printf("Solution 02:\n%s\n", s.solve(true))
}

func (s *Solver) solve(modern bool) string {
	scanner := s.scanner()
	cargo := setupCargo(scanner)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			split := strings.Split(txt, " ")
			load, _ := strconv.Atoi(split[1])
			src, _ := strconv.Atoi(split[3])
			dst, _ := strconv.Atoi(split[5])
			lifted := lift(&cargo[src-1], load, modern)
			cargo[dst-1] = append(cargo[dst-1], lifted...)
		}
	}
	return topWord(cargo)
}

func setupCargo(scanner *bufio.Scanner) (crg [][]byte) {
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

func lift(slot *[]byte, load int, modern bool) []byte {
	lifted := (*slot)[len(*slot)-load:]
	*slot = (*slot)[:len(*slot)-load]
	if !modern {
		reverse(&lifted)
	}
	return lifted
}

func reverse(s *[]byte) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func topWord(c [][]byte) string {
	var word []byte
	for _, s := range c {
		word = append(word, s[len(s)-1])
	}
	return string(word)
}
