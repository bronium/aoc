package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validDiff(a, b int, asc bool) bool {
	low, high := b, a
	if asc {
		low, high = a, b
	}

	diff := high - low
	return diff >= 1 && diff <= 3
}

func isSafe(levels []int) bool {
	last := levels[0]
	ASC := levels[1] > last

	for i := 1; i < len(levels); i++ {
		if !validDiff(last, levels[i], ASC) {
			return false
		}

		last = levels[i]
	}
	return true
}

func hasValidPermutation(report string) bool {
	levels_string := strings.Split(report, " ")
	levels := []int{}
	for _, l := range levels_string {
		number, _ := strconv.Atoi(l)
		levels = append(levels, number)
	}

	tests := [][]int{}
	tests = append(tests, levels)
	for i := range levels {
		new := []int{}
		for j, el := range levels {
			if j == i {
				continue
			}
			new = append(new, el)
		}
		tests = append(tests, new)
	}

	for _, report := range tests {
		if isSafe(report) {
			return true
		}
	}
	return false
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if hasValidPermutation(line) {
			total++
		}
	}
	fmt.Println(total)
}
