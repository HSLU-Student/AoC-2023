package day01

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day01 struct{}

var DIGITS = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func (d Day01) Part1(input string) util.Solution {
	starttime := time.Now()

	total := 0
	reg := regexp.MustCompile("[1-9]")
	for _, line := range util.SplitContentLine(input) {
		//get numbers form regex
		numbers := reg.FindAllString(line, -1)

		//combine numbers
		parsenum, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		if err != nil {
			break
		}
		total += parsenum
	}

	return util.NewSolution(total, 1, time.Since(starttime))
}

func (d Day01) Part2(input string) util.Solution {
	starttime := time.Now()

	total := 0
	//everthing that could contain something
	reg := regexp.MustCompile("o|t|f|s|e|n")

	for _, line := range util.SplitContentLine(input) {
		numbers := []string{}
		for index, char := range strings.Split(line, "") {
			//check char could be beginning of a digit word & is long enough for a match
			if reg.MatchString(char) && len(line[index:]) >= 3 {
				digit, exists := StringToDigitParse(line[index:])
				if exists {
					numbers = append(numbers, digit)
					continue
				}
			}

			//very disgusting solution - if we dont get an error stringconv was successful & char is a digit
			_, err := strconv.Atoi(char)
			if err == nil {
				numbers = append(numbers, char)
			}
		}
		parsenum, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		if err != nil {
			break
		}
		total += parsenum
	}
	return util.NewSolution(total, 2, time.Since(starttime))
}

func StringToDigitParse(pattern string) (string, bool) {
	patternlen := len(pattern)
	possiblewords := []string{}

	if patternlen >= 5 {
		possiblewords = append(possiblewords, pattern[:5])
	}

	if patternlen >= 4 {
		possiblewords = append(possiblewords, pattern[:4])
	}

	possiblewords = append(possiblewords, pattern[:3])

	for _, candiate := range possiblewords {
		digit, exists := DIGITS[candiate]
		if exists {
			return digit, true
		}
	}
	return "", false
}
