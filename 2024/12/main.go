package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}
type Matrix [][]string

type Region struct {
	area, perimeter, sides int
}

var Directions = []Point{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

var been = map[Point]bool{}
var uniqIndex int

func readMatrix(filename string) Matrix {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	matrix := Matrix{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for _, ch := range line {
			row = append(row, string(ch))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func (matrix *Matrix) calculate(target string, x, y int) (int, int) {
	(*matrix)[y][x] = strconv.Itoa(uniqIndex)
	been[Point{x, y}] = true

	totalArea := 1
	totalPerimeter := 0

	for _, direction := range Directions {
		newX := x + direction.x
		newY := y + direction.y

		if been[Point{newX, newY}] {
			continue
		}

		if newX < 0 || newY < 0 || newX >= len((*matrix)[0]) || newY >= len(*matrix) {
			totalPerimeter++
			continue
		}

		if (*matrix)[newY][newX] == target {
			area, perimeter := matrix.calculate(target, newX, newY)
			totalArea += area
			totalPerimeter += perimeter
		} else {
			totalPerimeter++
		}

	}

	return totalArea, totalPerimeter
}

func main() {
	matrix := readMatrix("input")

	uniqIndex = 0
	regions := map[string]Region{}
	for i, row := range matrix {
		for j, val := range row {
			_, isInt := strconv.Atoi(val)
			if isInt != nil {
				been = map[Point]bool{}
				area, perimeter := matrix.calculate(val, j, i)
				regions[strconv.Itoa(uniqIndex)] = Region{area: area, perimeter: perimeter}
				uniqIndex++
			}
		}
	}

	for i, row := range matrix {
		var currentTop string
		var currentBottom string
		var currentLeft string
		var currentRight string
		for j, val := range row {
			if i-1 < 0 || matrix[i-1][j] != val {
				if currentTop != val {
					currentTop = val
					region := regions[currentTop]
					region.sides++
					regions[currentTop] = region
				}
			} else {
				currentTop = ""
			}

			if i+1 >= len(matrix) || matrix[i+1][j] != val {
				if currentBottom != val {
					currentBottom = val
					region := regions[currentBottom]
					region.sides++
					regions[currentBottom] = region
				}
			} else {
				currentBottom = ""
			}

			val := matrix[j][i]
			if i-1 < 0 || matrix[j][i-1] != val {
				if currentLeft != val {
					currentLeft = val
					region := regions[currentLeft]
					region.sides++
					regions[currentLeft] = region
				}
			} else {
				currentLeft = ""
			}

			val = matrix[j][i]
			if i+1 >= len(matrix) || matrix[j][i+1] != val {
				if currentRight != val {
					currentRight = val
					region := regions[currentRight]
					region.sides++
					regions[currentRight] = region
				}
			} else {
				currentRight = ""
			}
		}
	}

	result := 0
	for _, region := range regions {
		result += region.area * region.sides
	}

	fmt.Println(result)
}
