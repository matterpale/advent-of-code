package y22d07

import (
	"bufio"
	"strconv"
	"strings"
)

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

func (d *directory) fetchSize() (size int) {
	for _, f := range d.files {
		size = size + f.size
	}
	for _, s := range d.subDirs {
		size = size + s.fetchSize()
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
