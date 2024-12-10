package main

import (
	"fmt"
	"os"
	"strings"
)

type Location struct {
	x, y int
}

func nextLocation(loc Location, dir string) Location {
	xDiff := 0
	yDiff := 0
	switch dir {
	case ">":
		xDiff = 1
	case "<":
		xDiff = -1
	case "^":
		yDiff = 1
	case "v":
		yDiff = -1
	}
	return Location{loc.x + xDiff, loc.y + yDiff}
}

func countHouses() int {
	data, _ := os.ReadFile("smol.txt")
	input := strings.Trim(string(data), "\n")

	houses := map[Location]bool{}
	currentSanta := Location{0, 0}
	currentRobo := Location{0, 0}
	houses[currentSanta] = true
	for ind, dir := range strings.Split(input, "") {
		if ind%2 == 0 {
			currentSanta = nextLocation(currentSanta, dir)
			houses[currentSanta] = true
		} else {
			currentRobo = nextLocation(currentRobo, dir)
			houses[currentRobo] = true
		}
	}

	return len(houses) 
}

func main() {
	fmt.Println(countHouses())
}
