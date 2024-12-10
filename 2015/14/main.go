package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	maxDistance := 0.0
	timestamp := 2503.0
	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile(`\d+`)
		values := re.FindAllString(line, 3)

		speed, _ := strconv.ParseFloat(values[0], 64)
		flyingTime, _ := strconv.ParseFloat(values[1], 64)
		restingTime, _ := strconv.ParseFloat(values[2], 64)

		cycleTime := flyingTime + restingTime
		intervals := math.Floor(timestamp / cycleTime)

		remainder := timestamp - intervals*cycleTime
		remainderFlying := max(0, min(flyingTime, remainder))

		totalFlying := intervals*flyingTime + remainderFlying
		totalDistance := totalFlying * speed

		maxDistance = max(maxDistance, totalDistance)
	}
	fmt.Println(maxDistance)
}
