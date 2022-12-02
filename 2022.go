package main

import (
	"sort"
	"strconv"
	"strings"
)

func Y2022_01(input string) (interface{}, interface{}) {
	elves := make([]int, 0)
	currentElf := 0
	for _, n := range strings.Split(input, "\n\n") {
		for _, m := range strings.Fields(n) {
			val, _ := strconv.Atoi(m)
			currentElf += val
		}
		elves = append(elves, currentElf)
		currentElf = 0
	}

	sort.Ints(elves)
	return elves[len(elves)-1], elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
}
