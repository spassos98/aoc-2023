package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func pickNode(pos int, seen map[int]bool, graph [][]int) int {
	option1 := graph[pos][0]
	option2 := graph[pos][1]
	_, seen1 := seen[option1]
	_, seen2 := seen[option2]
	if !seen2 {
		return option2
	}
	if !seen1 {
		return option1
	}
	return -1
}

func maxPath(initialPosition int, startPos []int, graph [][]int) (int, []int) {
	currPos := startPos
	steps := 1
	seen := make(map[int]bool)
	seen[initialPosition] = true
	loopPositions := make([]int, 1)
	loopPositions[0] = initialPosition
	for {
		node1 := currPos[0]
		node2 := currPos[1]
		loopPositions = append(loopPositions, node1)
		loopPositions = append(loopPositions, node2)
		if _, ok := seen[node1]; ok {
			break
		}
		if _, ok := seen[node2]; ok {
			break
		}
		seen[node1] = true
		seen[node2] = true
		nextNode1 := pickNode(node1, seen, graph)
		nextNode2 := pickNode(node2, seen, graph)
		if nextNode1 == -1 || nextNode2 == -1 {
			break
		}
		currPos[0] = nextNode1
		currPos[1] = nextNode2
		steps += 1
	}
	return steps, loopPositions
}

func buildGraph(matrix []string) (int, [][]int) {
	var graph [][]int
	var startPosition int
	rowLen := len(matrix[0])
	for row, matrixLine := range matrix {
		for col, symbol := range matrixLine {
			var neigh []int
			currPos := row*rowLen + col
			north := currPos - rowLen
			south := currPos + rowLen
			east := currPos + 1
			west := currPos - 1
			addNorth := false
			addSouth := false
			addEast := false
			addWest := false
			switch symbol {
			case '|':
				addNorth = true
				addSouth = true
			case '-':
				addEast = true
				addWest = true
			case 'L':
				addNorth = true
				addEast = true
			case 'J':
				addNorth = true
				addWest = true
			case '7':
				addSouth = true
				addWest = true
			case 'F':
				addSouth = true
				addEast = true
			case 'S':
				startPosition = currPos

			}
			if addNorth && north > 0 {
				neigh = append(neigh, north)
			}
			if addSouth && south < rowLen*len(matrix) {
				neigh = append(neigh, south)
			}
			if addEast && east < rowLen*len(matrix) && currPos%rowLen != (rowLen-1) {
				neigh = append(neigh, east)
			}
			if addWest && west > 0 && currPos%rowLen != 0 {
				neigh = append(neigh, west)
			}
			graph = append(graph, neigh)
		}
	}
	return startPosition, graph
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var matrix []string
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}

	rowLen := len(matrix[0])
	startPosition, graph := buildGraph(matrix)

	var startNeigh []int
	for x := 0; x < rowLen*len(matrix); x++ {
		neigh := graph[x]
		for idx := range neigh {
			if neigh[idx] == startPosition {
				startNeigh = append(startNeigh, x)
			}
		}
	}

	part := os.Args[1]
	sol, loopPositions := maxPath(startPosition, startNeigh, graph)
	if part == "1" {
		fmt.Printf("%d\n", sol)
		return
	}

	var transformedMatrix []string
	for row, line := range matrix {
		var transformedLine []rune
		for col, val := range line {
			currPos := row*rowLen + col
			if slices.Contains(loopPositions, currPos) {
				if val == '7' {
					transformedLine = append(transformedLine, '┐')
				} else if val == 'L' {
					transformedLine = append(transformedLine, '└')
				} else if val == 'F' {
					transformedLine = append(transformedLine, '┌')
				} else if val == 'J' {
					transformedLine = append(transformedLine, '┘')
				} else if val == '|' {
					transformedLine = append(transformedLine, '│')
				} else if val == '-' {
					transformedLine = append(transformedLine, '─')
				} else {
					transformedLine = append(transformedLine, val)
				}
			} else {
				transformedLine = append(transformedLine, '.')
			}
		}
		transformedMatrix = append(transformedMatrix, string(transformedLine))
	}
	for _, line := range transformedMatrix {
		fmt.Println(line)
	}

	nEnclosed := 0

	for _, line := range transformedMatrix {
		nEdges := 0
		openChar := '-'
		for _, val := range line {
			if val == '│' {
				nEdges += 1
			} else if val == '.' && nEdges%2 == 1 {
				nEnclosed += 1
			} else if val == '└' || val == '┌' {
				openChar = val
			} else if val == '┘' || val == '┐' {
				if openChar == '└' && val == '┐' {
					nEdges += 1
				} else if openChar == '┌' && val == '┘' {
					nEdges += 1
				}
			}
		}
	}
	fmt.Printf("%d\n", nEnclosed)
}
