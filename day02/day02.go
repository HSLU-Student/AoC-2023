package day02

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day02 struct{}

var maxValue = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func (d Day02) Part1(input string) util.Solution {
	starttime := time.Now()
	//build game registry
	total := 0
	for index, line := range util.SplitContentLine(input) {
		gamevalid := true
		parsedline := ParseGame(line)
		for _, pick := range parsedline {
			num, _ := strconv.Atoi(pick[0])
			if maxValue[pick[1]] < num {
				gamevalid = false
				break
			}
		}

		if gamevalid {
			total += index + 1
		}
	}

	return util.NewSolution(total, 1, time.Since(starttime))
}

func (d Day02) Part2(input string) util.Solution {
	starttime := time.Now()
	total := 0
	for _, line := range util.SplitContentLine(input) {
		parsedline := ParseGame(line)
		gamemap := map[string]int{
			"red":   1, //init with 1 to prevent erasure by 0 multiplication (only two cube colors exists)
			"green": 1,
			"blue":  1,
		}

		//find largest color draws
		for _, pick := range parsedline {
			num, _ := strconv.Atoi(pick[0])
			color := pick[1]
			if gamemap[color] < num {
				gamemap[color] = num
			}
		}

		//calculate power
		power := 1
		for _, num := range gamemap {
			power *= num
		}
		total += power
	}
	return util.NewSolution(total, 2, time.Since(starttime))
}

func ParseGame(gamestring string) [][]string {
	//remove game number
	reg := regexp.MustCompile("Game [0-9]*: ")
	nstring := reg.ReplaceAllString(gamestring, "")

	//combine sets & trim spaces
	nstring = strings.ReplaceAll(nstring, ";", ",")

	//split number & color
	picks := strings.Split(nstring, ",")

	slicedset := [][]string{}
	for _, pick := range picks {
		slicedset = append(slicedset, strings.Split(strings.Trim(pick, " "), " "))
	}

	return slicedset
}
