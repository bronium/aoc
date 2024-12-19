package main

import (
	"fmt"
	"os"
	"strings"
)

const Filename = "input.txt"

var patterns = []string{}
var cache = map[string]int{}

func readFromFile() ([]string, []string) {
	data, _ := os.ReadFile(Filename)
	parts := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	patterns := strings.Split(parts[0], ", ")
	combinations := strings.Split(parts[1], "\n")

	return patterns, combinations
}

func isPossible(target string) int {
	if target == "" {
		return 1
	}

	if value, ok := cache[target]; ok {
		return value
	}

	count := 0
	for _, pattern := range patterns {
		after, found := strings.CutPrefix(target, pattern)
		if found {
			count += isPossible(after)
		}
	}

	cache[target] = count
	return count
}

func main() {
	patternArr, combinations := readFromFile()
	patterns = patternArr

	part1_total := 0
	total := 0
	for _, combination := range combinations {
		current := isPossible(combination)
		total += current
		if current > 0 {
			part1_total++
		}
	}

	fmt.Println("First part: ", part1_total)
	fmt.Println("Second part:", total)
}
