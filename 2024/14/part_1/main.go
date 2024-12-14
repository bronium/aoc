package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const Width = 101
const Height = 103
const Seconds = 100

func mod(a, b int) int {
	return (a%b + b) % b
}

type Robot struct {
	x, y, velX, velY int
}

func (robot *Robot) move() {
	robot.x = mod(robot.x+robot.velX, Width)
	robot.y = mod(robot.y+robot.velY, Height)
}

func (robot Robot) quadrant() (int, bool) {
	switch true {
	case robot.x < Width/2 && robot.y < Height/2:
		return 0, true
	case robot.x > Width/2 && robot.y < Height/2:
		return 1, true
	case robot.x < Width/2 && robot.y > Height/2:
		return 2, true
	case robot.x > Width/2 && robot.y > Height/2:
		return 3, true
	default:
		return -1, false
	}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	security := [4]int{}
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`-?\d+`)
		values := re.FindAllString(line, 4)
		var nums [4]int
		for i, value := range values {
			nums[i], _ = strconv.Atoi(value)
		}
		robot := Robot{nums[0], nums[1], nums[2], nums[3]}

		for range Seconds {
			robot.move()
		}
		if quadrant, ok := robot.quadrant(); ok {
			security[quadrant] += 1
		}
	}

	total := 1
	for _, value := range security {
		total *= value
	}

	fmt.Println(total)
}
