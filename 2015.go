package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Y2015_01(input string) (interface{}, interface{}) {
	stop := strings.Count(input, "(") - strings.Count(input, ")")
	basement, floor := 0, 0
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
	return code - mem, requote - code
}
func Y2015_09(input string) (interface{}, interface{}) {
	type path struct {
		from, to string
		distance int
	}
	permutations := func(arr []int) [][]int {
		var helper func([]int, int)
		res := [][]int{}

		helper = func(arr []int, n int) {
			if n == 1 {
				tmp := make([]int, len(arr))
				copy(tmp, arr)
				res = append(res, tmp)
			} else {
				for i := 0; i < n; i++ {
					helper(arr, n-1)
					if n%2 == 1 {
						tmp := arr[i]
						arr[i] = arr[n-1]
						arr[n-1] = tmp
					} else {
						tmp := arr[0]
						arr[0] = arr[n-1]
						arr[n-1] = tmp
					}
				}
			}
		}
		helper(arr, len(arr))
		return res
	}
	paths := make([]path, 0)
	locs := ""
	for _, n := range strings.Split(input, "\n") {
		m := strings.Split(n, " = ")
		o := strings.Split(m[0], " to ")
		p, _ := strconv.Atoi(m[1])
		paths = append(paths, path{o[0], o[1], p})
		if !strings.Contains(locs, o[0]) {
			locs += o[0] + ","
		}
		if !strings.Contains(locs, o[1]) {
			locs += o[1] + ","
		}
	}
	slocs := strings.Split(locs[:len(locs)-1], ",")
	distance := []int{}
	for _, n := range permutations([]int{0, 1, 2, 3, 4, 5, 6, 7}) {
		acc := 0
		for j := 0; j < len(n)-1; j++ {
			for _, m := range paths {
				if (slocs[n[j]] == m.from && slocs[n[j+1]] == m.to) || (slocs[n[j]] == m.to && slocs[n[j+1]] == m.from) {
					acc += m.distance
					break
				}
			}
		}
		distance = append(distance, acc)
	}
	sort.Ints(distance)
	return distance[0], distance[len(distance)-1]
}
func Y2015_10(input string) (interface{}, interface{}) {
	lss := func(s string) (r string) {
		c := s[0]
		nc := 1
		for i := 1; i < len(s); i++ {
			d := s[i]
			if d == c {
				nc++
				continue
			}
			r += strconv.Itoa(nc) + string(c)
			c = d
			nc = 1
		}
		return r + strconv.Itoa(nc) + string(c)
	}
	forty := 0
	//this takes a long time (~30 mins), problem text implies that there's a way to split the problem space
	for i := 0; i < 50; i++ {
		if i == 40 {
			forty = utf8.RuneCountInString(input)
		}
		input = lss(input)
	}
	return forty, utf8.RuneCountInString(input)
}
func Y2015_11(input string) (interface{}, interface{}) {
	pass := []rune(input)
	reverse := func(s []rune) {
		for i := 0; i < 4; i++ {
			s[i], s[7-i] = s[7-i], s[i]
		}
	}
	var increment func([]rune)
	increment = func(s []rune) {
		for i, n := range s {
			if n == 'z' {
				s[i] = 'a'
				increment(s[1:])
				break
			} else {
				s[i]++
				break
			}
		}
	}
	double := func(s []rune) bool {
		found := false
		for i := 0; i < len(s)-1; i++ {
			a, b := i, i+1
			if s[a] == s[b] {
				for j := 0; j < len(s)-1; j++ {
					c, d := j, j+1
					if s[c] == s[d] && !(c == a || c == b || d == a || d == b) {
						found = true
						break
					}
				}
				if found {
					break
				}
			}
		}
		return found
	}
	straight := func(s []rune) bool {
		found := false
		for i := 0; i < len(s)-2; i++ {
			if s[i]+1 == s[i+1] && s[i+1]+1 == s[i+2] {
				found = true
			}
		}
		return found
	}
	first, firstpass := true, ""
	for {
		reverse(pass)
		increment(pass)
		reverse(pass)
		if strings.ContainsAny(string(pass), "iol") || !straight(pass) || !double(pass) {
			continue
		}
		if first {
			firstpass = string(pass)
			first = false
			continue
		}
		break
	}
	return firstpass, string(pass)
}
func Y2015_12(input string) (interface{}, interface{}) {
	sum := 0
	nored := false
	var parse func(interface{})
	parse = func(data interface{}) {
		switch data := data.(type) {
		case []interface{}:
			for _, n := range data {
				parse(n)
			}
		case map[string]interface{}:
			safe := true
			if nored {
				for k := range data {
					switch data[k].(type) {
					case string:
						if data[k].(string) == "red" {
							safe = false
						}
					}
				}
			}
			if safe {
				for _, v := range data {
					parse(v)
				}
			}
		case float64:
			sum += int(data)
		}
	}
	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		panic(err)
	}
	parse(data)
	oldsum := sum
	sum = 0
	nored = true
	parse(data)
	return oldsum, sum
}
func Y2015_13(input string) (interface{}, interface{}) {
	permutations := func(arr []int) [][]int {
		var helper func([]int, int)
		res := [][]int{}

		helper = func(arr []int, n int) {
			if n == 1 {
				tmp := make([]int, len(arr))
				copy(tmp, arr)
				res = append(res, tmp)
			} else {
				for i := 0; i < n; i++ {
					helper(arr, n-1)
					if n%2 == 1 {
						tmp := arr[i]
						arr[i] = arr[n-1]
						arr[n-1] = tmp
					} else {
						tmp := arr[0]
						arr[0] = arr[n-1]
						arr[n-1] = tmp
					}
				}
			}
		}
		helper(arr, len(arr))
		return res
	}
	feels := make(map[string]map[string]int, 0)
	names := []string{}
	for _, n := range strings.Split(input, "\n") {
		s := strings.Split(strings.ReplaceAll(n, ".", ""), " ")
		strength, _ := strconv.Atoi(s[3])
		if s[2] == "lose" {
			strength *= -1
		}
		if feels[s[0]] == nil {
			feels[s[0]] = make(map[string]int, 0)
			names = append(names, s[0])
		}
		feels[s[0]][s[10]] = strength
	}
	relations := make(map[string]int, 0)
	for k, v := range feels {
		for k2, v2 := range v {
			if _, prs := relations[k2+":"+k]; !prs {
				relations[k+":"+k2] = v2 + feels[k2][k]
			}
		}
	}
	for k, v := range relations {
		s := strings.Split(k, ":")
		feels[s[0]][s[1]] = v
		feels[s[1]][s[0]] = v
	}
	arrange := []int{}
	for _, n := range permutations([]int{0, 1, 2, 3, 4, 5, 6, 7}) {
		acc := 0
		for j := 0; j <= len(n)-1; j++ {
			if j == 7 {
				acc += feels[names[n[j]]][names[n[0]]]
			} else {
				acc += feels[names[n[j]]][names[n[j+1]]]
			}
		}
		arrange = append(arrange, acc)
	}
	sort.Ints(arrange)
	//part 2
	feels["Me"] = make(map[string]int, 0)
	for _, n := range names {
		feels["Me"][n] = 0
		feels[n]["Me"] = 0
	}
	names = append(names, "Me")
	arrange2 := []int{}
	for _, n := range permutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}) {
		acc := 0
		for j := 0; j <= len(n)-1; j++ {
			if j == 8 {
				acc += feels[names[n[j]]][names[n[0]]]
			} else {
				acc += feels[names[n[j]]][names[n[j+1]]]
			}
		}
		arrange2 = append(arrange2, acc)
	}
	sort.Ints(arrange2)
	return arrange[len(arrange)-1], arrange2[len(arrange2)-1]
}
func Y2015_14(input string) (interface{}, interface{}) {
	type reindeer struct {
		name                   string
		speed, endurance, rest int
	}
	type timer struct {
		distance, run_time, rest_time, points int
		resting                               bool
	}
	deers := []reindeer{}
	race := make(map[string]timer, 0)
	for _, n := range strings.Split(input, "\n") {
		parse := strings.Split(n, " ")
		speed, _ := strconv.Atoi(parse[3])
		endurance, _ := strconv.Atoi(parse[6])
		rest, _ := strconv.Atoi(parse[13])
		deers = append(deers, reindeer{parse[0], speed, endurance, rest})
		race[parse[0]] = timer{0, endurance, 0, 0, false}
	}
	// deers = append(deers, []reindeer{{"Comet", 14, 10, 127},{"Dancer", 16, 11, 162}}...)
	// race["Comet"] = timer{0,10,0,0,false}
	// race["Dancer"] = timer{0,11,0,0,false}
	for i := 1; i <= 2503; i++ {
		for _, n := range deers {
			entry := race[n.name]
			if entry.resting {
				entry.rest_time--
				if entry.rest_time == 0 {
					entry.resting = false
					entry.run_time = n.endurance
				}
			} else {
				entry.distance += n.speed
				entry.run_time--
				if entry.run_time == 0 {
					entry.resting = true
					entry.rest_time = n.rest
				}
			}
			race[n.name] = entry
		}
		winning := []string{}
		for k := range race {
			if len(winning) == 0 {
				winning = append(winning, k)
				continue
			}
			if race[k].distance > race[winning[0]].distance {
				temp := []string{k}
				winning = temp
			} else {
				if race[k].distance == race[winning[0]].distance {
					winning = append(winning, k)
				}
			}
		}
		// fmt.Printf("%v\t", winning)
		for _, n := range winning {
			entry := race[n]
			entry.points++
			race[n] = entry
		}
		// fmt.Printf("%04d:%v\n",i,race)
	}
	win, pointswin := 0, 0
	for k := range race {
		if race[k].distance > win {
			win = race[k].distance
		}
		if race[k].points > pointswin {
			pointswin = race[k].points
		}
		fmt.Printf("%s\t%d\n", k, race[k].points)
	}
	return win, pointswin
}
func Y2015_15(input string) (interface{}, interface{}) {
	ingredients := make(map[string][]int)
	for _, c := range strings.FieldsFunc(input, func(c rune) bool { return c == '\n' }) {
		m := strings.Split(c, ":")
		n := strings.Split(m[1], ",")
		nums := make([]int, 0)
		for _, o := range n {
			p := strings.Split(o, " ")
			a, _ := strconv.Atoi(p[2])
			nums = append(nums, a)
		}
		ingredients[m[0]] = nums
	}
	cookies := make([][]int, 0)
	litecookies := make([][]int, 0)
	for z := 1; z < 97; z++ {
		for y := 1; y < 97; y++ {
			for x := 1; x < 97; x++ {
				for w := 1; w < 97; w++ {
					if w+x+y+z == 100 {
						litecookies = append(litecookies, []int{w, x, y, z})
						if w+x < y {
							if x*5 > z {
								if z*5 > y {
									cookies = append(cookies, []int{w, x, y, z})
								}
							}
						}
					}
				}
			}
		}
	}
	highScore, calScore := 0, 0
	for i := range cookies {
		cap := cookies[i][0]*ingredients["Sprinkles"][0] + cookies[i][1]*ingredients["Butterscotch"][0] + cookies[i][2]*ingredients["Chocolate"][0] + cookies[i][3]*ingredients["Candy"][0]
		dur := cookies[i][0]*ingredients["Sprinkles"][1] + cookies[i][1]*ingredients["Butterscotch"][1] + cookies[i][2]*ingredients["Chocolate"][1] + cookies[i][3]*ingredients["Candy"][1]
		fla := cookies[i][0]*ingredients["Sprinkles"][2] + cookies[i][1]*ingredients["Butterscotch"][2] + cookies[i][2]*ingredients["Chocolate"][2] + cookies[i][3]*ingredients["Candy"][2]
		tex := cookies[i][0]*ingredients["Sprinkles"][3] + cookies[i][1]*ingredients["Butterscotch"][3] + cookies[i][2]*ingredients["Chocolate"][3] + cookies[i][3]*ingredients["Candy"][3]
		score := cap * dur * fla * tex
		if score > highScore {
			highScore = score
		}
	}
	for i := range litecookies {
		cap := litecookies[i][0]*ingredients["Sprinkles"][0] + litecookies[i][1]*ingredients["Butterscotch"][0] + litecookies[i][2]*ingredients["Chocolate"][0] + litecookies[i][3]*ingredients["Candy"][0]
		dur := litecookies[i][0]*ingredients["Sprinkles"][1] + litecookies[i][1]*ingredients["Butterscotch"][1] + litecookies[i][2]*ingredients["Chocolate"][1] + litecookies[i][3]*ingredients["Candy"][1]
		fla := litecookies[i][0]*ingredients["Sprinkles"][2] + litecookies[i][1]*ingredients["Butterscotch"][2] + litecookies[i][2]*ingredients["Chocolate"][2] + litecookies[i][3]*ingredients["Candy"][2]
		tex := litecookies[i][0]*ingredients["Sprinkles"][3] + litecookies[i][1]*ingredients["Butterscotch"][3] + litecookies[i][2]*ingredients["Chocolate"][3] + litecookies[i][3]*ingredients["Candy"][3]
		cal := litecookies[i][0]*ingredients["Sprinkles"][4] + litecookies[i][1]*ingredients["Butterscotch"][4] + litecookies[i][2]*ingredients["Chocolate"][4] + litecookies[i][3]*ingredients["Candy"][4]
		if cap < 0 {
			cap = 0
		}
		if dur < 0 {
			dur = 0
		}
		if fla < 0 {
			fla = 0
		}
		if tex < 0 {
			tex = 0
		}
		score := cap * dur * fla * tex
		if score > calScore && cal == 500 {
			calScore = score
		}
	}
	return highScore, calScore
}
func Y2015_16(input string) (interface{}, interface{}) {
	sues := make([]map[string]int, 1)
	sues[0] = map[string]int{"children": 3, "cats": 7, "samoyeds": 2, "pomeranians": 3, "akitas": 0, "vizslas": 0, "goldfish": 5, "trees": 5, "cars": 2, "perfumes": 1}
	for _, c := range strings.FieldsFunc(input, func(c rune) bool { return c == '\n' }) {
		thisSue := make(map[string]int)
		n := strings.Split(strings.Join(strings.Split(c, ":")[1:], ""), ",")
		for i := range n {
			keyval := strings.Split(strings.Trim(n[i], " "), " ")
			a, _ := strconv.Atoi(keyval[1])
			thisSue[keyval[0]] = a
		}
		sues = append(sues, thisSue)
	}
	gift, realGift := 0, 0
	for i, n := range sues {
		if i == 0 {
			continue
		}
		match, realMatch := 0, 0
		for k, v := range sues[0] {
			m, prs := n[k]
			if prs && m == v {
				match++
				if match > 2 {
					gift = i
				}
			}
			switch k {
			case "cats", "trees":
				if prs && v <= m {
					realMatch++
				}
			case "pomeranians", "goldfish":
				if prs && v >= m {
					realMatch++
				}
			default:
				if prs && v == m {
					realMatch++
				}
			}
			if realMatch > 2 && realGift == 0 {
				realGift = i
			}
		}
	}
	return gift, realGift
}
func Y2015_17(input string) (interface{}, interface{}) {
	buckets := make([]int, 0)
	eggnog := 150
	combos := make([][]int, 0)
	for _, n := range strings.Fields(input) {
		a, _ := strconv.Atoi(n)
		buckets = append(buckets, a)
	}
	for i := 1; i < len(buckets); i++ {
		for _, n := range combin.Combinations(len(buckets), i) {
			fill := make([]int, len(n))
			for j, m := range n {
				fill[j] = buckets[m]
			}
			sum := 0
			for i := range fill {
				sum += fill[i]
			}
			if sum == eggnog {
				combos = append(combos, fill)
			}
		}
	}
	fewest := make([][]int, 0)
	shortest := 99
	for i := range combos {
		if len(combos[i]) < shortest {
			shortest = len(combos[i])
		}
	}
	for i := range combos {
		if len(combos[i]) == shortest {
			fewest = append(fewest, combos[i])
		}
	}
	return len(combos), len(fewest)
}
