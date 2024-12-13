package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Button struct {
	x, y, price int
}
type Machine struct {
	buttonA, buttonB Button
	targetX, targetY int
}

const Offset = 10000000000000

func determinant(matrix [2][2]int) int {
	return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
}

func countPrice(machine Machine) int {
	denominator := determinant([2][2]int{
		{machine.targetX, machine.buttonB.x},
		{machine.targetY, machine.buttonB.y},
	})
	numerator := determinant([2][2]int{
		{machine.buttonA.x, machine.buttonB.x},
		{machine.buttonA.y, machine.buttonB.y},
	})

	if denominator%numerator != 0 {
		return 0
	}

	a := denominator / numerator
	b := (machine.targetX - a*machine.buttonA.x) / machine.buttonB.x

	return a*machine.buttonA.price + b*machine.buttonB.price
}

func main() {
	file, _ := os.ReadFile("input.txt")
	machines := strings.Split(string(file), "\n\n")
	total := 0
	for _, machine := range machines {
		re := regexp.MustCompile(`\d+`)
		values := re.FindAllString(machine, 6)
		var nums [6]int
		for i, value := range values {
			res, _ := strconv.Atoi(value)
			nums[i] = res
		}

		current := Machine{
			Button{nums[0], nums[1], 3},
			Button{nums[2], nums[3], 1},
			Offset + nums[4], Offset + nums[5],
		}
		total += countPrice(current)
	}

	fmt.Println(total)
}
