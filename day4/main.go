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
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input, _ := io.ReadAll(file)
	cards := parseInput(string(input))
	totalPoints := 0
	totalCards := 0

	t := time.Now()

	for i := 0; i < len(cards); i++ {
		totalPoints += cards[i].Points
		for x := 1; x < cards[i].Matches+1; x++ {
			cards[i+x].Amount += cards[i].Amount
		}
		totalCards += cards[i].Amount
	}
	fmt.Println(time.Since(t))

	fmt.Println(totalCards)
	fmt.Println(totalPoints)
}

func (c *Card) calculatePoints() {
	c.Matches = c.calculateMatches()
	if c.Matches == 0 {
		return
	}
	c.Points = 1 << (c.Matches - 1)
}

func (c *Card) calculateMatches() int {
	matches := 0
	for _, n := range c.ScratchedNums {
		if c.CardNumsMap[n] {
			matches++
		}
	}
	return matches
}

type Card struct {
	ID            int
	CardNumsMap   map[int]bool
	CardNums      []int
	ScratchedNums []int
	Matches       int
	Points        int
	Amount        int
}

func parseInput(input string) []Card {
	cards := []Card{}
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	for i, l := range lines {
		cleaned := strings.Split(l, ":")[1]
		split := strings.Split(cleaned, "|")

		nums1 := strings.Split(split[0], " ")
		nums2 := strings.Split(split[1], " ")
		_ = nums2

		card := Card{CardNumsMap: make(map[int]bool), ID: i + 1, Amount: 1}

		for _, n := range nums1 {
			num := strings.TrimSpace(n)
			if len(num) > 0 {
				numParsed, _ := strconv.ParseInt(num, 10, 64)
				card.CardNums = append(card.CardNums, int(numParsed))
				card.CardNumsMap[int(numParsed)] = true
			}
		}

		for _, n := range nums2 {
			num := strings.TrimSpace(n)
			if len(num) > 0 {
				numParsed, _ := strconv.ParseInt(num, 10, 64)
				card.ScratchedNums = append(card.ScratchedNums, int(numParsed))
			}
		}

		card.calculatePoints()
		cards = append(cards, card)
	}
	return cards
}
