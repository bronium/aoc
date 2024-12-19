package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Filename = "input.txt"
const BytesToRead = 1024
const Width = 71
const Height = 71

var Start = Point{0, 0}
var End = Point{Width - 1, Height - 1}

type Point struct {
	x, y int
}
type Queue []Point
type Matrix [Height][Width]rune
type VisitedMatrix [Height][Width]bool
type ParentsMatrix [Height][Width]Point
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

func initializeMatrix() Matrix {
	matrix := [Height][Width]rune{}

	for i, row := range matrix {
		for j := range row {
			matrix[i][j] = '.'
		}
	}

	return matrix
}

func readBytes() (bytes []Point) {
	file, _ := os.Open(Filename)
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		parts := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		bytes = append(bytes, Point{x, y})
	}

	return
}

func adjacentEdges(src Point, matrix Matrix) []Point {
	edges := []Point{}

	directions := []Point{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	for _, direction := range directions {
		potentialEdge := Point{src.x + direction.x, src.y + direction.y}

		if potentialEdge.x < 0 || potentialEdge.y < 0 || potentialEdge.x > Width-1 || potentialEdge.y > Height-1 {
			continue
		}

		if matrix[potentialEdge.y][potentialEdge.x] == '#' {
			continue
		}

		edges = append(edges, potentialEdge)
	}

	return edges
}

func bfs(maze *Maze) bool {
	root := Start
	maze.visited[root.y][root.x] = true
	maze.queue.push(root)

	for len(maze.queue) > 0 {
		value := maze.queue.pop()

		if value == End {
			return true
		}

		for _, edge := range adjacentEdges(value, maze.matrix) {
			if !maze.visited[edge.y][edge.x] {
				maze.visited[edge.y][edge.x] = true
				maze.parents[edge.y][edge.x] = value
				maze.queue.push(edge)
			}
		}
	}

	return false
}

func countSteps(parents ParentsMatrix) (total int) {
	current := Point{Width - 1, Height - 1}
	for current != Start {
		total++
		current = parents[current.y][current.x]
	}

	return total
}

func main() {
	bytes := readBytes()

	maze := Maze{
		initializeMatrix(),
		Queue{},
		VisitedMatrix{},
		ParentsMatrix{},
	}

	for i, corruption := range bytes {
		maze.matrix[corruption.y][corruption.x] = '#'

		if i == BytesToRead-1 {
			bfs(&maze)
			fmt.Println("First task: ", countSteps(maze.parents))
		}

		if i > BytesToRead-1 {
			maze.queue = Queue{}
			maze.visited = VisitedMatrix{}
			maze.parents = ParentsMatrix{}
			if !bfs(&maze) {
				first := bytes[i]
				fmt.Printf("Second task: %v,%v\n", first.x, first.y)
				break
			}
		}
	}
}
