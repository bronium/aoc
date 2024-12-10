package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	re_mul := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
	re_num := regexp.MustCompile(`\d+`)
	ops := re_mul.FindAllString(string(input), -1)

	total := 0
	enabled := true
	for _, op := range ops {
		if op == "don't()" {
			enabled = false
			continue
		}
		if op == "do()" {
			enabled = true
			continue
		}
		if enabled {
			nums_str := re_num.FindAllString(op, -1)
			nums := []int{}
			for _, v := range nums_str {
				num, _ := strconv.Atoi(v)
				nums = append(nums, num)
			}

			total += nums[0] * nums[1]
		}
	}

	fmt.Println(total)
}
