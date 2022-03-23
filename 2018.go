package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Y2018_01(input string) (interface{}, interface{}) {
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
func Y2018_02(input string) (interface{}, interface{}) {
	double, triple := 0, 0
	for _, n := range strings.Fields(input) {
		counts := make(map[rune]int)
		for _, ch := range "abcdefghijklmnopqrstuvwxyz" {
			counts[ch] = strings.Count(n, fmt.Sprintf("%c", ch))
		}
		for _, v := range counts {
			if v == 2 {
				double++
				break
			}
		}
		for _, v := range counts {
			if v == 3 {
				triple++
				break
			}
		}
	}
	common := ""
	for i, n := range strings.Fields(input) {
		for _, m := range strings.Fields(input)[i+1:] {
			diffs := 0
			for a, b := range n {
				if rune(m[a]) != b {
					diffs++
				}
			}
			if diffs == 1 {
				for k, o := range n {
					if rune(m[k]) != o {
						common = n[:k] + n[k+1:]
					}
				}
			}
		}
	}
	return double * triple, common
}