package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
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
	outcomes := []int{4, 1, 7, 8, 5, 2, 3, 9, 6}
	outcomes2 := []int{3, 1, 2, 4, 5, 6, 8, 9, 7}
	for _, n := range strings.FieldsFunc(input, func(c rune) bool { return c == '\n' }) {
		total += outcomes[(n[2]-88)*3+(n[0]-65)]
		total2 += outcomes2[(n[2]-88)*3+(n[0]-65)]
	}
	return total, total2
}
func Y2022_03(input string) (interface{}, interface{}) {
	sumDupes, sumBadges := 0, 0
	groups := make([]string, 0)
	currGroup := ""
	for i, n := range strings.Fields(input) {
		if (i+1)%3 == 0 {
			currGroup += n
			groups = append(groups, currGroup)
			currGroup = ""
		} else {
			currGroup += n + " "
		}
		first, second := n[:len(n)/2], n[len(n)/2:]
		for _, c := range first {
			if strings.ContainsRune(second, rune(c)) {
				if c > 96 {
					sumDupes += int(c) - 96
				} else {
					sumDupes += int(c) - 38
				}
				break
			}
		}
	}
	for _, n := range groups {
		sacks := strings.Split(n, " ")
		for _, c := range sacks[0] {
			if strings.ContainsRune(sacks[1], rune(c)) && strings.ContainsRune(sacks[2], rune(c)) {
				if c > 96 {
					sumBadges += int(c) - 96
				} else {
					sumBadges += int(c) - 38
				}
				break
			}
		}
	}
	return sumDupes, sumBadges
}
func Y2022_04(input string) (interface{}, interface{}) {
	totalContained, anyOverlap := 0, 0
	for _, n := range strings.Fields(input) {
		a := strings.Split(n, ",")
		b1 := strings.Split(a[0], "-")
		b2 := strings.Split(a[1], "-")
		c1, _ := strconv.Atoi(b1[0])
		c2, _ := strconv.Atoi(b1[1])
		c3, _ := strconv.Atoi(b2[0])
		c4, _ := strconv.Atoi(b2[1])
		elf1 := make([]int, 0)
		elf2 := make([]int, 0)
		for i := c1; i <= c2; i++ {
			elf1 = append(elf1, i)
		}
		for i := c3; i <= c4; i++ {
			elf2 = append(elf2, i)
		}
		for _, m := range elf1 {
			if slices.Contains(elf2, m) {
				anyOverlap++
				break
			}
		}
		fullyContained := false
		for _, m := range elf1 {
			if !slices.Contains(elf2, m) {
				fullyContained = false
				break
			}
			fullyContained = true
		}
		if fullyContained {
			totalContained++
			continue
		}
		for _, m := range elf2 {
			if !slices.Contains(elf1, m) {
				fullyContained = false
				break
			}
			fullyContained = true
		}
		if fullyContained {
			totalContained++
		}
	}
	return totalContained, anyOverlap
}
func Y2022_05(input string) (interface{}, interface{}) {
	instructionStart := 0
	stacks9000 := make(map[int][]rune)
	stacks9001 := make(map[int][]rune)
	start := make([]string, 0)
	for i, n := range strings.FieldsFunc(input, func(c rune) bool { return c == '\n' }) {
		if strings.Contains(n, "move") {
			instructionStart = i
			break
		}
		start = append(start, n)
	}
	for l, r := 0, len(start)-1; l < r; l, r = l+1, r-1 {
		start[l], start[r] = start[r], start[l]
	}
	for _, n := range strings.Fields(start[0]) {
		i, _ := strconv.Atoi(n)
		stacks9000[i] = make([]rune, 0)
		stacks9001[i] = make([]rune, 0)
	}
	for i := 1; i <= len(stacks9000); i++ {
		for j, c := range start {
			if j == 0 {
				continue
			}
			if rune(c[i*4-3]) == ' ' {
				continue
			}
			stacks9000[i] = append(stacks9000[i], rune(c[i*4-3]))
			stacks9001[i] = append(stacks9001[i], rune(c[i*4-3]))
		}
	}
	for i, n := range strings.FieldsFunc(input, func(c rune) bool { return c == '\n' }) {
		if i < instructionStart {
			continue
		}
		order := strings.Fields(n)
		amt, _ := strconv.Atoi(order[1])
		from, _ := strconv.Atoi(order[3])
		to, _ := strconv.Atoi(order[5])
		for j := 0; j < amt; j++ {
			var mv rune
			mv, stacks9000[from] = stacks9000[from][len(stacks9000[from])-1], stacks9000[from][:len(stacks9000[from])-1]
			stacks9000[to] = append(stacks9000[to], mv)
		}
		mv := make([]rune, amt)
		short := make([]rune, len(stacks9001[from])-amt)
		copy(mv, stacks9001[from][len(stacks9001[from])-amt:len(stacks9001[from])])
		copy(short, stacks9001[from][:len(stacks9001[from])-amt])
		fmt.Printf("%v\t%v\n", short, mv)
		stacks9001[from] = short
		stacks9001[to] = append(stacks9001[to], mv...)
	}
	ans9000, ans9001 := "", ""
	for i := 1; i <= len(stacks9000); i++ {
		ans9000 += string(stacks9000[i][len(stacks9000[i])-1])
		ans9001 += string(stacks9001[i][len(stacks9001[i])-1])
	}
	return ans9000, ans9001
}
func Y2022_06(input string) (interface{}, interface{}) {
	startPacket, startMessage := "", ""
	totalPacket, totalMessage := 0, 0
	for i := range input {
		totalPacket++
		startPacket = input[i : i+4]
		count := make(map[rune]int)
		for _, c := range startPacket {
			count[c]++
		}
		found := true
		for k := range count {
			if count[k] > 1 {
				found = false
			}
		}
		if found {
			break
		}
	}
	for i := range input {
		totalMessage++
		startMessage = input[i : i+14]
		count := make(map[rune]int)
		for _, c := range startMessage {
			count[c]++
		}
		found := true
		for k := range count {
			if count[k] > 1 {
				found = false
			}
		}
		if found {
			break
		}
	}
	return totalPacket + 3, totalMessage + 13
}
func Y2022_07(input string) (interface{}, interface{}) {
	type file struct {
		name string
		size int
	}
	type dir struct {
		name        string
		parent      *dir
		depth, size int
		childs      []*dir
		files       []*file
	}
	root := dir{"/", nil, 0, 0, make([]*dir, 0), make([]*file, 0)}
	currentDir := &root
	slicedInput := strings.FieldsFunc(input, func(c rune) bool { return c == '\n' })
	for i, n := range slicedInput {
		if i == 0 {
			continue
		}
		data := strings.Fields(n)
		if data[0] == "$" {
			if data[1] == "cd" {
				//change directory
				if data[2] == ".." {
					//go up
					currentDir = currentDir.parent
				} else {
					//go down
					for j, m := range currentDir.childs {
						if m.name == data[2] {
							currentDir = currentDir.childs[j]
						}
					}
				}
			} else {
				//process file list
				currentLine := i
				for {
					currentLine++
					if currentLine == len(slicedInput) {
						break
					}
					if strings.Fields(slicedInput[currentLine])[0] == "$" {
						break
					}
				}
				for _, m := range slicedInput[i+1 : currentLine] {
					if strings.Fields(m)[0] == "dir" {
						//add directory
						currentDir.childs = append(currentDir.childs, &dir{strings.Fields(m)[1], currentDir, currentDir.depth + 1, 0, make([]*dir, 0), make([]*file, 0)})
					} else {
						//add file
						size, _ := strconv.Atoi(strings.Fields(m)[0])
						currentDir.files = append(currentDir.files, &file{strings.Fields(m)[1], size})
					}
				}
			}
		} else {
			continue
		}
	}
	hundredKtotal := 0
	spaceRequired := 30000000 - (70000000 - root.size)
	deleteSize := root.size + 1
	var findAnswers func(d *dir)
	var calcDirSize func(d *dir)
	calcDirSize = func(d *dir) {
		if len(d.childs) != 0 {
			for i := range d.childs {
				calcDirSize(d.childs[i])
				d.size += d.childs[i].size
			}
		}
		for i := range d.files {
			d.size += d.files[i].size
		}
	}
	findAnswers = func(d *dir) {
		if len(d.childs) != 0 {
			for i := range d.childs {
				findAnswers(d.childs[i])
			}
		}
		if d.size <= 100000 {
			hundredKtotal += d.size
		}
		if d.size > spaceRequired && d.size < deleteSize {
			deleteSize = d.size
		}
	}
	calcDirSize(&root)
	findAnswers(&root)
	return hundredKtotal, deleteSize
}
func Y2022_08(input string) (interface{}, interface{}) {
	type tree struct {
		height  int
		visible bool
		score   int
		north   []int
		south   []int
		east    []int
		west    []int
	}
	trees := make([]tree, 0)
	layout := [][]int{}
	for _, c := range strings.FieldsFunc(input, func(c rune) bool { return c == '\n' }) {
		line := make([]int, 0)
		for _, n := range c {
			a, _ := strconv.Atoi(string(n))
			line = append(line, a)
		}
		layout = append(layout, line)
	}
	for i := range layout {
		for j := range layout[i] {
			z := tree{layout[i][j], false, 0, make([]int, 0), make([]int, 0), make([]int, 0), make([]int, 0)}
			n, s, e, w := true, true, true, true
			for k := i; k > 0; {
				k--
				z.north = append(z.north, layout[k][j])
				if z.height <= layout[k][j] {
					n = false
					break
				}
			}
			for k := i; k < 98; {
				k++
				z.south = append(z.south, layout[k][j])
				if z.height <= layout[k][j] {
					s = false
					break
				}
			}
			for k := j; k > 0; {
				k--
				z.west = append(z.west, layout[i][k])
				if z.height <= layout[i][k] {
					w = false
					break
				}
			}
			for k := j; k < 98; {
				k++
				z.east = append(z.east, layout[i][k])
				if z.height <= layout[i][k] {
					e = false
					break
				}
			}
			if n || s || w || e {
				z.visible = true
			}
			z.score = len(z.north) * len(z.south) * len(z.east) * len(z.west)
			trees = append(trees, z)
		}
	}
	vis, highScore := 0, 0
	for i := range trees {
		if trees[i].visible {
			vis++
		}
		if trees[i].score > highScore {
			highScore = trees[i].score
		}
	}
	return vis, highScore
}
