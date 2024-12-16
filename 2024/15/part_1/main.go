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

func (matrix *Matrix) move(coord Point, direction rune) bool {
	delta := Directions[direction]
	nextPos := Point{coord.x + delta.x, coord.y + delta.y}
	currentValue := (*matrix)[coord.y][coord.x]
	nextValue := (*matrix)[nextPos.y][nextPos.x]

	moveSelf := func() {
		if currentValue == '@' {
			dog = nextPos
		}
		(*matrix)[coord.y][coord.x] = '.'
		(*matrix)[nextPos.y][nextPos.x] = currentValue
	}

	switch nextValue {
	case '#':
		return false
	case '.':
		moveSelf()
		return true
	case 'O':
		if matrix.move(nextPos, direction) {
			moveSelf()
			return true
		}
	}

	return false
}

func (matrix Matrix) countCoordinates() int {
	total := 0
	for i, row := range matrix {
		for j, ch := range row {
			if ch == 'O' {
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
		matrix = append(matrix, []rune(line))
		if x := strings.Index(line, "@"); x > 0 {
			dog = Point{x, y}
		}
	}

	moves := parts[1]
	moves = strings.ReplaceAll(moves, "\n", "")

	return matrix, []rune(moves)
}

func main() {

	matrix, moves := readFromFile("input.txt")
	for _, direction := range moves {
		matrix.move(dog, direction)
	}

	fmt.Println(matrix.countCoordinates())
}
