package main

import (
	"bufio"
	"fmt"
	"os"
)

func isXmas(l1, l2, l3, l4 byte) bool {
	word := string(l1) + string(l2) + string(l3) + string(l4)
	return word == "XMAS" || word == "SAMX"
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
			// Vertical
			if i+3 < len(data) {
				if isXmas(data[i][j], data[i+1][j], data[i+2][j], data[i+3][j]) {
					count++
				}
			}
			// Horizontal
			if j+3 < len(row) {
				if isXmas(data[i][j], data[i][j+1], data[i][j+2], data[i][j+3]) {
					count++
				}
			}
			// Right Diagonal
			if i+3 < len(data) && j+3 < len(row) {
				if isXmas(data[i][j], data[i+1][j+1], data[i+2][j+2], data[i+3][j+3]) {
					count++
				}
			}
			// Left Diagonal
			if i+3 < len(data) && j-3 >= 0 {
				if isXmas(data[i][j], data[i+1][j-1], data[i+2][j-2], data[i+3][j-3]) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
