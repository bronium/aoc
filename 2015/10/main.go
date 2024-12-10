package main

import (
	"fmt"
	"strconv"
	"strings"
)

func generate(str string) string {
	var nextStr strings.Builder
	current := str[0]
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] == current {
			count++
		} else {
			nextStr.WriteString(strconv.Itoa(count))
			nextStr.WriteRune(rune(current))
			count = 1
			current = str[i]
		}
	}
	nextStr.WriteString(strconv.Itoa(count))
	nextStr.WriteRune(rune(current))
	return nextStr.String()
}

func main() {
	current := "3113322113"

	for i := 0; i < 50; i++ {
		current = generate(current)
	}

	fmt.Println(len(current))
}
