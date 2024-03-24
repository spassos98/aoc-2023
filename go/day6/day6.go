package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func getRaceWays(race Race) int {
	sum := 0
	for timePressed := 0; timePressed <= race.time; timePressed++ {
		speed := timePressed
		distanceTraveled := speed * (race.time - timePressed)
		if distanceTraveled > race.distance {
			sum += 1
		}
	}
	return sum
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	timeLine := scanner.Text()
	scanner.Scan()
	distanceLine := scanner.Text()
	timesStr := strings.Fields(strings.Split(timeLine, ":")[1])
	distanceStr := strings.Fields(strings.Split(distanceLine, ":")[1])

	var races []Race
	for idx, raceTime := range timesStr {
		timeValue, _ := strconv.Atoi(raceTime)
		distanceValue, _ := strconv.Atoi(distanceStr[idx])
		race := Race{timeValue, distanceValue}
		races = append(races, race)
	}

	sum := 1
	for _, race := range races {
		sum *= getRaceWays(race)
	}

	fmt.Printf("%d\n", sum)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	timeLine := scanner.Text()
	scanner.Scan()
	distanceLine := scanner.Text()
	timesStr := strings.Join(strings.Fields(strings.Split(timeLine, ":")[1]), "")
	distanceStr := strings.Join(strings.Fields(strings.Split(distanceLine, ":")[1]), "")

	timeValue, _ := strconv.Atoi(timesStr)
	timeValueF := float64(timeValue)
	distanceValue, _ := strconv.Atoi(distanceStr)
	distanceValueF := float64(distanceValue)
	num1 := -1*timeValueF + math.Sqrt(timeValueF*timeValueF-(4*distanceValueF))
	num2 := -1*timeValueF - math.Sqrt(timeValueF*timeValueF-(4*distanceValueF))
	den := -2.0
	sol1 := num1 / den
	sol2 := num2 / den
	var start float64
	var end float64
	if sol1 < sol2 {
		start = math.Ceil(sol1 + 0.0001)
		end = math.Floor(sol2 - 0.0001)
	} else {
		start = math.Ceil(sol2 + 0.0001)
		end = math.Floor(sol1 - 0.0001)
	}
	solRange := end - start + 1
	fmt.Printf("%d\n", int(solRange))
}

func main() {
	part := os.Args[1]
	if part == "2" {
		part2()
	} else {
		part1()
	}
}
