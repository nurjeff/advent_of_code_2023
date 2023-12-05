package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input, _ := io.ReadAll(file)
	sum, ratio := process(parseInput(string(input)))

	fmt.Println(sum)
	fmt.Println(int(ratio))
}

type Point struct {
	Char string
	X, Y int
}

type Number struct {
	Num        string
	Start, End Point
	Calculated bool
}

func process(data [][]string) (int, float64) {
	var numbers []Number
	var specialChars []Point
	totalSum := 0
	totalRatio := 0.0

	for x, row := range data {
		var bufferNum string
		startY := 0
		for y, char := range row {
			if unicode.IsDigit(rune(char[0])) {
				if bufferNum == "" {
					startY = y
				}
				bufferNum += char
			} else {
				if bufferNum != "" {
					numbers = append(numbers, Number{Num: bufferNum, Start: Point{X: x, Y: startY}, End: Point{X: x, Y: y - 1}})
					bufferNum = ""
				}
				if char != "." {
					specialChars = append(specialChars, Point{X: x, Y: y, Char: char})
				}
			}
		}
		if bufferNum != "" {
			numbers = append(numbers, Number{Num: bufferNum, Start: Point{X: x, Y: startY}, End: Point{X: x, Y: len(row) - 1}})
		}
	}

	for _, c := range specialChars {
		var partNums []string
		for i := range numbers {
			n := &numbers[i]
			if isAdjacent(n, &c) {
				if !n.Calculated {
					n.Calculated = true
					val, _ := strconv.Atoi(n.Num)
					totalSum += val
				}
				if c.Char == "*" {
					partNums = append(partNums, n.Num)
				}
			}
		}

		if len(partNums) == 2 {
			totalRatio += multiplyStrings(partNums[0], partNums[1])
		}
	}

	return totalSum, totalRatio
}

func isAdjacent(n *Number, c *Point) bool {
	return abs(n.End.X-c.X) <= 1 && abs(n.End.Y-c.Y) <= 1 || abs(n.Start.X-c.X) <= 1 && abs(n.Start.Y-c.Y) <= 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func multiplyStrings(num1, num2 string) float64 {
	first, _ := strconv.ParseInt(num1, 10, 64)
	second, _ := strconv.ParseInt(num2, 10, 64)
	return float64(first) * float64(second)
}

func parseInput(input string) [][]string {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	parsedData := make([][]string, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		parsedData[i] = strings.Split(line, "")
	}
	return parsedData
}
