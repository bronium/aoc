package main

import (
	"fmt"
	"os"
)

func findFloor() int {
	dat, _ := os.ReadFile("input.txt")
	input := string(dat)

	res := 0

	for ind, sym := range input {
		if sym == '(' {
			res++
		}
		if sym == ')' {
			if res == 0 {
				return ind + 1
			}
			res--
		}
	}

	return 0
}

func main() {
	fmt.Println(findFloor())
}
