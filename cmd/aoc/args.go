package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func getPuzzleString(year, day string) (string, error) {
	y, err := strconv.Atoi(year)
	if err != nil {
		return "", errors.New("invalid year")
	} else if y < 15 {
		return "", fmt.Errorf("aoc dates back to only 2015 (year entered: %d)", y)
	} else if y > time.Now().Year()-2000 || !currentAocYear(y) {
		return "", fmt.Errorf("have patience (year entered: %d)", y)
	}

	d, err := strconv.Atoi(day)
	if err != nil {
		return "", errors.New("invalid year")
	} else if d < 1 || d > 25 {
		return "", fmt.Errorf("advent calendar only spans 01-25 (day entered: %d)", d)
	} else if !currentAocDay(y, d) {
		return "", fmt.Errorf("have patience (year and day entered: %d %d)", y, d)
	}

	return fmt.Sprintf("y%dd%02d", y, d), nil
}

func currentAocYear(year int) bool {
	return year == time.Now().Year()-2000 && time.Now().After(time.Date(time.Now().Year(), time.December, 1, 6, 0, 0, 0, time.Local))
}

func currentAocDay(year, day int) bool {
	return year == time.Now().Year()-2000 && time.Now().After(time.Date(time.Now().Year(), time.December, day, 6, 0, 0, 0, time.Local))
}
