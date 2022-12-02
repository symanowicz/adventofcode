package main

import (
	"fmt"
	"math"
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
		c          string
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
		start := n.x + n.y*1000
		for i := 0; i < n.h; i++ {
			for j := 0; j < n.w; j++ {
				switch fabric[start+j] {
				case '.':
					fabric[start+j] = 'O'
				case 'O':
					fabric[start+j] = 'X'
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
		start := n.x + n.y*1000
		for i := 0; i < n.h; i++ {
			for j := 0; j < n.w; j++ {
				if fabric[start+j] == 'X' {
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
		start  time.Time
		length time.Duration
	}
	type guard struct {
		asleep  []sleep
		total   time.Duration
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
			g.asleep[len(g.asleep)-1].length = t.Sub(g.asleep[len(g.asleep)-1].start)
			g.total += g.asleep[len(g.asleep)-1].length
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

// this function can be shorted by using a field struct to contain edge, marker, safe
func Y2018_06(input string) (interface{}, interface{}) {
	type point struct {
		x, y int
	}
	p := make([]point, 0)
	for _, n := range strings.FieldsFunc(input, func(c rune) bool { return c == '\n' }) {
		s := strings.Split(n, ", ")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		p = append(p, point{x, y})
	}
	field := make([]rune, 160000)
	marker := "abcdefghijklmnopqrstuvwxyABCDEFGHIJKLMNOPQRSTUVWXY"

	for i := range field {
		x := i % 400
		y := i / 400
		closest := 999999
		for j, m := range p {
			dist := int(math.Abs(float64(x-m.x)) + math.Abs(float64(y-m.y)))
			if dist < closest {
				field[i] = rune(marker[j])
				closest = dist
			} else if dist == closest {
				field[i] = '.'
			}
		}
	}
	m := make(map[rune]int)
	for _, n := range marker {
		m[n] = 0
	}
	for _, n := range field {
		m[n]++
	}
	for i, n := range field {
		x := i % 400
		y := i / 400
		if x == 0 || x == 399 {
			delete(m, n)
		}
		if y == 0 || y == 399 {
			delete(m, n)
		}
	}
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	//part 2
	for i := range field {
		x := i % 400
		y := i / 400
		dist := make([]int, 0)
		sum := 0
		for _, m := range p {
			dist = append(dist, int(math.Abs(float64(x-m.x))+math.Abs(float64(y-m.y))))
		}
		for _, n := range dist {
			sum += n
		}
		if sum < 10000 {
			field[i] = '#'
		}
	}
	safe := 0
	for _, n := range field {
		if n == '#' {
			safe++
		}
	}
	return max, safe
}

// find better way to duplicate the map struct?
func Y2018_07(input string) (interface{}, interface{}) {
	type step struct {
		prereqs             []string
		available, assigned bool
		time                int
	}
	steps := make(map[string]*step)
	steps2 := make(map[string]*step)
	for i, n := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		steps[string(n)] = new(step)
		steps[string(n)].available = true
		steps[string(n)].assigned = false
		steps[string(n)].time = 60 + i
		steps2[string(n)] = new(step)
		steps2[string(n)].available = true
		steps2[string(n)].assigned = false
		steps2[string(n)].time = 60 + i
	}
	r := regexp.MustCompile(`.* ([A-Z]){1} .* ([A-Z]){1} .*`)
	for _, n := range strings.FieldsFunc(input, func(c rune) bool { return c == '\n' }) {
		m := r.FindStringSubmatch(n)
		for k := range steps {
			if m[2] == k {
				steps[k].prereqs = append(steps[k].prereqs, m[1])
				steps[k].available = false
				break
			}
		}
		for k := range steps2 {
			if m[2] == k {
				steps2[k].prereqs = append(steps2[k].prereqs, m[1])
				steps2[k].available = false
				break
			}
		}
	}
	order := ""
	for {
		done := true
		for _, n := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
			if _, p := steps[string(n)]; p {
				if !steps[string(n)].available {
					p_done := true
					for _, m := range steps[string(n)].prereqs {
						if _, p2 := steps[m]; p2 {
							p_done = false
						}
					}
					if p_done {
						steps[string(n)].available = true
					}
					done = false
				}
				if steps[string(n)].available {
					order += string(n)
					delete(steps, string(n))
					break
				}
			}
		}
		if done {
			break
		}
	}
	timer := 0
	order2 := ""
	elfs := make([]string, 5)
	for {
		done := true
		for i := range elfs {
			if elfs[i] != "" {
				if steps2[elfs[i]].time == 0 {
					order2 += elfs[i]
					delete(steps2, elfs[i])
					elfs[i] = ""
				} else {
					steps2[elfs[i]].time--
					done = false
				}
			}
			if elfs[i] == "" {
				for _, n := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
					if _, p := steps2[string(n)]; p {
						if !steps2[string(n)].available {
							p_done := true
							for _, m := range steps2[string(n)].prereqs {
								if _, p2 := steps2[m]; p2 {
									p_done = false
								}
							}
							if p_done {
								steps2[string(n)].available = true
							}
							done = false
						}
						if steps2[string(n)].available && !steps2[string(n)].assigned {
							elfs[i] = string(n)
							steps2[string(n)].assigned = true
							break
						}
					}
				}
			}
		}
		if done {
			break
		}
		timer++
	}
	return order, timer
}
