package main

import (
	"bufio"
	"fmt"
	"os"
)

const Filename = "input.txt"
const BytesToRead = 1024

var Directions = []Point{
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}
var width int
var height int
var start Point
var end Point

type Point struct {
	x, y int
}
type Queue []Point
type Matrix [][]rune
type VisitedMatrix [][]bool
type ParentsMatrix [][]Point
type Maze struct {
	matrix  Matrix
	queue   Queue
	visited VisitedMatrix
	parents ParentsMatrix
}

func (queue *Queue) push(el Point) {
	*queue = append(*queue, el)
}

func (queue *Queue) pop() Point {
	last := (*queue)[0]
	*queue = (*queue)[1:]
	return last
}

func readFromFile() (matrix Matrix) {
	file, _ := os.Open(Filename)
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		row := []rune{}
		for x, ch := range line {
			if ch == 'S' {
				start = Point{x, y}
			}
			if ch == 'E' {
				end = Point{x, y}
			}
			row = append(row, ch)
		}
		matrix = append(matrix, row)
	}

	height = len(matrix)
	width = len(matrix[0])

	return
}

func bfs(maze *Maze, cheatPoint Point) bool {
	root := start
	maze.visited[root.y][root.x] = true
	maze.queue.push(root)

	for len(maze.queue) > 0 {
		value := maze.queue.pop()

		if value == end {
			return true
		}

		for _, direction := range Directions {
			edge := Point{value.x + direction.x, value.y + direction.y}

			if edge.x < 0 || edge.y < 0 || edge.x > width-1 || edge.y > height-1 {
				continue
			}

			if maze.matrix[edge.y][edge.x] == '#' && edge != cheatPoint {
				continue
			}

			if !maze.visited[edge.y][edge.x] {
				maze.visited[edge.y][edge.x] = true
				maze.parents[edge.y][edge.x] = value
				maze.queue.push(edge)
			}
		}
	}

	return false
}

func countSteps(parents *ParentsMatrix) (total int) {
	current := end

	for current != start {
		total++
		current = (*parents)[current.y][current.x]
	}

	return total
}

func initializeMatrix[T any]() *[][]T {
	visitedMatrix := make([][]T, height)
	for i := range visitedMatrix {
		visitedMatrix[i] = make([]T, width)
	}

	return &visitedMatrix
}

func (matrix Matrix) print() {
	for _, row := range matrix {
		for _, el := range row {
			fmt.Print(string(el), " ")
		}
		fmt.Println()
	}
}

func (matrix ParentsMatrix) print() {
	for _, row := range matrix {
		for _, el := range row {
			if (el == Point{}) {
				fmt.Print(".", " ")
			} else {
				fmt.Print("o", " ")
			}

		}
		fmt.Println()
	}
}

func main() {
	matrix := readFromFile()
	visitedMatrix := *initializeMatrix[bool]()
	parentsMatrix := *initializeMatrix[Point]()

	maze := Maze{
		matrix,
		Queue{},
		visitedMatrix,
		parentsMatrix,
	}

	bfs(&maze, Point{0, 0})
	base := countSteps(&maze.parents)

	count := 0
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if matrix[y][x] == '#' {

				maze.visited = *initializeMatrix[bool]()
				maze.parents = *initializeMatrix[Point]()
				maze.queue = Queue{}
				bfs(&maze, Point{x, y})
				delta := base - countSteps(&maze.parents)

				if delta >= 100 {
					count++
				}
			}
		}
	}

	fmt.Println(count)

}
