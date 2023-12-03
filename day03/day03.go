package day03

import (
	"regexp"
	"strconv"
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day03 struct{}

var SYMCORD = map[int]map[int]bool{}

func (d Day03) Part1(input string) util.Solution {
	starttime := time.Now()

	//build symbol map
	reg := regexp.MustCompile("[^0-9|.|\n]")
	for index, line := range util.SplitContentLine(input) {
		//filter lines with no hits
		if !reg.MatchString(line) {
			continue
		}

		indicies := reg.FindAllStringIndex(line, -1)
		//init inner map of line
		SYMCORD[index] = make(map[int]bool)
		for _, indice := range indicies {
			//place cords
			SYMCORD[index][indice[0]] = true
		}

	}

	//find numbers
	total := 0
	regnum := regexp.MustCompile(`\d+`)
	for lineno, line := range util.SplitContentLine(input) {
		indices := regnum.FindAllStringIndex(line, -1)
		numbers := regnum.FindAllString(line, -1)
		for whichnum, indece := range indices {
			if computeLookArounds(lineno, indece) {
				num, _ := strconv.Atoi(numbers[whichnum])
				total += num
			}
		}
	}
	return util.NewSolution(total, 1, time.Since(starttime))
}

func (d Day03) Part2(input string) util.Solution {
	starttime := time.Now()

	//calculate total
	total := 0
	return util.NewSolution(total, 2, time.Since(starttime))
}

func computeLookArounds(lineNo int, indices []int) bool {
	//expand indeces
	expandindeces, _ := util.Range(indices[0], indices[len(indices)-1])
	//check lefthand side
	if SYMCORD[lineNo][expandindeces[0]-1] {
		return true
	}

	//check righthand side
	if SYMCORD[lineNo][expandindeces[len(expandindeces)-1]+1] {
		return true
	}

	//check top & bottom
	for _, indice := range expandindeces {
		//top
		if SYMCORD[lineNo-1][indice] {
			return true
		}
		//bottom
		if SYMCORD[lineNo+1][indice] {
			return true
		}
	}

	//check top diagonal
	//left
	if SYMCORD[lineNo-1][expandindeces[0]-1] {
		return true
	}
	//right
	if SYMCORD[lineNo-1][expandindeces[len(expandindeces)-1]+1] {
		return true
	}

	//check bottom diagonal
	//left
	if SYMCORD[lineNo+1][expandindeces[0]-1] {
		return true
	}
	//right
	if SYMCORD[lineNo+1][expandindeces[len(expandindeces)-1]+1] {
		return true
	}

	//no match
	return false
}
