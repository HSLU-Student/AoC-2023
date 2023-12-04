package day04

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day04 struct{}

func (d Day04) Part1(input string) util.Solution {
	starttime := time.Now()

	total := 0
	reg := regexp.MustCompile("^Card +[0-9]+: +")
	for _, card := range util.SplitContentLine(input) {
		content := strings.Split(reg.ReplaceAllString(card, ""), " | ")

		//little hack to build a set
		winningnum := map[int]struct{}{}
		for _, num := range ParseNumbers(content[0]) {
			winningnum[num] = struct{}{}
		}

		//query set for hits
		cardtotal := 0
		for _, num := range ParseNumbers(content[1]) {
			_, hit := winningnum[num]
			if hit {
				if cardtotal == 0 {
					cardtotal = 1
				} else {
					cardtotal = cardtotal * 2
				}
			}
		}
		total += cardtotal
	}

	return util.NewSolution(total, 1, time.Since(starttime))
}

func (d Day04) Part2(input string) util.Solution {
	starttime := time.Now()

	//building recursion map
	recursionmap := map[int][]int{}
	reg := regexp.MustCompile("^Card +[0-9]+: +")
	for game, card := range util.SplitContentLine(input) {
		content := strings.Split(reg.ReplaceAllString(card, ""), " | ")

		//little hack to build a set
		winningnum := map[int]struct{}{}
		for _, num := range ParseNumbers(content[0]) {
			winningnum[num] = struct{}{}
		}

		//query set for no hits
		hits := 0
		for _, num := range ParseNumbers(content[1]) {
			_, hit := winningnum[num]
			if hit {
				hits += 1
			}
		}

		//build recursion map
		if hits != 0 {
			copycards, _ := util.Range(game+2, game+2+hits)
			recursionmap[game+1] = copycards
		}
	}

	//now run everthing against it
	total := 0
	lookuptable := map[int]int{}
	for game := range util.SplitContentLine(input) {
		total += RecursiveGameExecution(game+1, recursionmap, lookuptable)
	}
	return util.NewSolution(total, 2, time.Since(starttime))
}

func ParseNumbers(numstr string) []int {
	reg := regexp.MustCompile(`\d+`)
	numbersstr := reg.FindAllString(numstr, -1)

	numbers := []int{}
	for _, numstr := range numbersstr {
		numi, _ := strconv.Atoi(numstr)
		numbers = append(numbers, numi)
	}
	return numbers
}

func RecursiveGameExecution(gameno int, recursionmap map[int][]int, lookuptable map[int]int) int {
	//shortcut if lookuptable entry exists
	lookup, exists := lookuptable[gameno]
	if exists {
		return lookup
	}
	_, gencopys := recursionmap[gameno]
	if gencopys {
		recres := 0
		for _, copy := range recursionmap[gameno] {
			recres += RecursiveGameExecution(copy, recursionmap, lookuptable)
		}
		lookuptable[gameno] = recres + 1
		return lookuptable[gameno] //+1 because every card by itself has a value of 1
	} else {
		return 1
	}
}
