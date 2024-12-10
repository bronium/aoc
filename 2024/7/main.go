package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(target int, args []int) bool {
	if len(args) == 1 {
		return target == args[0]
	}

	sum := args[0] + args[1]
	newArgs := append([]int{sum}, args[2:]...)
	if calculate(target, newArgs) {
		return true
	}

	product := args[0] * args[1]
	newArgs = append([]int{product}, args[2:]...)
	if calculate(target, newArgs) {
		return true
	}

	arg1 := strconv.Itoa(args[0])
	arg2 := strconv.Itoa(args[1])
	var builder strings.Builder
	builder.WriteString(arg1)
	builder.WriteString(arg2)
	res, _ := strconv.Atoi(builder.String())
	newArgs = append([]int{res}, args[2:]...)
	if calculate(target, newArgs) {
		return true
	}

	return false
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])

		opsStrings := strings.Split(parts[1], " ")
		ops := []int{}
		for _, el := range opsStrings {
			op, _ := strconv.Atoi(el)
			ops = append(ops, op)
		}

		if calculate(target, ops) {
			count += target
		}
	}

	fmt.Printf("count: %v\n", count)
}
