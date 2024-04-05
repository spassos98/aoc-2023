package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Galaxy struct {
	row int
	col int
}

func abs(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}

func galaxyDistance(a Galaxy, b Galaxy) int {
	return abs(a.row-b.row) + abs(a.col-b.col)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var originalGalaxy []string
	for scanner.Scan() {
		line := scanner.Text()
		originalGalaxy = append(originalGalaxy, line)
	}
	var emptyRows []int
	for row, line := range originalGalaxy {
		isEmpty := true
		for _, val := range line {
			if val == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyRows = append(emptyRows, row)
		}
	}

	var emptyColumns []int
	for col := 0; col < len(originalGalaxy[0]); col += 1 {
		isEmpty := true
		for row := 0; row < len(originalGalaxy); row += 1 {
			if originalGalaxy[row][col] == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyColumns = append(emptyColumns, col)
		}

	}
	fmt.Printf("Empty rows %v\n", emptyRows)
	fmt.Printf("Empty cols %v\n", emptyColumns)

	expandedGalaxyCols := len(emptyColumns) + len(originalGalaxy)
	var expandedGalaxy []string
	for row, line := range originalGalaxy {
		if slices.Contains(emptyRows, row) {
			emptyRow := strings.Repeat(".", expandedGalaxyCols)
			expandedGalaxy = append(expandedGalaxy, emptyRow)
		}
		var expandedLine []rune
		for col, val := range line {
			if slices.Contains(emptyColumns, col) {
				expandedLine = append(expandedLine, val)
			}
			expandedLine = append(expandedLine, val)
		}
		expandedGalaxy = append(expandedGalaxy, string(expandedLine))
	}

	for _, line := range expandedGalaxy {
		fmt.Printf("%v\n", line)
	}

	var galaxies []Galaxy
	for row, line := range expandedGalaxy {
		for col, val := range line {
			if val == '#' {
				galaxies = append(galaxies, Galaxy{row, col})
			}
		}
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += galaxyDistance(galaxies[i], galaxies[j])
		}
	}
	fmt.Printf("%d\n", sum)
}
