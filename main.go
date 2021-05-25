package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

type puzzle struct {
	year, day int
	input, output string
	solution func(string) (int, int)
}

func (p puzzle) solve() string {
	data, err := ioutil.ReadFile(p.input)
	check(err)
	a, b := p.solution(string(data))
	return fmt.Sprintf("*** %d *** [%d]\n" + p.output, p.year, p.day, a, b)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	puzz := map[string]puzzle{
		"2015 1": {2015,1,"2015/1","End Floor: %d\nEnter Basement at Position: %d\n",Y2015_01},
		"2015 2": {2015,2,"2015/2","Total Wrapping Paper: %d\nTotal Ribbon: %d\n",Y2015_02},
		"2015 3": {2015,3,"2015/3","Houses Visited This Year: %d\nHouses Visited Next Year: %d\n",Y2015_03},
		"2015 4": {2015,4,"2015/4","5 0s: %d\n6 0s: %d\n",Y2015_04},
		"2015 5": {2015,5,"2015/5","Nice Strings: %d\nNicer Strings: %d\n",Y2015_05},
		"2015 6": {2015,6,"2015/6","Lights Lit: %d\nTotal Brightness: %d\n",Y2015_06},
		"2015 7": {2015,7,"2015/7","Signal a: %d\nSignal a w/b override: %d",Y2015_07},
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
		fmt.Println(puzz[*fYear+" "+*fDay].solve())
	}
}
