package y22d10

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

type Device struct {
	cpu *CPU
	crt *CRT
}

func (d *Device) Boot() {
	var wg sync.WaitGroup
	wg.Add(2)
	go d.cpu.Run(&wg)
	go d.crt.Run(&wg)
	wg.Wait()
}

type CPU struct {
	scanner  *bufio.Scanner
	outChan  chan bool
	cycle, x int
}

func (cpu *CPU) Run(waiter *sync.WaitGroup) {
	defer waiter.Done()
	for cpu.scanner.Scan() {
		txt := cpu.scanner.Text()
		if txt != "" {
			split := strings.Split(txt, " ")
			switch split[0] {
			case "noop":
				cpu.sendPixel()
			case "addx":
				cpu.sendPixel()
				cpu.sendPixel()
				add, _ := strconv.Atoi(split[1])
				cpu.x = cpu.x + add
			}
		}
	}
}

func (cpu *CPU) sendPixel() {
	if math.Abs(float64(cpu.x-(cpu.cycle%40))) < 1.5 {
		cpu.outChan <- true
	} else {
		cpu.outChan <- false
	}
	cpu.cycle++
}

type CRT struct {
	inChan chan bool
	screen Screen
}

func (crt *CRT) Run(waiter *sync.WaitGroup) {
	defer waiter.Done()
	for i, row := range crt.screen {
		for j := range row {
			select {
			case px, ok := <-crt.inChan:
				if !ok {
					crt.screen.display()
				}
				crt.screen[i][j] = px
			}
		}
	}
	crt.screen.display()
}

type Screen [crtHeight][crtWidth]bool

func (s *Screen) display() {
	pxMap := map[bool]string{true: "#", false: "."}
	for _, r := range s {
		for _, p := range r {
			fmt.Print(pxMap[p])
		}
		fmt.Println("")
	}
}
