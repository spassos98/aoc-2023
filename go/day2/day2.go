package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	red   int
	green int
	blue  int
}

func (m *Move) updateMove(value int, color string) {
	switch color {
	case "red":
		m.red = m.red + value
	case "green":
		m.green = m.green + value
	case "blue":
		m.blue = m.blue + value
	}
}

func (m *Move) updateMaxMove(move Move) {
	m.red = max(m.red, move.red)
	m.green = max(m.green, move.green)
	m.blue = max(m.blue, move.blue)
}

var redCubes int = 12
var greenCubes int = 13
var blueCubes int = 14

func getGameNumber(gameString string) int {
	numberString := gameString[5:]
	number, _ := strconv.Atoi(numberString)
	return number
}

func parseMove(moveString string) Move {
	cubeCountStrings := strings.Split(moveString, ",")
	move := Move{0, 0, 0}
	for _, cubeCountString := range cubeCountStrings {
		numberAndColor := strings.Split(strings.Trim(cubeCountString, " "), " ")
		number, _ := strconv.Atoi(numberAndColor[0])
		color := numberAndColor[1]
		move.updateMove(number, color)
	}
	return move
}

func getMoves(movesString string) []Move {
	movesStrings := strings.Split(movesString, ";")
	var moves []Move
	for _, moveString := range movesStrings {
		currMove := parseMove(moveString)
		moves = append(moves, currMove)
	}
	return moves
}

func isValidMove(move Move) bool {
	return move.red <= redCubes && move.blue <= blueCubes && move.green <= greenCubes
}

func isValidGame(gameLine string) int {
	firstSplit := strings.Split(gameLine, ":")
	gameString := firstSplit[0]
	gameNumber := getGameNumber(gameString)
	movesString := firstSplit[1]
	moves := getMoves(movesString)
	for _, move := range moves {
		if !isValidMove(move) {
			return 0
		}
	}
	return gameNumber
}

func getGameValue(gameLine string) int {
	firstSplit := strings.Split(gameLine, ":")
	movesString := firstSplit[1]
	moves := getMoves(movesString)
	maxMove := Move{0, 0, 0}
	for _, move := range moves {
		maxMove.updateMaxMove(move)
	}
	return maxMove.red * maxMove.blue * maxMove.green
}

func day2Part2() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		gameLine := scanner.Text()
		sum += getGameValue(gameLine)
	}
	fmt.Printf("%d\n", sum)
}

func day2() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		gameLine := scanner.Text()
		sum += isValidGame(gameLine)
	}
	fmt.Printf("%d\n", sum)
}
