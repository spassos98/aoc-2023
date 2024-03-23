package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var limit = 1 << 32

func getRangeValue(startRange int, rangeLength int, mapValues [][]MapRange) (int, int) {
	for _, ranges := range mapValues {
		foundValidRange := false
		smallestDiff := limit
		for _, currRange := range ranges {
			if foundValidRange {
				break
			}
			startDiff := currRange.sourceStart - startRange
			if startDiff >= 0 {
				smallestDiff = min(smallestDiff, startDiff)
			}
			if currRange.sourceStart <= startRange && startRange < currRange.sourceStart+currRange.rangeLength {
				startDiff = startRange - currRange.sourceStart
				foundValidRange = true
				if rangeLength > currRange.rangeLength-startDiff {
					rangeLength = currRange.rangeLength - startDiff
				}
				startRange = currRange.destinationStart + startDiff
			}
		}
		if !foundValidRange {
			if smallestDiff < limit {
				rangeLength = smallestDiff
			}
		}
	}

	return startRange, rangeLength
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	seedsLine := scanner.Text()
	scanner.Scan()

	firstLine := true
	for idx := 0; idx < 7; idx++ {
		var ranges []string
		keepReading := false
		for scanner.Scan() {
			if firstLine {
				firstLine = false
				continue
			}
			line := scanner.Text()
			if len(line) < 1 {
				firstLine = true
				mapsLines = append(mapsLines, ranges)
				keepReading = true
				break
			} else {
				ranges = append(ranges, line)
			}
		}
		if !keepReading {
			mapsLines = append(mapsLines, ranges)
		}
	}

	var mapValues [][]MapRange
	for _, val := range mapsLines {
		m := processMapLines(val)
		mapValues = append(mapValues, m)
	}

	var seedValues []int
	seedStrings := strings.Fields(strings.Split(seedsLine, ":")[1])
	for _, seed := range seedStrings {
		val, _ := strconv.Atoi(seed)
		seedValues = append(seedValues, val)
	}

	minLocation := 1 << 32

	for idx := 0; idx < len(seedValues); idx += 2 {
		startRange := seedValues[idx]
		rangeLength := seedValues[idx+1]
		for {
			locationValue, lengthValue := getRangeValue(startRange, rangeLength, mapValues)
			minLocation = min(locationValue, minLocation)
			if lengthValue >= rangeLength {
				break
			} else {
				startRange = startRange + lengthValue
				rangeLength = rangeLength - lengthValue
			}
		}
	}
	fmt.Printf("%d\n", minLocation)
}
