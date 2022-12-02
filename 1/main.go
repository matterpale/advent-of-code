package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("1/input")
	if err != nil {
		panic(err.Error())
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	elves := []int{0}
	var i int
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if txt == "" {
			i++
			elves = append(elves, 0)
			continue
		}
		n, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			panic("atoi")
		}
		elves[i] += n
	}

	var max1, max2, max3 int
	for _, v := range elves {
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}
	fmt.Println(max1)
	fmt.Println(max2)
	fmt.Println(max3)
	fmt.Println(max1 + max2 + max3)
}
