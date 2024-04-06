package main

import (
	"bufio"
	"fmt"
	"os"
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

func galaxyDistance(a Galaxy, b Galaxy, emptyRows []int, emptyColums []int, emptyValue int) int {
	baseDistance := abs(a.row-b.row) + abs(a.col-b.col)
	startRow := min(a.row, b.row)
	endRow := max(a.row, b.row)
	for _, row := range emptyRows {
		if startRow <= row && row <= endRow {
			baseDistance += emptyValue - 1
		}
	}

	startCol := min(a.col, b.col)
	endCol := max(a.col, b.col)
	for _, col := range emptyColums {
		if startCol <= col && col <= endCol {
			baseDistance += emptyValue - 1
		}
	}
	return baseDistance
}

func main() {
	part := os.Args[1]
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

	var galaxies []Galaxy
	for row, line := range originalGalaxy {
		for col, val := range line {
			if val == '#' {
				galaxies = append(galaxies, Galaxy{row, col})
			}
		}
	}

	emptyValue := 0
	if part == "2" {
		emptyValue = 1000000
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += galaxyDistance(galaxies[i], galaxies[j], emptyRows, emptyColumns, emptyValue)
		}
	}
	fmt.Printf("%d\n", sum)
}
