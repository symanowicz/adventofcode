package main

import (
	"fmt"
	"regexp"
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
func Y2018_03(input string) (interface{}, interface{}) {
	type order struct {
		c string
		x, y, w, h int
	}

	r := regexp.MustCompile(`.*@ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)`)
	orders := make([]order, 0)
	for _, n := range strings.FieldsFunc(input, func(c rune) bool { return c == '\n' }) {
		var o order
		m := r.FindStringSubmatch(n)
		o.c = m[0]
		o.x, _ = strconv.Atoi(m[1])
		o.y, _ = strconv.Atoi(m[2])
		o.w, _ = strconv.Atoi(m[3])
		o.h, _ = strconv.Atoi(m[4])
		orders = append(orders, o)
	}
	var fabric [1000000]rune
	for i := range fabric {
		fabric[i] = '.'
	}
	for _, n := range orders {
		start := n.x + n.y * 1000
		for i := 0; i < n.h; i++ {
			for j := 0; j < n.w; j++ {
				switch fabric[start + j] {
				case '.':
					fabric[start + j] = 'O'
				case 'O':
					fabric[start + j] = 'X'
				}
			}
			start += 1000
		}
	}
	total := 0
	for i := range fabric {
		if fabric[i] == 'X' {
			total++
		}
	}
	claim := ""
	for _, n := range orders {
		intact := true
		start := n.x + n.y * 1000
		for i := 0; i < n.h; i++ {
			for j := 0; j < n.w; j++ {
				if fabric[start + j] == 'X' {
					intact = false
				}
			}
			start += 1000
		}
		if intact {
			claim = n.c
			break
		}
	}
	return total, claim
}