package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type WP struct {
	L   *WP
	R   *WP
	Val string
}

var WPMap map[string]*WP

var moves []string

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input, _ := io.ReadAll(file)
	lines := strings.Split(strings.ReplaceAll(string(input), "\r\n", "\n"), "\n")
	moves = strings.Split(lines[0], "")

	WPMap = make(map[string]*WP)

	currentNodes := []*WP{}

	// first fill base nodes
	for _, e := range lines[2:] {
		id := e[:3]
		_, ok := WPMap[id]
		if ok {
			continue
		}

		WPMap[id] = &WP{Val: id}
	}

	// build linked graph
	for _, e := range lines[2:] {
		id := e[:3]
		lwp := e[7:10]
		rwp := e[12:15]
		n := WPMap[id]
		if id[2:] == "A" {
			currentNodes = append(currentNodes, n)
		}
		n.L = WPMap[lwp]
		n.R = WPMap[rwp]
	}

	// Dirty LCM Solution
	// Only works because it loops around to the starting node, which is not specified. meh.
	maxSteps := []int{}
	for _, e := range currentNodes {
		currentNode := e
		currentStep := 0
		totalSteps := 0

		for {
			if moves[currentStep] == "L" {
				currentNode = currentNode.L
			} else {
				currentNode = currentNode.R
			}
			totalSteps++
			currentStep++
			if currentStep > len(moves)-1 {
				currentStep = 0
			}
			if currentNode.Val[2:] == "Z" {
				maxSteps = append(maxSteps, totalSteps)
				break
			}
		}
	}

	fmt.Println(lcmSlice(maxSteps))

	// Brute force solution
	// might take the latter half of your life though
	/*
		amount := 0
		i := 0
		allOnFin := true
			for {
				for x := 0; x < len(currentNodes); x++ {
					if (moves[i]) == "L" {
						currentNodes[x] = currentNodes[x].L
					} else {
						currentNodes[x] = currentNodes[x].R
					}
					if currentNodes[x].Val[2:] != "Z" {
						allOnFin = false
					}
				}
				i++
				if i > len(moves)-1 {
					i = 0
				}
				amount++
				if allOnFin {
					fmt.Println(amount)
					return
				}

				allOnFin = true
			}
	*/
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmSlice(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	result := nums[0]
	for _, num := range nums[1:] {
		result = lcm(result, num)
	}
	return result
}
