package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Map [][]int
type Coord struct {
	x, y int
}
type Path map[Coord]struct{}

var Directions = [4]Coord{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func (matrix Map) walk(coord Coord, path Path, summitsVisited Path) (int, int) {
	currentHeight := matrix[coord.y][coord.x]

	if currentHeight == 9 {
		_, visited := summitsVisited[coord]

		if visited {
			return 1, 0
		}

		summitsVisited[coord] = struct{}{}
		return 1, 1

	}
	path[coord] = struct{}{}

	trailheads := 0
	ratings := 0
	for _, delta := range Directions {
		next := Coord{coord.x + delta.x, coord.y + delta.y}

		if next.x < len(matrix[0]) && next.x >= 0 && next.y < len(matrix) && next.y >= 0 {
			_, beenThere := path[next]
			heightDiff := matrix[next.y][next.x] - currentHeight
			if !beenThere && heightDiff == 1 {
				trailhead, rating := matrix.walk(next, path, summitsVisited)
				trailheads += trailhead
				ratings += rating
			}
		}
	}

	delete(path, coord)
	return trailheads, ratings
}

func readFromFile(filename string) (Map, []Coord) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	matrix := Map{}
	starts := []Coord{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		nums := []int{}
		for j, ch := range line {
			num, _ := strconv.Atoi(string(ch))
			nums = append(nums, num)
			if num == 0 {
				starts = append(starts, Coord{j, i})
			}
		}
		matrix = append(matrix, nums)
	}

	return matrix, starts
}

func main() {
	matrix, starts := readFromFile("input.txt")

	totalTrailheads, totalRatings := 0, 0
	for _, start := range starts {
		path := Path{}
		summitsVisited := Path{}
		trailheads, ratings := matrix.walk(start, path, summitsVisited)
		totalTrailheads += trailheads
		totalRatings += ratings
	}

	fmt.Println("Total trailheads: ", totalTrailheads, "\nTotal ratings: ", totalRatings)
}
