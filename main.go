package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

type puzzle struct {
	year, day     int
	input, output string
	solution      func(string) (interface{}, interface{})
}

func (p puzzle) solve() string {
	data, err := ioutil.ReadFile(p.input)
	check(err)
	a, b := p.solution(string(data))
	return fmt.Sprintf("*** %d *** [%d]\n"+p.output, p.year, p.day, a, b)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	puzz := map[string]puzzle{
		"2015 1": {2015, 1, "2015/1", "End Floor: %v\nEnter Basement at Position: %v\n", Y2015_01},
		"2015 2": {2015, 2, "2015/2", "Total Wrapping Paper: %v\nTotal Ribbon: %v\n", Y2015_02},
		"2015 3": {2015, 3, "2015/3", "Houses Visited This Year: %v\nHouses Visited Next Year: %v\n", Y2015_03},
		"2015 4": {2015, 4, "2015/4", "5 0s: %v\n6 0s: %v\n", Y2015_04},
		"2015 5": {2015, 5, "2015/5", "Nice Strings: %v\nNicer Strings: %v\n", Y2015_05},
		"2015 6": {2015, 6, "2015/6", "Lights Lit: %v\nTotal Brightness: %v\n", Y2015_06},
		"2015 7": {2015, 7, "2015/7", "Signal a: %v\nSignal a w/b override: %v\n", Y2015_07},
		"2015 8": {2015, 8, "2015/8", "Difference: %v\nRequote Difference: %v\n", Y2015_08},
		"2015 9": {2015, 9, "2015/9", "Shortest Route: %v\nLongest Route: %v\n", Y2015_09},
		"2018 1": {2018, 1, "2018/1", "Resulting Frequency: %v\nFirst Repeated Frequency: %v\n", Y2018_01},
		"2018 2": {2018, 2, "2018/2", "Checksum: %v\nCommon Letters: %v\n", Y2018_02},
		"2018 3": {2018, 3, "2018/3", "Square Inches Contested: %v\nNon-Overlapping Claim ID: %v\n", Y2018_03},
		"2018 4": {2018, 4, "2018/4", "Most Minutes Asleep: %v\nSleeps The Same Minute: %v\n", Y2018_04},
		"2018 5": {2018, 5, "2018/5", "Polymer Length: %v\nImproved Polymer Length: %v\n", Y2018_05},
		"2018 6": {2018, 6, "2018/6", "Largest area: %v\nMost Populous: %v\n", Y2018_06},
		"2018 7": {2018, 7, "2018/7", "Instruction Order: %v\nTime to Complete: %v\n", Y2018_07},
	}
	fDay := flag.String("day", "1", "which day to solve for")
	fYear := flag.String("year", "2015", "which year to solve for")
	fAll := flag.Bool("all", false, "output all solutions, overrides -day and -year")
	flag.Parse()
	if *fAll {
		for k := range puzz {
			fmt.Println(puzz[k].solve())
		}
	} else {
		i, prs := puzz[*fYear+" "+*fDay]
		if prs {
			fmt.Println(i.solve())
		} else {
			fmt.Println("Solution Not Found")
		}
	}
}
