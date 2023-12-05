package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Solution without regex
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(processInput(input, processPart1))
	fmt.Println(processInput(input, processPart2))
}

func processInput(input []string, processor func(string) uint32) uint32 {
	var sum uint32
	for _, line := range input {
		sum += processor(line)
	}
	return sum
}

func processPart1(line string) uint32 {
	first, _ := extractDigit(line, true)
	last, _ := extractDigit(line, false)
	return 10*first + last
}

func processPart2(line string) uint32 {
	nums := map[string]uint32{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	first := findNumber(line, nums, true)
	last := findNumber(line, nums, false)
	return first*10 + last
}

func extractDigit(s string, first bool) (uint32, bool) {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length; i++ {
		index := i
		if !first {
			index = length - 1 - i
		}
		if runes[index] >= '0' && runes[index] <= '9' {
			return uint32(runes[index] - '0'), true
		}
	}
	return 0, false
}

func findNumber(s string, nums map[string]uint32, first bool) uint32 {
	for len(s) > 0 {
		for k, v := range nums {
			if (first && strings.HasPrefix(s, k)) || (!first && strings.HasSuffix(s, k)) {
				return v
			}
		}
		if first {
			s = s[1:]
		} else {
			s = s[:len(s)-1]
		}
	}
	return 0
}
