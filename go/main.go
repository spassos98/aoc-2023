package main

import (
	"fmt"
	"os"
)

func main() {
	dayName := os.Args[1]
	fmt.Printf("Running code for day %s\n", dayName)
	switch dayName {
	case "1":
		day1()
	case "2":
		day2()
	}
}
