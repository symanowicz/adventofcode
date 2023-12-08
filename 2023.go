package main

import (
	//"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Y2023_01(input string) (interface{}, interface{}) {
	r := regexp.MustCompile(`[0-9]`)
	r2 := regexp.MustCompile(`[0-9]|oneight|twone|threeight|fiveight|sevenine|eightwo|eighthree|nineight|one|two|three|four|five|six|seven|eight|nine`)
	sum, sum2 := 0, 0
	for _, n := range strings.Fields(input) {
		m := r.FindAllString(n, -1)
		o := r2.FindAllString(n, -1)
		idx := 0
		ins := []string{}
		for x := range o {
			switch o[x] {
			case "one":
				o[x] = "1"
			case "two":
				o[x] = "2"
			case "three":
				o[x] = "3"
			case "four":
				o[x] = "4"
			case "five":
				o[x] = "5"
			case "six":
				o[x] = "6"
			case "seven":
				o[x] = "7"
			case "eight":
				o[x] = "8"
			case "nine":
				o[x] = "9"
			case "oneight":
				idx = x
				ins = append(ins, "1")
				ins = append(ins, "8")
			case "twone":
				idx = x
				ins = append(ins, "2")
				ins = append(ins, "1")
			case "threeight":
				idx = x
				ins = append(ins, "3")
				ins = append(ins, "8")
			case "fiveight":
				idx = x
				ins = append(ins, "5")
				ins = append(ins, "8")
			case "sevenine":
				idx = x
				ins = append(ins, "7")
				ins = append(ins, "9")
			case "eightwo":
				idx = x
				ins = append(ins, "8")
				ins = append(ins, "2")
			case "eighthree":
				idx = x
				ins = append(ins, "8")
				ins = append(ins, "3")
			case "nineight":
				idx = x
				ins = append(ins, "9")
				ins = append(ins, "8")
			}
		}
		if len(ins) > 0 {
			o = slices.Replace(o, idx,idx+1, ins...)
		}
		a, _ := strconv.Atoi(m[0] + m[len(m)-1])
		b, _ := strconv.Atoi(o[0] + o[len(o)-1])
		sum += a
		sum2 += b
	}
	return sum,sum2
}

func Y2023_02(input string) (interface{}, interface{}) {
	sum, sum2 := 0, 0
	for _, n := range strings.FieldsFunc(input, func(c rune) bool { return c == '\n' }) {
		m := strings.Split(n, ":")
		idx, _ := strconv.Atoi(strings.Split(m[0], " ")[1])
		//fmt.Printf("%d:\n", idx)
		games := strings.Split(strings.Trim(m[1], " "), "; ")
		counts := make([]map[string]int, 0)
		for _, o := range games {
			p := strings.Split(o, ", ")
			//fmt.Printf("\t%v\n", p)
			num := make(map[string]int, 0)
			for _, q := range p {
				r := strings.Split(q," ")
				//fmt.Printf("\t\t%v\n",r)
				s, _ := strconv.Atoi(r[0])
				num[r[1]] = s
			}
			counts = append(counts, num)
		}
		//fmt.Printf("%3d: %v\n",idx,counts)
		possible := true
		max := make(map[string]int)
		max["red"] = 0
		max["green"] = 0
		max["blue"] = 0
		for _, t := range counts {
			if t["red"] > 12 || t["green"] > 13 || t["blue"] > 14 {
				possible = false
			}
			if t["red"] > max["red"] {
				max["red"] = t["red"]
			}
			if t["green"] > max["green"] {
				max["green"] = t["green"]
			}
			if t["blue"] > max["blue"] {
				max["blue"] = t["blue"]
			}
		}
		if possible {
			sum += idx
		}
		sum2 += max["red"] * max["green"] * max["blue"]
	}
	return sum,sum2
}