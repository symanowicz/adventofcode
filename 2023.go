package main

import (
	"fmt"
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
		fmt.Printf("%v\t%v\n", n,o)
		sum += a
		sum2 += b
	}
	return sum,sum2
}