package y22d11

import (
	"bufio"
	"strconv"
	"strings"
)

type monkey struct {
	items       []*item
	operation   func(*item)
	test        func(int) bool
	targets     [2]*monkey
	inspections int
}

type item struct {
	worryLevel int
}

func (m *monkey) throwAll() {
	for len(m.items) > 0 {
		m.throw()
	}
}

func (m *monkey) throw() {
	it := m.items[0]
	m.inspect(it)
	if m.test(it.worryLevel) {
		m.targets[0].items = append(m.targets[0].items, it)
	} else {
		m.targets[1].items = append(m.targets[1].items, it)
	}
	m.items = m.items[1:]
}

func (m *monkey) inspect(it *item) {
	m.operation(it)
	m.inspections++
}

func initMonkeys(scanner *bufio.Scanner, relief bool) []*monkey {
	var monkeys []*monkey
	var currentMonkey *monkey
	var targets [][2]int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		trimmed := strings.TrimLeft(txt, " ")
		split := strings.Split(trimmed, " ")
		switch split[0] {
		case "Monkey":
			currentMonkey = &monkey{}
			monkeys = append(monkeys, currentMonkey)
		case "Starting":
			s := strings.Split(txt, ": ")
			items := strings.Split(s[1], ", ")
			for _, worryLevel := range items {
				wl, _ := strconv.Atoi(worryLevel)
				currentMonkey.items = append(currentMonkey.items, &item{worryLevel: wl})
			}
		case "Operation:":
			op := split[4]
			var self bool
			x, err := strconv.Atoi(split[5])
			if err != nil {
				if split[5] == "old" {
					self = true
				}
			}
			switch op {
			case "+":
				if self {
					currentMonkey.operation = func(i *item) {
						i.worryLevel = relax(2*i.worryLevel, relief)
					}
				} else {
					currentMonkey.operation = func(i *item) {
						i.worryLevel = relax(i.worryLevel+x, relief)
					}
				}
			case "*":
				if self {
					currentMonkey.operation = func(i *item) {
						i.worryLevel = relax(i.worryLevel*i.worryLevel, relief)
					}
				} else {
					currentMonkey.operation = func(i *item) {
						i.worryLevel = relax(i.worryLevel*x, relief)
					}
				}
			}
		case "Test:":
			divisible, _ := strconv.Atoi(split[3])
			currentMonkey.test = func(i int) bool {
				return i%divisible == 0
			}
		case "If":
			target, _ := strconv.Atoi(split[5])
			switch split[1] {
			case "true:":
				targets = append(targets, [2]int{target})
			case "false:":
				targets[len(targets)-1][1] = target
			}
		}
	}

	for i, m := range monkeys {
		m.targets[0], m.targets[1] = monkeys[targets[i][0]], monkeys[targets[i][1]]
	}
	return monkeys
}

func relax(worryLevel int, relief bool) int {
	if relief {
		return worryLevel / 3
	}
	return worryLevel
}
