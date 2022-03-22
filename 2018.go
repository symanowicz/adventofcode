package main

import (
	"strconv"
	"strings"
)

func Y2018_01(input string) (int, int) {
	x, y := 0, 0
	found := false
	freqs := make(map[int]int)
	for _, n := range strings.Fields(input) {
		i, _ := strconv.Atoi(n)
		x += i
	}
	for !found {
		for _, n := range strings.Fields(input) {
			i, _ := strconv.Atoi(n)
			y += i
			val, pres := freqs[y]
			if pres {
				y = val
				found = true
				break
			} else {
				freqs[y] = y
			}
		}
	}
	return x, y
}