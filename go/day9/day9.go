package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func transformArray(sequenceLine []string) []int {
	var intArray []int
	for _, val := range sequenceLine {
		transformVal, _ := strconv.Atoi(val)
		intArray = append(intArray, transformVal)
	}
	return intArray
}

func getDifferenceSequence(arr []int) []int {
	var differenceSequence []int
	for idx := 0; idx < len(arr)-1; idx++ {
		diff := arr[idx+1] - arr[idx]
		differenceSequence = append(differenceSequence, diff)
	}
	return differenceSequence
}

func sequenceIsZero(arr []int) bool {
	for _, val := range arr {
		if val != 0 {
			return false
		}
	}
	return true
}

func extrapolateSequence(arr []int) int {
	diffSeq := getDifferenceSequence(arr)
	if sequenceIsZero(diffSeq) {
		return arr[len(arr)-1]
	}
	extrapolateDiff := extrapolateSequence(diffSeq)
	return arr[len(arr)-1] + extrapolateDiff
}

func extrapolateSequencePart2(arr []int) int {
	diffSeq := getDifferenceSequence(arr)
	if sequenceIsZero(diffSeq) {
		return arr[len(arr)-1]
	}
	extrapolateDiff := extrapolateSequencePart2(diffSeq)
	return arr[0] - extrapolateDiff
}

func main() {
	part := os.Args[1]
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sequenceLine := strings.Fields(line)
		arr := transformArray(sequenceLine)
		var val int
		if part == "2" {
			val = extrapolateSequencePart2(arr)
		} else {
			val = extrapolateSequence(arr)
		}
		sum += val
	}
	fmt.Printf("%d\n", sum)
}
