package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const Filename = "input.txt"

type PortMap map[string][]string

func readFromFile() PortMap {
	file, _ := os.Open(Filename)
	scanner := bufio.NewScanner(file)

	portMap := PortMap{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		port1, port2 := parts[0], parts[1]
		portMap[port1] = append(portMap[port1], port2)
		portMap[port2] = append(portMap[port2], port1)
	}

	return portMap
}

func main() {
	portMap := readFromFile()

	triples := map[string]bool{}
	for key, ports := range portMap {
		if len(ports) > 1 {
			for i := 0; i < len(ports)-1; i++ {
				for j := i + 1; j < len(ports); j++ {
					if slices.Contains(portMap[ports[i]], ports[j]) {
						arr := []string{key, ports[i], ports[j]}
						slices.Sort(arr)
						for _, val := range arr {
							if val[0] == 't' {
								triples[strings.Join(arr, ",")] = true
								continue
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(len(triples))
}
