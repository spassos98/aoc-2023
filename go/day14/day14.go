package main

import (
	"bufio"
	"fmt"
	"os"
)

func getColumnWeight(matrix []string, col int) int {
	nextRockPos := 0
	totalWeight := 0
	for i := 0; i < len(matrix); i++ {
		currSpace := matrix[i][col]
		if currSpace == '#' {
			nextRockPos = i + 1
		} else if currSpace == 'O' {
			totalWeight += len(matrix[0]) - nextRockPos
			nextRockPos += 1
		}
	}
	return totalWeight
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var matrix []string
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}
	sum := 0
	ncols := len(matrix[0])
	for col := 0; col < ncols; col++ {
		sum += getColumnWeight(matrix, col)
	}
	fmt.Println(sum)
}
