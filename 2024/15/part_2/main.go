package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Matrix [][]rune

var Directions = map[rune]Point{
	'^': {0, -1},
	'v': {0, 1},
	'>': {1, 0},
	'<': {-1, 0},
}

var dog = Point{}

func (matrix *Matrix) move(direction rune) {
	delta := Directions[direction]
	nextPos := Point{dog.x + delta.x, dog.y + delta.y}

	moveDog := func() {
		(*matrix)[dog.y][dog.x] = '.'
		dog = nextPos
		(*matrix)[dog.y][dog.x] = '@'
	}

	switch (*matrix)[nextPos.y][nextPos.x] {
	case '.':
		moveDog()
	case '[', ']':
		if matrix.moveBox(nextPos, direction, false) {
			moveDog()
		}
	}
}

func (matrix *Matrix) moveBox(coord Point, direction rune, dry bool) bool {
	left, right := Point{}, Point{}
	switch (*matrix)[coord.y][coord.x] {
	case '[':
		left = coord
		right = Point{coord.x + 1, coord.y}
	case ']':
		left = Point{coord.x - 1, coord.y}
		right = coord
	}

	delta := Directions[direction]
	nextLeftPos := Point{left.x + delta.x, left.y + delta.y}
	nextRightPos := Point{right.x + delta.x, right.y + delta.y}
	nextLeftValue := (*matrix)[nextLeftPos.y][nextLeftPos.x]
	nextRightValue := (*matrix)[nextRightPos.y][nextRightPos.x]

	if nextLeftValue == '#' || nextRightValue == '#' {
		return false
	}

	switch direction {
	case '^', 'v':
		if nextLeftValue == '[' {
			if !matrix.moveBox(nextLeftPos, direction, dry) {
				return false
			}
		}

		if nextLeftValue == ']' && !matrix.moveBox(nextLeftPos, direction, true) ||
			nextRightValue == '[' && !matrix.moveBox(nextRightPos, direction, true) {
			return false
		}

		if nextLeftValue == ']' {
			matrix.moveBox(nextLeftPos, direction, dry)
		}
		if nextRightValue == '[' {
			matrix.moveBox(nextRightPos, direction, dry)
		}
	case '>':
		if nextRightValue == '[' && !matrix.moveBox(nextRightPos, direction, dry) {
			return false
		}
	case '<':
		if nextLeftValue == ']' && !matrix.moveBox(nextLeftPos, direction, dry) {
			return false
		}
	}

	if !dry {
		(*matrix)[left.y][left.x] = '.'
		(*matrix)[right.y][right.x] = '.'

		(*matrix)[nextLeftPos.y][nextLeftPos.x] = '['
		(*matrix)[nextRightPos.y][nextRightPos.x] = ']'
	}

	return true
}

func (matrix Matrix) countCoordinates() int {
	total := 0
	for i, row := range matrix {
		for j, ch := range row {
			if ch == '[' {
				total += 100*i + j
			}
		}
	}
	return total
}

func readFromFile(filename string) (Matrix, []rune) {
	data, _ := os.ReadFile(filename)
	parts := strings.Split(string(data), "\n\n")

	matrixLines := strings.Split(parts[0], "\n")
	matrix := Matrix{}
	for y, line := range matrixLines {
		row := []rune{}
		for x, ch := range line {
			switch ch {
			case '#':
				row = append(row, '#', '#')
			case 'O':
				row = append(row, '[', ']')
			case '.':
				row = append(row, '.', '.')
			case '@':
				row = append(row, '@', '.')
				dog = Point{x * 2, y}
			}
		}
		matrix = append(matrix, row)
	}

	moves := parts[1]
	moves = strings.ReplaceAll(moves, "\n", "")

	return matrix, []rune(moves)
}

func main() {
	matrix, moves := readFromFile("input.txt")
	for _, direction := range moves {
		matrix.move(direction)
	}

	fmt.Println(matrix.countCoordinates())
}
