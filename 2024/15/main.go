package main

import (
	"os"
	"strings"
)

type Delta struct {
	x, y int
}

type Matrix [][]rune

var Directions = map[rune]Delta{
	'^': {0, -1},
	'v': {0, 1},
	'>': {1, 0},
	'<': {-1, 0},
}

func (matrix *Matrix) move(direction rune) {
}

func (matrix *Matrix) moveBox(direction rune) bool {
}

func readFromFile(filename string) (Matrix, []rune) {
	data, _ := os.ReadFile(filename)
	parts := strings.Split(string(data), "\n\n")

	matrixLines := strings.Split(parts[0], "\n")
	matrix := Matrix{}
	for _, line := range matrixLines {
		matrix = append(matrix, []rune(line))
	}

	moves := parts[1]
	moves = strings.ReplaceAll(moves, "\n", "")

	return matrix, []rune(moves)
}

func main() {
	matrix, moves := readFromFile("smol.txt")
	for _, direction := range moves {
		matrix.move(direction)
	}
}
