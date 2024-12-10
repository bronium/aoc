package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Disk []int

func (disk *Disk) fillFromFile(filename string) {
	input, _ := os.ReadFile(filename)
	parts := strings.Split(string(input), "")

	ind := 0
	for i, el := range parts {
		num, _ := strconv.Atoi(el)

		value := -1
		if i%2 == 0 {
			value = i / 2
		}

		for range num {
			(*disk)[ind] = value
			ind++
		}
	}
}

func (disk Disk) checksum() int {
	checksum := 0
	for i, el := range disk {
		if el > 0 {
			checksum += i * el
		}
	}
	return checksum
}

func main() {
	blocks := make(Disk, 100_000)
	blocks.fillFromFile("../input")

	touched := map[int]bool{}
	buff := []int{}
	for i := len(blocks) - 1; i >= 0; i-- {
		curr := blocks[i]

		if len(buff) > 0 && curr != buff[0] {
			voidLength := 0
			voidStartIndex := 0
			for j := 0; j <= i+1; j++ {
				if blocks[j] == -1 {
					if voidLength == 0 {
						voidStartIndex = j
					}
					voidLength++
				} else {
					if voidLength >= len(buff) {
						for k := voidStartIndex; k < voidStartIndex+len(buff); k++ {
							blocks[k] = buff[0]
						}
						for k := i + 1; k < i+1+len(buff); k++ {
							blocks[k] = -1
						}

						break
					} else {
						voidLength = 0
						voidStartIndex = 0
					}
				}
			}

			touched[buff[0]] = true
			buff = nil
		}

		if (len(buff) == 0 || buff[0] == curr) && curr > 0 && !touched[curr] {
			buff = append(buff, curr)
		}

	}

	fmt.Println(blocks.checksum())
}
