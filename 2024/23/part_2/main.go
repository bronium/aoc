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
type PortGroups [][]string

var portMap = PortMap{}
var portGroups = PortGroups{}

func (groups *PortGroups) findGroup(el string) {
	for i, group := range *groups {
		applies := true
		for _, port := range group {
			if !slices.Contains(portMap[port], el) {
				applies = false
				break
			}
		}

		if applies {
			(*groups)[i] = append(group, el)
		}
	}
}

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
	portMap = readFromFile()

	for key, ports := range portMap {
		portGroups = append(portGroups, []string{key})
		for _, port := range ports {
			portGroups.findGroup(port)
		}
		delete(portMap, key)
	}

	maxLen, ind := 0, 0
	for i, group := range portGroups {
		if len(group) > maxLen {
			maxLen = len(group)
			ind = i
		}
	}

	group := portGroups[ind]
	slices.Sort(group)
	fmt.Println(strings.Join(group, ","))
}
