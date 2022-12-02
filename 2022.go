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

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	return elves[0], elves[0] + elves[1] + elves[2]
}
func Y2022_02(input string) (interface{}, interface{}) {
	// matrix of outcomes
	//  Part 1 |  Part 2
	//   A B C |   A B C
	// X 4 1 7 | X 3 1 2
	// Y 8 5 2 | Y 4 5 6
	// Z 3 9 6 | Z 8 9 7
	total, total2 := 0, 0
	outcomes := []int{4,1,7,8,5,2,3,9,6}
	outcomes2 := []int{3,1,2,4,5,6,8,9,7}
	for _, n := range strings.FieldsFunc(input, func(c rune) bool { return c == '\n' }) {
		total += outcomes[(n[2]-88)*3+(n[0]-65)]
		total2 += outcomes2[(n[2]-88)*3+(n[0]-65)]
	}
	return total,total2
}