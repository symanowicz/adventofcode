package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
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
			claim = strings.Fields(n.c)[0][1:]
			break
		}
	}
	return total, claim
}
func Y2018_04(input string) (interface{}, interface{}) {
	type sleep struct {
		start time.Time
		length time.Duration
	}
	type guard struct {
		asleep []sleep
		total time.Duration
		minutes map[int]int
	}
	//Removes '#' to make id parsing easier, splits input on newline
	sorted := strings.FieldsFunc(strings.ReplaceAll(input, "#", ""), func(c rune) bool { return c == '\n' })
	//Sorts by timestamp (easy no?)
	sort.Strings(sorted)
	r := regexp.MustCompile(`\[(.*)\] [[:alpha:]]{5} ([[:alnum:]]+).*`)
	const timeForm = "2006-01-02 15:04"
	guards := make(map[int]guard)
	current_id := 0
	for _, n := range sorted {
		m := r.FindStringSubmatch(n)
		switch m[2] {
		//push new sleep period for current guard
		case "asleep":
			t, _ := time.Parse(timeForm, m[1])
			g := guards[current_id]
			g.asleep = append(g.asleep, sleep{t, 0})
			guards[current_id] = g
		//set length of current sleep period for current guard
		case "up":
			t, _ := time.Parse(timeForm, m[1])
			g := guards[current_id]
			g.asleep[len(g.asleep) - 1].length = t.Sub(g.asleep[len(g.asleep) - 1].start)
			g.total += g.asleep[len(g.asleep) - 1].length
			guards[current_id] = g
		//guard change, check list if guard exists, otherwise create guard and push to list
		default:
			i, _ := strconv.Atoi(m[2])
			_, prs := guards[i]
			if prs {
				current_id = i
			} else {
				guards[i] = guard{make([]sleep, 0), 0, make(map[int]int)}
				current_id = i
			}
		}
	}
	//count minute occurences
	for _, v := range guards {
		for _, s := range v.asleep {
			for i := s.start.Minute(); i < s.start.Add(s.length).Minute(); i++ {
				v.minutes[i]++
			}
		}
	}
	//who was asleep the longest?
	sleeper, longest := 0, 0
	for k, v := range guards {
		if len(v.asleep) != 0 {
			if int(v.total) > longest {
				longest = int(v.total)
				sleeper = k
			}
		}
	}
	//most likely minute to be sleeping
	minute, max := 0, 0
	for k, v := range guards[sleeper].minutes {
		if v > max {
			max = v
			minute = k
		}
	}
	//part 2
	sleeper2, minute2, max2 := 0, 0, 0
	for k, v := range guards {
		if len(v.asleep) != 0 {
			for l, w := range v.minutes {
				if w > max2 {
					sleeper2 = k
					minute2 = l
					max2 = w
				}
			}
		}
	}
	return sleeper * minute, sleeper2 * minute2
}
func Y2018_05(input string) (interface{}, interface{}) {
	length := 0
	polymer := input
	for {
		length = len(polymer)
		for _, n := range "abcdefghijklmnopqrstuvwxyz" {
			polymer = strings.Replace(polymer, string(n)+strings.ToUpper(string(n)), "", -1)
			polymer = strings.Replace(polymer, strings.ToUpper(string(n))+string(n), "", -1)
		}
		if length == len(polymer) {
			break
		}
	}
	totals := make([]int, 0)
	for _, m := range "abcdefghijklmnopqrstuvwxyz" {
		improved := input
		improved = strings.Replace(improved, string(m), "", -1)
		improved = strings.Replace(improved, strings.ToUpper(string(m)), "", -1)
		for {
			length = len(improved)
			for _, n := range "abcdefghijklmnopqrstuvwxyz" {
				improved = strings.Replace(improved, string(n)+strings.ToUpper(string(n)), "", -1)
				improved = strings.Replace(improved, strings.ToUpper(string(n))+string(n), "", -1)
			}
			if length == len(improved) {
				totals = append(totals, length)
				break
			}
		}
	}
	min := 9999
	for _, n := range totals {
		if n < min {
			min = n
		}
	}
	return len(polymer), min
}