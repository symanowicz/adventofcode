package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type puzzle struct {
	year, day int
	output    string
	solution  func(string) (interface{}, interface{})
}

func (p puzzle) solve() string {
	if _, e := os.Stat(fmt.Sprintf("%d", p.year)); errors.Is(e, os.ErrNotExist) {
		e = os.Mkdir(fmt.Sprintf("%d", p.year), 0770)
		check(e)
	}
	if _, e := os.Stat(fmt.Sprintf("%d/%d", p.year, p.day)); errors.Is(e, os.ErrNotExist) {
		sess := os.Getenv("AOC_SESSION")
		if len(sess) == 0 {
			log.Fatal("no input file found for this puzzle and no session key provided to fetch input...exiting")
		}
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", p.year, p.day), nil)
		check(err)
		req.AddCookie(&http.Cookie{Name: "session", Value: sess})
		resp, err := client.Do(req)
		check(err)
		s, err := io.ReadAll(resp.Body)
		check(err)
		err = os.WriteFile(fmt.Sprintf("%d/%d", p.year, p.day), s, 0660)
		check(err)
	}
	data, err := os.ReadFile(fmt.Sprintf("%d/%d", p.year, p.day))
	check(err)
	a, b := p.solution(string(data))
	return fmt.Sprintf("*** %d *** [%d]\n"+p.output, p.year, p.day, a, b)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	puzz := map[string]puzzle{
		"2015 1":  {2015, 1, "End Floor: %v\nEnter Basement at Position: %v\n", Y2015_01},
		"2015 2":  {2015, 2, "Total Wrapping Paper: %v\nTotal Ribbon: %v\n", Y2015_02},
		"2015 3":  {2015, 3, "Houses Visited This Year: %v\nHouses Visited Next Year: %v\n", Y2015_03},
		"2015 4":  {2015, 4, "5 0s: %v\n6 0s: %v\n", Y2015_04},
		"2015 5":  {2015, 5, "Nice Strings: %v\nNicer Strings: %v\n", Y2015_05},
		"2015 6":  {2015, 6, "Lights Lit: %v\nTotal Brightness: %v\n", Y2015_06},
		"2015 7":  {2015, 7, "Signal a: %v\nSignal a w/b override: %v\n", Y2015_07},
		"2015 8":  {2015, 8, "Difference: %v\nRequote Difference: %v\n", Y2015_08},
		"2015 9":  {2015, 9, "Shortest Route: %v\nLongest Route: %v\n", Y2015_09},
		"2015 10": {2015, 10, "40x Length: %v\n50x Length: %v\n", Y2015_10},
		"2015 11": {2015, 11, "Next Password: %v\nNexter Password: %v\n", Y2015_11},
		"2015 12": {2015, 12, "Sum: %v\nWithout Red: %v\n", Y2015_12},
		"2015 13": {2015, 13, "Total Happy: %v\nPlus Me: %v\n", Y2015_13},
		"2015 14": {2015, 14, "Winning Distance: %v\nWinning Points: %v\n", Y2015_14},
		"2018 1":  {2018, 1, "Resulting Frequency: %v\nFirst Repeated Frequency: %v\n", Y2018_01},
		"2018 2":  {2018, 2, "Checksum: %v\nCommon Letters: %v\n", Y2018_02},
		"2018 3":  {2018, 3, "Square Inches Contested: %v\nNon-Overlapping Claim ID: %v\n", Y2018_03},
		"2018 4":  {2018, 4, "Most Minutes Asleep: %v\nSleeps The Same Minute: %v\n", Y2018_04},
		"2018 5":  {2018, 5, "Polymer Length: %v\nImproved Polymer Length: %v\n", Y2018_05},
		"2018 6":  {2018, 6, "Largest area: %v\nMost Populous: %v\n", Y2018_06},
		"2018 7":  {2018, 7, "Instruction Order: %v\nTime to Complete: %v\n", Y2018_07},
		"2022 1":  {2022, 1, "Most Calories: %v\nTop Three: %v\n", Y2022_01},
		"2022 2":  {2022, 2, "Total Score: %v\nCorrect Total Score: %v\n", Y2022_02},
		"2022 3":  {2022, 3, "Sum of Duplicates: %v\nSum of Badges: %v\n", Y2022_03},
		"2022 4":  {2022, 4, "Total Fully Contained: %v\nAny Overlap: %v\n", Y2022_04},
		"2022 5":  {2022, 5, "Top Crates for Cratemover 9000: %v\nTop Crates for Cratemover 9001: %v\n", Y2022_05},
		"2022 6":  {2022, 6, "Characters to start-of-packet: %v\nCharacters to start-of-message: %v\n", Y2022_06},
		"2022 7":  {2022, 7, "Sum of Directories under 100K: %v\nSmallest directory to make room for update: %v\n", Y2022_07},
		"2022 8":  {2022, 8, "Visible Trees: %v\nHighest scenic score: %v\n", Y2022_08},
	}
	fDay := flag.String("day", "1", "which day to solve for")
	fYear := flag.String("year", "2015", "which year to solve for")
	fAll := flag.Bool("all", false, "output all solutions, overrides -day and -year")
	flag.Parse()
	if *fAll {
		years := []int{2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022}
		days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
		for i := range years {
			for j := range days {
				p, prs := puzz[fmt.Sprintf("%d %d", years[i], days[j])]
				if prs {
					fmt.Println(p.solve())
				}
			}
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
