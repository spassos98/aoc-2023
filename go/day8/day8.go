package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MovementOption struct {
	left  string
	right string
}

var movementMap map[string]MovementOption

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func getNextPosition(currPos string, movement byte) string {
	if movement == 'L' {
		return movementMap[currPos].left
	}
	return movementMap[currPos].right
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	movements := scanner.Text()
	scanner.Scan()
	movementMap = make(map[string]MovementOption)
	var currPos []string
	for scanner.Scan() {
		movementLine := scanner.Text()
		firstSplit := strings.Split(movementLine, "=")
		source := strings.Trim(firstSplit[0], " ")
		optionsSplit := strings.Split(firstSplit[1], ",")
		leftOption := optionsSplit[0][2:]
		rightOption := optionsSplit[1][1:4]
		movementMap[source] = MovementOption{leftOption, rightOption}
		if source[2] == 'A' {
			currPos = append(currPos, source)
		}
	}

	var stepsToAchieveEnd []int
	for _, startNode := range currPos {
		currNode := startNode
		steps := 0
		for {
			stepPos := steps % len(movements)
			movement := movements[stepPos]
			currNode = getNextPosition(currNode, movement)
			steps += 1
			if currNode[2] == 'Z' {
				stepsToAchieveEnd = append(stepsToAchieveEnd, steps)
				break
			}
		}
	}

	sum := LCM(stepsToAchieveEnd[0], stepsToAchieveEnd[1], stepsToAchieveEnd...)
	fmt.Printf("%d\n", sum)
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	movements := scanner.Text()
	scanner.Scan()
	movementMap = make(map[string]MovementOption)
	for scanner.Scan() {
		movementLine := scanner.Text()
		firstSplit := strings.Split(movementLine, "=")
		source := strings.Trim(firstSplit[0], " ")
		optionsSplit := strings.Split(firstSplit[1], ",")
		leftOption := optionsSplit[0][2:]
		rightOption := optionsSplit[1][1:4]
		movementMap[source] = MovementOption{leftOption, rightOption}
	}

	currPos := "AAA"
	steps := 0
	for currPos != "ZZZ" {
		stepPos := steps % len(movements)
		movement := movements[stepPos]
		if movement == 'L' {
			currPos = movementMap[currPos].left
		} else {
			currPos = movementMap[currPos].right
		}
		steps += 1
	}
	fmt.Printf("%d\n", steps)
}

func main() {
	part := os.Args[1]
	if part == "2" {
		part2()
	} else {
		part1()
	}
}
