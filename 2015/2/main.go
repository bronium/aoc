package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	sumArea := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "x")
		l, _ := strconv.Atoi(nums[0])
		w, _ := strconv.Atoi(nums[1])
		h, _ := strconv.Atoi(nums[2])
		ribbon := min(l+w, w+h, h+l) * 2
		bow := l * w * h
		sumArea += ribbon + bow
	}
	fmt.Println(sumArea)

	file.Close()
}
