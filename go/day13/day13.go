package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkVerticalReflection(matrix []string, pos int) bool {
	l := pos
	r := pos + 1
	for l >= 0 && r < len(matrix) {
		if matrix[l] != matrix[r] {
			return false
		}
		l--
		r++
	}
	return true
}

func findDuplicateRow(matrix []string) int {
	for idx, currLine := range matrix {
		if idx < len(matrix)-1 {
			nextLine := matrix[idx+1]
			if currLine == nextLine && checkVerticalReflection(matrix, idx) {
				return idx + 1
			}
		}
	}
	return 0
}

func checkHorizontalReflection(matrix []string, pos int) bool {
	l := pos
	r := pos + 1
	for l >= 0 && r < len(matrix[0]) {
		lCol := getCol(matrix, l)
		rCol := getCol(matrix, r)
		if lCol != rCol {
			return false
		}
		l--
		r++
	}
	return true
}

func getCol(matrix []string, pos int) string {
	var col []byte
	for row := 0; row < len(matrix); row += 1 {
		col = append(col, matrix[row][pos])
	}
	return string(col)
}

func findDuplicateColumn(matrix []string) int {
	for col := 0; col < len(matrix[0]); col++ {
		if col < len(matrix[0])-1 {
			currCol := getCol(matrix, col)
			nextCol := getCol(matrix, col+1)
			if currCol == nextCol && checkHorizontalReflection(matrix, col) {
				return col + 1
			}
		}
	}
	return 0
}

func getMatrixValue(matrix []string) int {
	summaryVal := 0
	colPos := findDuplicateColumn(matrix)
	if colPos != 0 {
		summaryVal = colPos
	} else {
		rowPos := findDuplicateRow(matrix)
		if rowPos != 0 {
			summaryVal = rowPos * 100
		}
	}
	return summaryVal
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var matrix []string
	sum := 0
	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sum += getMatrixValue(matrix)
			matrix = make([]string, 0)
			count += 1
			continue
		}
		matrix = append(matrix, line)
	}
	sum += getMatrixValue(matrix)
	fmt.Printf("%d\n", sum)
}
