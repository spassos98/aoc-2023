package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func processCard(cardLine string) int {
	firstSplit := strings.Split(cardLine, ":")
	gameInfo := firstSplit[1]
	numbersLine := strings.Split(gameInfo, "|")
	winningNumbers := strings.Fields(numbersLine[0])
	playNumbers := strings.Fields(numbersLine[1])
	matchCount := 0
	for _, number := range playNumbers {
		if slices.Contains(winningNumbers, number) {
			matchCount += 1
		}
	}

	if matchCount > 0 {
		return 1 << (matchCount - 1)
	}
	return 0
}

func main() {
	part := os.Args[1]
	if part == "2" {
		part2()
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		sum := 0
		for scanner.Scan() {
			sum += processCard(scanner.Text())
		}
		fmt.Printf("%d\n", sum)
	}
}
