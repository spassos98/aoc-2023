package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isDigit(testingString string, startPos int) int {
	digits := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for idx, digitString := range digits {
		limit := min(len(testingString), startPos+len(digitString))
		subStr := strings.Clone(testingString)[startPos:limit]
		if subStr == digitString {
			return idx + 1
		}
	}
	return 0
}

func getFirstAndLastDigit(calibrationValue string) (int, int) {
	firstDigit := 0
	lastDigit := 0
	for i := 0; i < len(calibrationValue); i++ {
		char := calibrationValue[i]
		digitValue := isDigit(calibrationValue, i)

		if digitValue == 0 && char >= '1' && char <= '9' {
			digitValue = int(char - '0')
		}

		if digitValue != 0 {
			if firstDigit != 0 {
				lastDigit = digitValue
			} else {
				firstDigit = digitValue
			}
		}
	}

	if lastDigit == 0 {
		lastDigit = firstDigit
	}

	return firstDigit, lastDigit
}

func day1() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		firstDigit, lastDigit := getFirstAndLastDigit(scanner.Text())
		sum += firstDigit*10 + lastDigit
	}
	fmt.Println(sum)
}
