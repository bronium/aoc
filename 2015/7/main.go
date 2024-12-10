package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Wires map[string]string

func parseArg(arg string, wires *Wires) uint16 {
	number, err := strconv.Atoi(arg)

	if err != nil {
		res := parsePrompt((*wires)[arg], arg, wires)
		return res
	}
``
	return uint16(number)
}

func parsePrompt(prompt string, wire string, wires *Wires) uint16 {
	parts := strings.Split(prompt, " ")

	var result uint16
	switch len(parts) {
	case 1:
		return parseArg(parts[0], wires)
	case 2:
		if parts[0] != "NOT" {
			panic("Fatal. Unknown two operand command")
		}
		result = ^parseArg(parts[1], wires)
	case 3:
		arg1 := parseArg(parts[0], wires)
		op := parts[1]
		arg2 := parseArg(parts[2], wires)
		switch op {
		case "AND":
			result = arg1 & arg2
		case "OR":
			result = arg1 | arg2
		case "LSHIFT":
			result = arg1 << arg2
		case "RSHIFT":
			result = arg1 >> arg2
		default:
			panic("Fatal. Unknown three operand command")
		}
	default:
		panic("Fatal. Too many operands")
	}

	(*wires)[wire] = strconv.Itoa(int(result))
	return result
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	wires := Wires{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		wires[parts[1]] = parts[0]
	}

	wire := "a"
	fmt.Println(parsePrompt(wires[wire], wire, &wires))
}
