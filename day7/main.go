package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var scoreMap = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 1,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type Hand struct {
	Hand         string
	OriginalHand string
	Bid          int
	Type         int
}

var hands []Hand

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input, _ := io.ReadAll(file)
	lines := strings.Split(strings.ReplaceAll(string(input), "\r\n", "\n"), "\n")

	t := time.Now()
	for _, r := range lines {
		l := strings.Split(r, " ")
		bidParsed, _ := strconv.ParseInt(l[1], 10, 64)
		bid := int(bidParsed)

		hand := Hand{
			Hand:         l[0],
			OriginalHand: l[0],
			Bid:          bid,
		}

		countMap := make(map[rune]int)
		maxRune := rune(hand.Hand[0])
		maxSoFar := 0
		for _, e := range hand.Hand {
			if e != 'J' {
				countMap[e]++
				if countMap[e] > maxSoFar {
					maxRune = e
					maxSoFar = countMap[e]
				}
			}
		}

		hand.Hand = strings.ReplaceAll(hand.Hand, "J", string(maxRune))
		hand.GetLevel()
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		return WinsVsOther(&hands[i], &hands[j])
	})

	totalWin := 0
	for i, r := range hands {
		totalWin += (len(hands) - i) * r.Bid
	}
	fmt.Println(totalWin, time.Since(t))
}

const (
	_ = iota
	FiveOfAKind
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func WinsVsOther(thisHand *Hand, otherHand *Hand) bool {
	if thisHand.Type > otherHand.Type {
		return false
	}
	if thisHand.Type < otherHand.Type {
		return true
	}
	ind := 0
	for {
		lhv := thisHand.OriginalHand[ind]
		rhv := otherHand.OriginalHand[ind]
		if lhv == rhv {
			ind++
			continue
		}

		lhs := scoreMap[string(thisHand.OriginalHand[ind])]
		rhs := scoreMap[string(otherHand.OriginalHand[ind])]
		return lhs > rhs
	}
}

func (hand *Hand) GetLevel() {
	s := []rune(hand.Hand)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	c1, c2, c3, c4, c5 := s[0], s[1], s[2], s[3], s[4]
	switch {
	case c1 == c5:
		hand.Type = FiveOfAKind
	case c1 == c4 || c2 == c5:
		hand.Type = FourOfAKind
	case (c1 == c3 && c4 == c5) || (c1 == c2 && c3 == c5):
		hand.Type = FullHouse
	case c1 == c3 || c2 == c4 || c3 == c5:
		hand.Type = ThreeOfAKind
	case (c1 == c2 && c3 == c4) || (c2 == c3 && c4 == c5) || (c1 == c2 && c4 == c5):
		hand.Type = TwoPair
	case c1 == c2 || c2 == c3 || c3 == c4 || c4 == c5:
		hand.Type = OnePair
	default:
		hand.Type = HighCard
	}
}
