package main

import (
	"fmt"
	"os"
)

func main() {
	dayName := os.Args[1]
	part := "1"
	if len(os.Args) >= 3 {
		part = os.Args[2]
	}
	fmt.Printf("Running code for day %s part %s\n", dayName, part)
	switch dayName {
	case "1":
		day1()
	case "2":
		if part == "1" {
			day2()
		} else {
			day2Part2()
		}
	}
}
