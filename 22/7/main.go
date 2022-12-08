package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	input    = "22/7/input"
	total    = 70_000_000
	required = 30_000_000
)

func main() {
	readFile1, _ := os.Open(input)
	readFile2, _ := os.Open(input)

	fileScanner1 := bufio.NewScanner(readFile1)
	fileScanner1.Split(bufio.ScanLines)
	fmt.Printf("Solution 1: %d\n", solve(fileScanner1, false))

	fileScanner2 := bufio.NewScanner(readFile2)
	fileScanner2.Split(bufio.ScanLines)
	fmt.Printf("Solution 2: %d\n", solve(fileScanner2, true))
}

func solve(scanner *bufio.Scanner, task2 bool) int {
	root := prepFileSystem(scanner)
	used := root.getSetSize()
	if !task2 {
		return root.sumSmallDirSizesWithDuplicities()
	}

	free := total - used
	need := required - free
	var get int
	if need > 0 {
		get = root.findSmallestGarbageDirSize(used, need)
	}
	return get
}

type file struct {
	size int
}

func newFile(size int) *file {
	return &file{size}
}

type directory struct {
	super   *directory
	subDirs map[string]*directory
	files   map[string]*file
	size    int
}

func newDir(super *directory) *directory {
	return &directory{
		super:   super,
		subDirs: make(map[string]*directory),
		files:   make(map[string]*file),
	}
}

func (d *directory) getSetSize() (size int) {
	for _, f := range d.files {
		size = size + f.size
	}
	for _, s := range d.subDirs {
		size = size + s.getSetSize()
	}
	d.size = size
	return
}

func (d *directory) sumSmallDirSizesWithDuplicities() (sum int) {
	for _, s := range d.subDirs {
		if size := s.size; size <= 100_000 {
			sum = sum + size
		}
		sum = sum + s.sumSmallDirSizesWithDuplicities()
	}
	return
}

func (d *directory) findSmallestGarbageDirSize(candidate, need int) int {
	for _, s := range d.subDirs {
		if size := s.size; size >= need {
			if size < candidate {
				candidate = size
			}
			candidate = s.findSmallestGarbageDirSize(candidate, need)
		}
	}
	return candidate
}

func prepFileSystem(scanner *bufio.Scanner) *directory {
	curDir := newDir(nil)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			split := strings.Split(txt, " ")
			switch split[0] {
			case "$":
				if split[1] == "cd" {
					switch split[2] {
					case "/":
					case "..":
						curDir = curDir.super
					default:
						curDir = curDir.subDirs[split[2]]
					}
				}
			case "dir":
				curDir.subDirs[split[1]] = newDir(curDir)
			default:
				size, _ := strconv.Atoi(split[0])
				curDir.files[split[1]] = newFile(size)
			}
		}
	}
	for curDir.super != nil {
		curDir = curDir.super
	}
	return curDir
}
