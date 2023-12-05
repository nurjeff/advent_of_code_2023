package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	powerSum := 0
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		maxPossibleMap := process(scanner.Text())

		// Part 1 solution
		if maxPossibleMap["red"] <= 12 &&
			maxPossibleMap["blue"] <= 14 &&
			maxPossibleMap["green"] <= 13 {
			sum += i + 1
		}

		// Part 2 solution
		powerSum += int(maxPossibleMap["red"] * maxPossibleMap["blue"] * maxPossibleMap["green"])
	}

	fmt.Println(sum)
	fmt.Println(powerSum)
}

// With this approach, both part solutions can be calculated in one go
func process(input string) map[string]int64 {
	parts := strings.Split(input, ": ")

	iterations := strings.Split(parts[1], ";")
	maxMap := make(map[string]int64)

	for _, i := range iterations {
		items := strings.Split(strings.TrimSpace(i), ",")
		for _, item := range items {
			valCol := strings.Fields(strings.TrimSpace(item))
			if len(valCol) < 2 {
				continue
			}
			v, _ := strconv.ParseInt(valCol[0], 10, 64)
			if v > maxMap[valCol[1]] {
				maxMap[valCol[1]] = v
			}
		}
	}

	return maxMap
}
