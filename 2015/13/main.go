package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseRow(row string) (source string, amount int, target string) {
	words := strings.Split(row, " ")

	var coef int
	switch words[2] {
	case "gain":
		coef = 1
	case "lose":
		coef = -1
	default:
		panic("This should only be gain or lose")
	}

	source = words[0]
	amount, _ = strconv.Atoi(words[3])
	amount *= coef
	target, _ = strings.CutSuffix(words[len(words)-1], ".")
	return
}

func main() {

	relations := map[string]map[string]int{}
	addRelation := func(source, target string, value int) {
		if relations[source] == nil {
			relations[source] = map[string]int{}
		}
		relations[source][target] = value
	}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		source, amount, target := parseRow(line)
		addRelation(source, target, amount)
	}

	guests := []string{}
	for key := range relations {
		guests = append(guests, key)
	}
	guests = append(guests, "Me :)")

	countHappiness := func(guests []string) int {
		total := 0

		var guest1, guest2 string

		for i := 0; i < len(guests)-1; i++ {
			guest1 = guests[i]
			guest2 = guests[i+1]
			total += relations[guest1][guest2] + relations[guest2][guest1]
		}
		guest1 = guests[0]
		guest2 = guests[len(guests)-1]
		total += relations[guest1][guest2] + relations[guest2][guest1]

		return total
	}

	maxHappiness := 0
	var generate func(k int, arr []string)
	generate = func(k int, arr []string) {
		if k == 1 {
			current := countHappiness(arr)
			maxHappiness = max(maxHappiness, current)
		} else {
			for i := 0; i < k; i++ {
				generate(k-1, arr)
				if k%2 == 0 {
					swap := arr[i]
					arr[i] = arr[k-1]
					arr[k-1] = swap
				} else {
					swap := arr[0]
					arr[0] = arr[k-1]
					arr[k-1] = swap
				}
			}
		}
	}

	generate(len(guests), guests)

	fmt.Println(maxHappiness)

}
