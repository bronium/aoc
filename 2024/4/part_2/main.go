package main

import (
	"bufio"
	"fmt"
	"os"
)

func isMas(b1, b2, b3 byte) bool {
	word := string(b1) + string(b2) + string(b3)
	return word == "MAS" || word == "SAM"
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	count := 0
	for i, row := range data {
		for j := range row {
			if !(i-1 >= 0 && i+1 < len(data) && j-1 >= 0 && j+1 < len(row)) {
				continue
			}

			if isMas(data[i-1][j-1], data[i][j], data[i+1][j+1]) && isMas(data[i-1][j+1], data[i][j], data[i+1][j-1]) {
				count++
			}
		}
	}

	fmt.Println(count)
}
