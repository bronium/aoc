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
const Seconds = 10000
const VerticalThreshhold = 10

func mod(a, b int) int {
	return (a%b + b) % b
}

type Matrix [Height][Width]int

func (matrix Matrix) print() {
	for _, row := range matrix {
		for _, value := range row {
			if value == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%v ", value)
			}
		}
		fmt.Println()
	}
}

func (matrix Matrix) hasPattern() bool {
	for i := 0; i < Width; i++ {
		count := 0
		for j := 0; j < Height; j++ {
			if matrix[j][i] > 0 {
				count++
				if count >= VerticalThreshhold {
					return true
				}
			} else {
				count = 0
			}
		}
	}
	return false
}

type Robot struct {
	x, y, velX, velY int
}

func (robot *Robot) move() {
	robot.x = mod(robot.x+robot.velX, Width)
	robot.y = mod(robot.y+robot.velY, Height)
}

func readFromFile(filename string) []Robot {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	robots := []Robot{}
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`-?\d+`)
		values := re.FindAllString(line, 4)
		var nums [4]int
		for i, value := range values {
			nums[i], _ = strconv.Atoi(value)
		}
		robot := Robot{nums[0], nums[1], nums[2], nums[3]}

		robots = append(robots, robot)

	}
	return robots
}

func main() {
	robots := readFromFile("input.txt")
	for i := range Seconds {
		matrix := Matrix{}
		for i, robot := range robots {
			robots[i].move()
			matrix[robot.y][robot.x] += 1
		}

		if matrix.hasPattern() {
			println(i + 1)
			matrix.print()
		}
	}
}
