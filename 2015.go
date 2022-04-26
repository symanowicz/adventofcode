package main

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Y2015_01(input string) (interface{}, interface{}) {
	stop := strings.Count(input, "(") - strings.Count(input, ")")
	basement := 0
	floor := 0
	for i, n := range input {
		if n == rune('(') {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			basement = i + 1
			break
		}
	}
	return stop, basement
}
func Y2015_02(input string) (interface{}, interface{}) {
	packages := make([][]int, 0)
	for _, n := range strings.Fields(input) {
		p := make([]int, 3)
		for i, j := range strings.Split(n, "x") {
			if k, err := strconv.Atoi(j); err == nil {
				p[i] = k
			}
		}
		sort.Ints(p)
		packages = append(packages, p)
	}
	paper, ribbon := 0, 0
	for _, n := range packages {
		paper += (n[0]*n[1]+n[1]*n[2]+n[2]*n[0])*2 + n[0]*n[1]
		ribbon += (n[0]+n[1])*2 + n[0]*n[1]*n[2]
	}
	return paper, ribbon
}
func Y2015_03(input string) (interface{}, interface{}) {
	homes, next := make(map[string]int), make(map[string]int)
	x, y, nx, ny, rx, ry := 0, 0, 0, 0, 0, 0
	homes[strconv.Itoa(x)+" "+strconv.Itoa(y)]++
	next[strconv.Itoa(nx)+" "+strconv.Itoa(ny)]++
	next[strconv.Itoa(rx)+" "+strconv.Itoa(ry)]++
	for i, n := range input {
		switch n {
		case rune('^'):
			if i%2 == 0 {
				nx++
			} else {
				rx++
			}
			x++
		case rune('v'):
			if i%2 == 0 {
				nx--
			} else {
				rx--
			}
			x--
		case rune('>'):
			if i%2 == 0 {
				ny++
			} else {
				ry++
			}
			y++
		case rune('<'):
			if i%2 == 0 {
				ny--
			} else {
				ry--
			}
			y--
		}
		if i%2 == 0 {
			next[strconv.Itoa(nx)+" "+strconv.Itoa(ny)]++
		} else {
			next[strconv.Itoa(rx)+" "+strconv.Itoa(ry)]++
		}
		homes[strconv.Itoa(x)+" "+strconv.Itoa(y)]++
	}
	return len(homes), len(next)
}
func Y2015_04(input string) (interface{}, interface{}) {
	i, j := 0, 0
	s, t := "", ""
	for i = 0; !(strings.HasPrefix(s, "00000")); i++ {
		s = fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
	}
	for j = 0; !(strings.HasPrefix(t, "000000")); j++ {
		t = fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(j))))
	}
	return i - 1, j - 1
}
func Y2015_05(input string) (interface{}, interface{}) {
	nice := 0
	nicer := 0
	good_re := regexp.MustCompile(`(?m)(aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz)`)
	bad_re := regexp.MustCompile(`(?m)(ab|cd|pq|xy)`)
	for _, n := range strings.Fields(input) {
		if !(bad_re.Match([]byte(n))) {
			if good_re.Match([]byte(n)) {
				if strings.Count(n, "a")+strings.Count(n, "e")+strings.Count(n, "i")+strings.Count(n, "o")+strings.Count(n, "u") >= 3 {
					nice++
				}
			}
		}
		trigraph_match := false
		digraph_match := false
		for i := range n {
			if i < len(n)-3 {
				for j := range n[i+2:] {
					if j > 0 {
						if n[i:i+2] == n[j+i+1:j+i+3] {
							digraph_match = true
						}
					}
				}
			}
			if i < len(n)-2 {
				if n[i : i+3][0] == n[i : i+3][2] {
					trigraph_match = true
				}
			}
		}
		if trigraph_match && digraph_match {
			nicer++
		}
	}
	return nice, nicer
}
func Y2015_06(input string) (interface{}, interface{}) {
	lights := make([]bool, 1000000)
	brights := make([]int, 1000000)
	for _, n := range strings.Split(input, "\n") {
		dim := make([]int, 0)
		//replace this loop with a regex
		for _, o := range strings.Fields(n) {
			if strings.Contains(o, ",") {
				for _, p := range strings.Split(o, ",") {
					if q, err := strconv.Atoi(p); err == nil {
						dim = append(dim, q)
					}
				}
			}
		}
		switch {
		case strings.Contains(n, "on"):
			for i := dim[1]; i <= dim[3]; i++ {
				for j := dim[0]; j <= dim[2]; j++ {
					lights[j+i*1000] = true
					brights[j+i*1000] += 1
				}
			}
		case strings.Contains(n, "off"):
			for i := dim[1]; i <= dim[3]; i++ {
				for j := dim[0]; j <= dim[2]; j++ {
					lights[j+i*1000] = false
					if brights[j+i*1000] != 0 {
						brights[j+i*1000]--
					}
				}
			}
		case strings.Contains(n, "toggle"):
			for i := dim[1]; i <= dim[3]; i++ {
				for j := dim[0]; j <= dim[2]; j++ {
					lights[j+i*1000] = !lights[j+i*1000]
					brights[j+i*1000] += 2
				}
			}
		}
	}
	lite, brite := 0, 0
	for _, n := range lights {
		if n {
			lite++
		}
	}
	for _, n := range brights {
		brite += n
	}
	return lite, brite
}
func Y2015_07(input string) (interface{}, interface{}) {
	signals := make(map[string]string)
	for _, n := range strings.Split(input, "\n") {
		m := strings.Split(n, " -> ")
		signals[m[1]] = m[0]
	}
	var solve func(string, string) uint16
	solve = func(input, key string) uint16 {
		n := strings.Split(input, " ")
		a := uint16(0)
		switch {
		case strings.Contains(input, "AND"):
			i, e := strconv.ParseUint(n[0], 0, 16)
			if e != nil {
				a = solve(signals[n[0]], n[0]) & solve(signals[n[2]], n[2])
			} else {
				a = uint16(i) & solve(signals[n[2]], n[2])
			}
		case strings.Contains(input, "OR"):
			a = solve(signals[n[0]], n[0]) | solve(signals[n[2]], n[2])
		case strings.Contains(input, "NOT"):
			a = 65535 - solve(signals[n[1]], n[1])
		case strings.Contains(input, "LSHIFT"):
			m, _ := strconv.Atoi(n[2])
			a = solve(signals[n[0]], n[0]) << m
		case strings.Contains(input, "RSHIFT"):
			m, _ := strconv.Atoi(n[2])
			a = solve(signals[n[0]], n[0]) >> m
		default:
			i, e := strconv.ParseUint(input, 0, 16)
			if e != nil {
				a = solve(signals[input], input)
			} else {
				a = uint16(i)
			}
		}
		signals[key] = strconv.Itoa(int(a))
		return a
	}
	a := solve(signals["a"], "a")
	for _, n := range strings.Split(input, "\n") {
		m := strings.Split(n, " -> ")
		signals[m[1]] = m[0]
	}
	signals["b"] = strconv.Itoa(int(a))
	b := solve(signals["a"], "a")
	return int(a), int(b)
}
func Y2015_08(input string) (interface{}, interface{}) {
	code, mem, requote := 0, 0, 0
	for _, n := range strings.Split(input, "\n") {
		code += utf8.RuneCountInString(n)
		s, _ := strconv.Unquote(n)
		mem += utf8.RuneCountInString(s)
		requote += utf8.RuneCountInString(strconv.Quote(n))
	}
	return code-mem, requote-code
}
