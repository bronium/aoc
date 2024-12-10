package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Disk []int

func (disk *Disk) next() (int, int) {
	for i, el := range *disk {
		if el == 0 {
			(*disk)[i] = -1
			return i, 0
		}

		if el > 0 {
			(*disk)[i] = -1
			return i, el
		}
	}

	return -1, -1
}

func (disk *Disk) pop() int {
	for i := len(*disk) - 1; i >= 0; i-- {
		if (*disk)[i] > 0 {
			(*disk)[i]--
			return i
		}
	}
	return -1
}

func main() {
	input, _ := os.ReadFile("input")
	parts := strings.Split(string(input), "")
	blocks := Disk{}
	voids := Disk{}
	for i, el := range parts {
		num, _ := strconv.Atoi(el)
		if i%2 == 0 {
			blocks = append(blocks, num)
		} else {
			voids = append(voids, num)
		}
	}

	checksum := 0
	ind := 0
	for {
		next, nextCount := blocks.next()
		if next == -1 {
			break
		}
		if nextCount == 0 {
			ind++
			continue
		}
		for range nextCount {
			checksum += ind * next
			ind++
		}
		_, voidCount := voids.next()
		for range voidCount {
			popped := blocks.pop()
			if popped == -1 {
				break
			}
			checksum += ind * popped
			ind++
		}
	}

	fmt.Println(checksum)
}
