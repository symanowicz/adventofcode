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
		"2015 8": {2015,8,"2015/8","undefined: %d\nundefined: %d",Y2015_08},
		"2018 1": {2018,1,"2018/1","Resulting Frequency: %d\nFirst Repeated Frequency: %d",Y2018_01},
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
