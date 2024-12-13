package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const Mark = '#'

type Matrix [][]rune
type Point struct {
	x, y int
}
type Frequencies map[rune][]Point

func (matrix *Matrix) mark(point, diff Point) {
	if diff.x == 0 && diff.y == 0 {
		return
	}

	antinode := Point{point.x, point.y}

	for {
		if antinode.x < 0 || antinode.x >= len((*matrix)[0]) || antinode.y < 0 || antinode.y >= len(*matrix) {
			return
		}

		(*matrix)[antinode.x][antinode.y] = Mark

		antinode.x += diff.x
		antinode.y += diff.y
	}
}

func (matrix Matrix) countX() int {
	count := 0
	for _, row := range matrix {
		for _, el := range row {
			if el == Mark {
				count++
			}
		}
	}
	return count
}

func readFromFile(filename string) Matrix {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	matrix := Matrix{}
	for scanner.Scan() {
		line := scanner.Text()

		arr := []rune{}
		for _, sym := range line {
			arr = append(arr, sym)
		}
		matrix = append(matrix, arr)
	}

	return matrix
}

func generateFrequencies(matrix Matrix) Frequencies {
	freqs := Frequencies{}
	for i, row := range matrix {
		for j, item := range row {
			if unicode.In(item, unicode.Number, unicode.Letter) {
				freqs[item] = append(freqs[item], Point{i, j})
			}
		}
	}
	return freqs
}

func main() {
	matrix := readFromFile("input")
	freqs := generateFrequencies(matrix)

	for _, points := range freqs {
		for _, first := range points {
			for _, second := range points {
				diff1 := Point{first.x - second.x, first.y - second.y}
				diff2 := Point{second.x - first.x, second.y - first.y}
				matrix.mark(first, diff1)
				matrix.mark(second, diff2)
			}
		}
	}

	fmt.Println(matrix.countX())
}
