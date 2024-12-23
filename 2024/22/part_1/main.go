package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func generateN(secret, n int) int {
	last := secret % 10
	diffArr := []int{}
	for range n {
		secret = generate(secret)
		current := secret % 10
		diffArr = append(diffArr, current-last)
		last = current

	}
	return secret
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		total += generateN(num, 2000)
	}
	fmt.Println(total)
}
