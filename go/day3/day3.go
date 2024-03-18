package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c rune) bool {
	return !isDigit(c) && c != '.'
}
func isPartNumber(engine []string, linePos int, startPos int, endPos int) bool {
	xStartPos := max(0, linePos-1)
	yStartPos := max(0, startPos-1)
	xEndPos := min(linePos+1, len(engine)-1)
	yEndPos := min(endPos+1, len(engine[linePos])-1)
	hasSymbol := false

SearchSymbol:
	for i := xStartPos; i <= xEndPos; i++ {
		for j := yStartPos; j <= yEndPos; j++ {
			if isSymbol(rune(engine[i][j])) {
				hasSymbol = true
				break SearchSymbol
			}
		}
	}
	return hasSymbol
}

func checkLine(engine []string, engineLinePos int) int {
	readingNumber := false
	engineLine := engine[engineLinePos]
	currNumber := ""
	startPos := -1
	endPos := -1
	sum := 0
	for idx, val := range engineLine {
		if isDigit(val) {
			if !readingNumber {
				readingNumber = true
				startPos = idx
			}
			currNumber += string(val)
		} else {
			if readingNumber {
				endPos = idx - 1
				if isPartNumber(engine, engineLinePos, startPos, endPos) {
					numValue, _ := strconv.Atoi(currNumber)
					sum += numValue
				}
			}
			readingNumber = false
			currNumber = ""
			startPos = -1
			endPos = -1
		}
	}
	if readingNumber {
		endPos = len(engine[engineLinePos]) - 1
		if isPartNumber(engine, engineLinePos, startPos, endPos) {
			numValue, _ := strconv.Atoi(currNumber)
			sum += numValue
		}
	}
	return sum
}

func main() {
	part := os.Args[1]
	if part == "a" {
		scanner := bufio.NewScanner(os.Stdin)

		engine := make([]string, 0, 100)
		for scanner.Scan() {
			line := scanner.Text()
			engine = append(engine, line)
		}
		sum := 0
		for idx := range engine {
			sum += checkLine(engine, idx)
		}
		fmt.Printf("%d\n", sum)
	} else {
		fmt.Printf("Running part 2\n")
		part2()
	}
}
