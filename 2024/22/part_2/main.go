package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Sequence [4]int

var sequences = map[Sequence]int{}
var historyMap = map[Sequence]bool{}

func mix(secret, value int) int {
	return secret ^ value
}

func prune(secret int) int {
	return secret % 16777216
}

func generate(secret int) int {
	secret = prune(mix(secret, secret*64))
	secret = prune(mix(secret, secret/32))
	return prune(mix(secret, secret*2048))
}

func clearHistoryMap() {
	for k := range historyMap {
		delete(historyMap, k)
	}
}

func generateN(secret, n int) int {
	last := secret % 10
	clearHistoryMap()
	diffArr := [4]int{}

	for i := range n {
		secret = generate(secret)
		current := secret % 10
		diffArr[0], diffArr[1], diffArr[2], diffArr[3] = diffArr[1], diffArr[2], diffArr[3], current-last
		last = current
		if i > 2 {
			if !historyMap[diffArr] {
				sequences[diffArr] += current
				historyMap[diffArr] = true
			}
		}
	}
	return secret
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		generateN(num, 2000)
	}

	max := 0
	for _, v := range sequences {
		if v > max {
			max = v
		}
	}

	fmt.Println(max)
}
