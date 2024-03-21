package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MapRange struct {
	sourceStart      int
	destinationStart int
	rangeLength      int
}

var mapsLines [][]string

func processRange(r string) MapRange {
	values := strings.Fields(r)
	destinationStart, _ := strconv.Atoi(values[0])
	sourceStart, _ := strconv.Atoi(values[1])
	rangeLength, _ := strconv.Atoi(values[2])

	return MapRange{sourceStart, destinationStart, rangeLength}
}

func processMapLines(mapLines []string) []MapRange {
	var arr []MapRange
	for _, mapLine := range mapLines {
		arr = append(arr, processRange(mapLine))
	}
	return arr
}

func getMapping(v int, m []MapRange) int {
	fmt.Printf("Mapping %d\n", v)
	for _, r := range m {
		if r.sourceStart <= v && v <= r.sourceStart+r.rangeLength {
			diffVal := v - r.sourceStart
			return diffVal + r.destinationStart
		}
	}
	return v
}

func getSeedLocation(seed int, mapValues [][]MapRange) int {
	currVal := seed
	for _, m := range mapValues {
		newVal := getMapping(currVal, m)
		currVal = newVal
	}
	return currVal
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	seedsLine := scanner.Text()
	fmt.Println(seedsLine)
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

	minLocation := 1<<32 - 1
	for _, seed := range seedValues {
		minLocation = min(minLocation, getSeedLocation(seed, mapValues))
	}

	fmt.Printf("%d\n", minLocation)
}

func main() {
	part1()
}
