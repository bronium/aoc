package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	paths := map[string]map[string]int{}

	addPath := func(src, dest string, dist int) {
		if paths[src] == nil {
			paths[src] = map[string]int{}
		}
		if paths[dest] == nil {
			paths[dest] = map[string]int{}
		}
		paths[src][dest] = dist
		paths[dest][src] = dist
	}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		places := strings.Split(parts[0], " to ")
		dist, _ := strconv.Atoi(parts[1])
		from, to := places[0], places[1]

		addPath(from, to, dist)
	}

	cities := []string{}
	for k := range paths {
		cities = append(cities, k)
	}

	countDistance := func(arr []string) int {
		total := 0
		for i := 0; i < len(arr)-1; i++ {
			total += paths[arr[i]][arr[i+1]]
		}
		return total
	}

	longest := 0
	var generate func(k int, arr []string)
	generate = func(k int, arr []string) {
		if k == 1 {
			longest = max(longest, countDistance(arr))
		} else {
			for i := 0; i < k; i++ {
				generate(k-1, arr)
				if k%2 == 0 {
					swap := arr[i]
					arr[i] = arr[k-1]
					arr[k-1] = swap
				} else {
					swap := arr[0]
					arr[0] = arr[k-1]
					arr[k-1] = swap
				}
			}
		}
	}

	generate(len(cities), cities)

	fmt.Println(longest)
}
