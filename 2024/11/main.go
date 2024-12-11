package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const TTL = 75

type CacheKey struct {
	value int
	ttl   int
}
type Cache map[CacheKey]int

var GlobalCache Cache = Cache{}

func count(value int, ttl int) int {
	if ttl == 0 {
		return 1
	}
	cached, ok := GlobalCache[CacheKey{value, ttl}]
	if ok {
		return cached
	}

	var res int
	if value == 0 {
		res = count(1, ttl-1)
	} else if (int(math.Log10(float64(value)))+1)%2 == 0 {
		valueStr := strconv.Itoa(value)
		separator := len(valueStr) / 2
		first, _ := strconv.Atoi(valueStr[:separator])
		second, _ := strconv.Atoi(valueStr[separator:])

		res = count(first, ttl-1) + count(second, ttl-1)
	} else {
		res = count(value*2024, ttl-1)
	}

	GlobalCache[CacheKey{value, ttl}] = res
	return res
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	total := 0
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		total += count(num, TTL)
	}

	fmt.Println(total)
}
