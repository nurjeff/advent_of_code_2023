package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input_2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input, _ := io.ReadAll(file)
	lines := strings.Split(strings.ReplaceAll(string(input), "\r\n", "\n"), "\n")
	times := strings.Split(lines[0], " ")
	distances := strings.Split(lines[1], " ")

	for i, e := range times {
		t, _ := strconv.ParseInt(e, 10, 64)
		d, _ := strconv.ParseInt(distances[i], 10, 64)
		sets = append(sets, Set{Time: int(t), Distance: int(d)})
	}

	answer := 1
	for _, e := range sets {
		winsamounts := 0
		for i := 0; i <= e.Time; i++ {
			holdingTime := e.Time - i
			distance := holdingTime * (e.Time - holdingTime)
			if distance > e.Distance {
				winsamounts++
			}
		}
		answer *= winsamounts

	}
	fmt.Println(answer)
}

type Set struct {
	Time     int
	Distance int
}

var sets []Set
