package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type ScracthCard struct {
	cardNumber     int
	winningNumbers []string
	playNumbers    []string
}

var cardsGenerated []int
var cards []ScracthCard

func getNumberOfGeneratedCards(cardNumber int) int {
	if cardsGenerated[cardNumber] != 0 {
		return cardsGenerated[cardNumber]
	}
	generatedCards := getGeneratedCards(cards[cardNumber-1], len(cards))
	sum := 1
	for _, val := range generatedCards {
		sum += getNumberOfGeneratedCards(val)
	}
	return sum
}

func getGeneratedCards(card ScracthCard, maxCard int) []int {
	matchingNumbers := 0
	for _, number := range card.playNumbers {
		if slices.Contains(card.winningNumbers, number) {
			matchingNumbers += 1
		}
	}
	var cardsGenerated []int
	for x := 1; x <= matchingNumbers; x++ {
		generatedCard := x + card.cardNumber
		if generatedCard <= maxCard {
			cardsGenerated = append(cardsGenerated, generatedCard)
		} else {
			break
		}
	}

	return cardsGenerated
}

func processLine(cardLine string) ScracthCard {
	firstSplit := strings.Split(cardLine, ":")
	cardNumber, _ := strconv.Atoi(strings.Fields(firstSplit[0])[1])
	lineNumbers := strings.Split(firstSplit[1], "|")
	winningNumbers := strings.Fields(lineNumbers[0])
	playNumbers := strings.Fields(lineNumbers[1])
	return ScracthCard{cardNumber, winningNumbers, playNumbers}
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, processLine(line))
	}

	for x := 0; x <= len(cards)+1; x++ {
		cardsGenerated = append(cardsGenerated, 0)
	}

	sum := 0
	for x := len(cards); x >= 1; x-- {
		sum += getNumberOfGeneratedCards(x)
	}

	fmt.Printf("%d\n", sum)
}
