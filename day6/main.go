package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	parseInput()
}

type Set struct {
	Time     int
	Distance int
}

var sets []Set

func parseInput() {
	file, err := os.Open("input_2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input, _ := io.ReadAll(file)
	inputString := string(input)

	lines := strings.Split(strings.ReplaceAll(inputString, "\r\n", "\n"), "\n")

	times := strings.Split(lines[0], " ")
	distances := strings.Split(lines[1], " ")

	for i, e := range times {
		t, _ := strconv.ParseInt(e, 10, 64)
		d, _ := strconv.ParseInt(distances[i], 10, 64)
		sets = append(sets, Set{Time: int(t), Distance: int(d)})
	}

	t := time.Now()
	answer := 1
	for _, e := range sets {
		winsamounts := 0
		for i := 0; i <= e.Time; i++ {
			holdingTime := e.Time - i
			remainingTime := e.Time - holdingTime
			distance := holdingTime * remainingTime
			if distance > e.Distance {
				winsamounts++
			}
		}
		answer *= winsamounts

	}
	fmt.Println(answer, time.Since(t))
}
