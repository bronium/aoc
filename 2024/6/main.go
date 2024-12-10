package main

import (
	"fmt"
	"os"
	"strings"
)

type Guard struct {
	x, y, directionX, directionY int
}

func (self *Guard) move() {
	self.x += self.directionX
	self.y += self.directionY
}

func (self *Guard) rollback() {
	self.x -= self.directionX
	self.y -= self.directionY
}

func (self *Guard) turn() {
	switch true {
	case self.directionX == 0 && self.directionY == -1:
		self.directionX = 1
		self.directionY = 0
	case self.directionX == 1 && self.directionY == 0:
		self.directionX = 0
		self.directionY = 1
	case self.directionX == 0 && self.directionY == 1:
		self.directionX = -1
		self.directionY = 0
	case self.directionX == -1 && self.directionY == 0:
		self.directionX = 0
		self.directionY = -1
	}
}

func (self Guard) recordMovementAxis(direction BeenAxis) BeenAxis {
	if self.directionY == -1 {
		direction.up = true
	}
	if self.directionY == 1 {
		direction.down = true
	}
	if self.directionX == -1 {
		direction.left = true
	}
	if self.directionX == 1 {
		direction.right = true
	}

	return direction
}

func (self Guard) visitedBefore(axis BeenAxis) bool {
	if axis.up && self.directionY == -1 {
		return true
	}

	if axis.down && self.directionY == 1 {
		return true
	}

	if axis.left && self.directionX == -1 {
		return true
	}

	if axis.right && self.directionX == 1 {
		return true
	}

	return false
}

type BeenAxis struct {
	up, down, left, right bool
}

type Tile struct {
	state string
	been  BeenAxis
}

func isInfiniteLoop(guard Guard, matrix [][]Tile) bool {
	for {
		newX, newY := guard.x+guard.directionX, guard.y+guard.directionY
		if newX < 0 || newY < 0 || newY >= len(matrix) || newX >= len(matrix[0]) {
			break
		}
		if matrix[newY][newX].state == "#" {
			guard.turn()
			continue
		}

		if guard.visitedBefore(matrix[guard.y][guard.x].been) {
			return true
		}

		matrix[guard.y][guard.x].been = guard.recordMovementAxis(matrix[guard.y][guard.x].been)
		guard.move()
	}

	return false
}

func copyMatrix(src [][]Tile) [][]Tile {
	dest := [][]Tile{}
	for _, row := range src {
		tiles := []Tile{}
		for _, el := range row {
			tile := Tile{el.state, BeenAxis{el.been.up, el.been.down, el.been.left, el.been.right}}
			tiles = append(tiles, tile)
		}
		dest = append(dest, tiles)
	}
	return dest
}

func printMatrix(matrix [][]Tile) {
	for _, row := range matrix {
		fmt.Println()
		for _, el := range row {
			fmt.Print(el.state, " ")
		}
	}
	fmt.Println()
}

func main() {
	data, _ := os.ReadFile("based_debug")
	lines := strings.Split(string(data), "\n")

	matrix := [][]Tile{}
	guard := Guard{}
	for y, line := range lines {
		runes := []Tile{}
		for x, item := range line {
			tile := Tile{state: string(item)}
			runes = append(runes, tile)
			if item == '^' {
				guard = Guard{x, y, 0, -1}
			}
		}
		matrix = append(matrix, runes)
	}

	count := 0
	for i, row := range matrix {
		for j, el := range row {
			if el.state == "." {

				matrixCopy := copyMatrix(matrix)
				matrixCopy[i][j].state = "#"

				if isInfiniteLoop(guard, matrixCopy) {
					count++
				}
			}
		}
	}
	fmt.Println(count)

	// printMatrix(matrix)
	// fmt.Printf("isInfiniteLoop(guard, matrixCopy): %v\n", isInfiniteLoop(guard, matrix))

}
