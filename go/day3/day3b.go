package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var gearMap = make(map[string][]int)

func updateMap(i int, j int, value int) {
	key := fmt.Sprintf("%d-%d", i, j)
	if val, ok := gearMap[key]; ok {
		gearMap[key] = append(val, value)
	} else {
		gearMap[key] = make([]int, 0)
		gearMap[key] = append(gearMap[key], value)
	}
	fmt.Printf("\t\t\t\t")
	fmt.Println(gearMap[key])
}

func processPartNumber(engine []string, linePos int, startPos int, endPos int, numValue int) {
	xStartPos := max(0, linePos-1)
	yStartPos := max(0, startPos-1)
	xEndPos := min(linePos+1, len(engine)-1)
	yEndPos := min(endPos+1, len(engine[linePos])-1)

	for i := xStartPos; i <= xEndPos; i++ {
		for j := yStartPos; j <= yEndPos; j++ {
			if isSymbol(rune(engine[i][j])) {
				if engine[i][j] == '*' {
					updateMap(i, j, numValue)
				}
			}
		}
	}
}

func processLinePart2(engine []string, engineLinePos int) {
	readingNumber := false
	engineLine := engine[engineLinePos]
	currNumber := ""
	startPos := -1
	endPos := -1
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
				numValue, _ := strconv.Atoi(currNumber)
				processPartNumber(engine, engineLinePos, startPos, endPos, numValue)
			}
			readingNumber = false
			currNumber = ""
			startPos = -1
			endPos = -1
		}
	}
	if readingNumber {
		endPos = len(engine[engineLinePos]) - 1
		numValue, _ := strconv.Atoi(currNumber)
		processPartNumber(engine, engineLinePos, startPos, endPos, numValue)
	}
}

func part2() {

	scanner := bufio.NewScanner(os.Stdin)

	engine := make([]string, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		engine = append(engine, line)
	}
	for idx := range engine {
		processLinePart2(engine, idx)
	}
	sum := 0
	for _, b := range gearMap {
		if len(b) == 2 {
			sum += b[0] * b[1]
		}

	}
	fmt.Printf("%d\n", sum)
}
