package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var memo [][][][]int

func arrangementIsValid(arrangement []rune, expectedGroups []int) int {
	isGroup := false
	groupCount := 0
	var groups []int
	for _, val := range arrangement {
		if !isGroup && val == '#' {
			isGroup = true
			groupCount = 1
		} else if isGroup && val == '#' {
			groupCount += 1
		} else if isGroup && val == '.' {
			isGroup = false
			groups = append(groups, groupCount)
		}
	}
	if isGroup {
		groups = append(groups, groupCount)
	}
	if slices.Equal(groups, expectedGroups) {
		return 1
	}
	return 0
}

func setArrangement(arrangement []rune, pos int, groups []int) int {
	if pos == len(arrangement) {
		return arrangementIsValid(arrangement, groups)
	}
	if arrangement[pos] != '?' {
		return setArrangement(arrangement, pos+1, groups)
	}

	operationalArr := make([]rune, len(arrangement))
	copy(operationalArr, arrangement)
	operationalArr[pos] = '.'

	damagedArr := make([]rune, len(arrangement))
	copy(damagedArr, arrangement)
	damagedArr[pos] = '#'

	operationalValue := setArrangement(operationalArr, pos+1, groups)
	damagedValue := setArrangement(damagedArr, pos+1, groups)

	return operationalValue + damagedValue
}

func countArrangements(line string) int {
	split := strings.Split(line, " ")
	arrLine := split[0]
	var arrangement []rune
	for _, val := range arrLine {
		arrangement = append(arrangement, val)
	}

	groupLine := split[1]
	strGroups := strings.Split(groupLine, ",")
	var groups []int
	for _, val := range strGroups {
		number, _ := strconv.Atoi(val)
		groups = append(groups, number)
	}
	value := setArrangement(arrangement, 0, groups)
	return value
}

func setArrangement2(arrangement []rune, pos int, expectedGroups []int, isGroup bool, groupCount int, groupPos int) int {

	if pos > 0 {
		val := arrangement[pos-1]

		if !isGroup && val == '#' {
			isGroup = true
			groupCount = 1
		} else if isGroup && val == '#' {
			groupCount += 1
		} else if isGroup && val == '.' {
			if groupPos >= len(expectedGroups) {
				return 0
			}
			if expectedGroups[groupPos] != groupCount {
				return 0
			}
			isGroup = false
			groupPos += 1
			groupCount = 0
		}
	}

	isGroupPos := 0
	if isGroup {
		isGroupPos = 1
	}
	if memo[pos][groupPos][groupCount][isGroupPos] != -1 {
		return memo[pos][groupPos][groupCount][isGroupPos]
	}

	remainingChars := len(arrangement) - pos
	remainingGroups := len(expectedGroups) - groupPos
	if remainingGroups > 1+remainingChars/2 {
		memo[pos][groupPos][groupCount][isGroupPos] = 0
		return 0
	}

	if pos == len(arrangement) {
		// fmt.Printf("Arr %s\n", string(arrangement))
		// fmt.Printf("\tEx Groups%v\n", expectedGroups)
		// fmt.Printf("\tisGroup %t groupCount %d groupPos %d\n", isGroup, groupCount, groupPos)
		if isGroup {
			if groupPos == len(expectedGroups)-1 && expectedGroups[groupPos] == groupCount {
				memo[pos][groupPos][groupCount][isGroupPos] = 1
				return 1
			}
			return 0
		} else {
			if groupPos == len(expectedGroups) {
				memo[pos][groupPos][groupCount][isGroupPos] = 1
				return 1
			}
			return 0
		}
	}

	if arrangement[pos] != '?' {
		return setArrangement2(arrangement, pos+1, expectedGroups, isGroup, groupCount, groupPos)
	}

	operationalArr := make([]rune, len(arrangement))
	copy(operationalArr, arrangement)
	operationalArr[pos] = '.'

	damagedArr := make([]rune, len(arrangement))
	copy(damagedArr, arrangement)
	damagedArr[pos] = '#'

	operationalValue := setArrangement2(operationalArr, pos+1, expectedGroups, isGroup, groupCount, groupPos)
	damagedValue := setArrangement2(damagedArr, pos+1, expectedGroups, isGroup, groupCount, groupPos)

	possibleArrs := operationalValue + damagedValue

	memo[pos][groupPos][groupCount][isGroupPos] = possibleArrs
	return possibleArrs
}

func countArrangements2(line string) int {
	split := strings.Split(line, " ")
	arrLine := split[0]

	newArrLine := arrLine + strings.Repeat("?"+arrLine, 4)

	var arrangement []rune
	for _, val := range newArrLine {
		arrangement = append(arrangement, val)
	}

	groupLine := split[1]
	strGroups := strings.Split(groupLine, ",")
	var groups []int
	for _, val := range strGroups {
		number, _ := strconv.Atoi(val)
		groups = append(groups, number)
	}

	newGroupsLen := len(groups) * 5
	newGroups := make([]int, newGroupsLen)
	for i := 0; i < newGroupsLen; i += len(groups) {
		copy(newGroups[i:], groups)
	}

	value := setArrangement2(arrangement, 0, newGroups, false, 0, 0)
	return value
}

func resetMemo() {
	l := len(memo)
	for a := 0; a < l; a++ {
		for b := 0; b < l; b++ {
			for c := 0; c < l; c++ {
				memo[a][b][c][0] = -1
				memo[a][b][c][1] = -1
			}

		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	count := 0
	l := 150
	part := os.Args[1]

	for a := 0; a < l; a++ {
		var newArr [][][]int
		for b := 0; b < l; b++ {
			var newArrb [][]int
			for c := 0; c < l; c++ {
				newArrb = append(newArrb, []int{-1, -1})
			}
			newArr = append(newArr, newArrb)

		}
		memo = append(memo, newArr)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if part == "2" {
			resetMemo()
			sum += countArrangements2(line)
		} else {
			sum += countArrangements(line)
		}
		count += 1
		fmt.Println(count)
	}
	fmt.Printf("%d\n", sum)
}
