package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	data  string
	index int
}

func (data Data) look() (byte, bool) {
	index := data.index
	if index == len(data.data) {
		return 0, true
	}
	return data.data[index], false
}

func (data *Data) consumeWord() string {
	acc := ""
	for {
		sym, _ := data.look()
		if !(sym >= 'a' && sym <= 'z') {
			break
		}
		acc += string(sym)
		data.index++
	}
	return acc
}

func (data *Data) consumeNumber() string {
	acc := ""
	for {
		sym, _ := data.look()
		if !(sym == '-' || (sym >= '0' && sym <= '9')) {
			break
		}
		acc += string(sym)
		data.index++
	}
	return acc
}

func (data *Data) count() int {
	brackets := []string{}
	objectTotals := []int{0}
	isObject := func() bool {
		return len(brackets) > 0 && brackets[len(brackets)-1] == "{"
	}

	invalidObject := false
	nestedBrackets := []string{}
	for {
		token, end := data.consumeToken()
		if end {
			break
		}

		if invalidObject {
			if token == "}" && len(nestedBrackets) == 0 {
				invalidObject = false
			} else {
				if token == "{" {
					nestedBrackets = append(nestedBrackets, token)
				}
				if token == "}" {
					if nestedBrackets[len(nestedBrackets)-1] == "{" {
						nestedBrackets = nestedBrackets[:len(nestedBrackets)-1]
					} else {
						panic("WRONG BRACKET")
					}

				}
			}
			continue
		}

		if strings.ContainsAny(token, "[{") {
			brackets = append(brackets, token)
			if token == "{" {
				objectTotals = append(objectTotals, 0)
			}
		}
		if strings.ContainsAny(token, "]}") {
			if token == "}" {
				lastTotal := objectTotals[len(objectTotals)-1]
				objectTotals = objectTotals[:len(objectTotals)-1]
				objectTotals[len(objectTotals)-1] += lastTotal
			}

			brackets = brackets[:len(brackets)-1]
		}
		if token == "red" && isObject() {
			invalidObject = true
			nestedBrackets = []string{}
			brackets = brackets[:len(brackets)-1]
			objectTotals = objectTotals[:len(objectTotals)-1]

			continue
		}

		num, err := strconv.Atoi(token)
		if err == nil {
			objectTotals[len(objectTotals)-1] += num
		}
	}

	total := 0
	for _, value := range objectTotals {
		total += value
	}
	return total
}

func (data *Data) consumeToken() (string, bool) {
	sym, end := data.look()
	if end {
		return "", true
	}

	var token string
	switch {
	case sym >= 'a' && sym <= 'z':
		token = data.consumeWord()
	case sym == '-' || (sym >= '0' && sym <= '9'):
		token = data.consumeNumber()
	default:
		token = string(sym)
		data.index++
	}

	return token, false
}

func main() {
	input, _ := os.ReadFile("input.txt")
	data := Data{string(input), 0}

	fmt.Println("Total: ", data.count())
}
