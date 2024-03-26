package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type HandValue int

const (
	None      HandValue = iota
	HighCard  HandValue = iota
	OnePair   HandValue = iota
	TwoPair   HandValue = iota
	ThreeKind HandValue = iota
	FullHouse HandValue = iota
	FourKind  HandValue = iota
	FiveKind  HandValue = iota
)

type Play struct {
	handValue HandValue
	hand      string
}

type Bid struct {
	play Play
	bid  int
}

func compareCards(a rune, b rune) int {
	if a == b {
		return 0
	}

	if a == 'J' {
		return -1
	}

	if b == 'J' {
		return 1
	}

	if '2' <= a && a <= '9' {
		if '2' <= b && b <= '9' {
			// a and b are numbers
			if a > b {
				return 1
			}
			return -1
		} else {
			// a is number and b is letter
			return -1
		}
	} else if '2' <= b && b <= '9' {
		// a is letter and b is number
		return 1
	}
	// a and b are letters
	if a == 'A' {
		return 1
	}
	if b == 'A' {
		return -1
	}

	if a == 'K' {
		return 1
	}
	if b == 'K' {
		return -1
	}

	if a == 'Q' {
		return 1
	}
	if b == 'Q' {
		return -1
	}

	if a == 'T' {
		return 1
	}
	if b == 'T' {
		return -1
	}
	return 1
}

func comparePlays(a Play, b Play) int {
	if a.handValue == b.handValue {
		for idx := 0; idx < len(a.hand); idx++ {
			compVal := compareCards(rune(a.hand[idx]), rune(b.hand[idx]))
			if compVal != 0 {
				return compVal
			}
		}
		return 0
	}
	if a.handValue > b.handValue {
		return 1
	} else {
		return -1
	}
}

func compareBids(a Bid, b Bid) int {
	return comparePlays(a.play, b.play)
}

func getPlay(handLine string) Play {
	cardCount := make(map[rune]int)
	for _, card := range handLine {
		if _, ok := cardCount[card]; ok {
			cardCount[card] += 1
		} else {
			cardCount[card] = 1
		}
	}

	nJokers := 0
	if val, ok := cardCount['J']; ok {
		nJokers = val
	}

	var repetitionArr []int

	for _, repetition := range cardCount {
		repetitionArr = append(repetitionArr, repetition)
	}

	slices.Sort(repetitionArr)
	switch lArr := len(repetitionArr); lArr {
	case 1:
		return Play{FiveKind, handLine}
	case 2:
		if repetitionArr[1] == 4 {
			if nJokers > 0 {
				return Play{FiveKind, handLine}
			}
			return Play{FourKind, handLine}
		}
		if nJokers > 0 {
			return Play{FiveKind, handLine}
		}
		return Play{FullHouse, handLine}
	case 3:
		if repetitionArr[2] == 3 {
			if nJokers > 0 {
				return Play{FourKind, handLine}
			}
			return Play{ThreeKind, handLine}
		}
		if nJokers == 1 {
			return Play{FullHouse, handLine}
		} else if nJokers == 2 {
			return Play{FourKind, handLine}
		}
		return Play{TwoPair, handLine}
	case 4:
		if nJokers > 0 {
			return Play{ThreeKind, handLine}
		}
		return Play{OnePair, handLine}
	default:
		if nJokers > 0 {
			return Play{OnePair, handLine}
		}
		return Play{HighCard, handLine}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var bids []Bid
	for scanner.Scan() {
		bidLine := scanner.Text()
		splits := strings.Fields(bidLine)
		hand := splits[0]
		bid, _ := strconv.Atoi(splits[1])
		play := getPlay(hand)
		bids = append(bids, Bid{play, bid})

	}
	slices.SortStableFunc(bids, compareBids)
	fmt.Printf("%v\n", bids)
	sum := 0
	for idx, bid := range bids {
		sum += bid.bid * (idx + 1)
	}
	fmt.Printf("%d\n", sum)
}
