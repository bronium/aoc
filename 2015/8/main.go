package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countDiff(s string) int {
	// parsed, _ := strconv.Unquote(s)
	// return len(s) - len(parsed)

	specialSymbols := strings.Count(s, "\\") + strings.Count(s, "\"")

	return specialSymbols + 2
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		total += countDiff(scanner.Text())
	}
	fmt.Println(total)
}
