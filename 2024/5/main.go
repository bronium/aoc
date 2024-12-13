package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	parts := strings.Split(string(data), "\n\n")

	rulesInput := strings.Split(parts[0], "\n")

	rules := map[int][]int{}
	for _, rule := range rulesInput {
		ruleParts := strings.Split(rule, "|")
		ruleBefore, _ := strconv.Atoi(ruleParts[0])
		ruleAfter, _ := strconv.Atoi(ruleParts[1])

		rules[ruleBefore] = append(rules[ruleBefore], ruleAfter)
	}

	updateStrings := strings.Split(parts[1], "\n")
	updates := [][]int{}

	for _, line := range updateStrings {
		arr := strings.Split(line, ",")
		arrInt := []int{}
		for _, el := range arr {
			num, _ := strconv.Atoi(el)
			arrInt = append(arrInt, num)
		}
		updates = append(updates, arrInt)
	}

	checkUpdate := func(update []int) (bool, int) {
		for ind, page := range update {
			afterPages := rules[page]

			for _, before := range afterPages {
				found := slices.Index(update, before)
				if found > -1 && found < ind {
					return false, found
				}
			}

		}
		return true, -1
	}

	sum := 0
	for _, update := range updates {
		correct, ind := checkUpdate(update)
		if !correct {
			for {
				update = append(update, update[ind])
				update = slices.Delete(update, ind, ind+1)

				valid, newInd := checkUpdate(update)
				if valid {
					fmt.Println(update)
					break
				}
				ind = newInd
			}
			mid := len(update) / 2
			sum += update[mid]
		}
	}

	fmt.Println(sum)
}
