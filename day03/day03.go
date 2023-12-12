package day03

import (
	"math"
	"regexp"
	"strconv"
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day03 struct{}

type CordMap map[int]map[int]bool

func (d Day03) Part1(input string) util.Solution {
	starttime := time.Now()

	//get symbol cord map
	reg := regexp.MustCompile("[^0-9|.|\n]")
	crmp := newCordGrid(input, *reg)

	//find numbers
	total := 0
	regnum := regexp.MustCompile(`\d+`)
	for lineno, line := range util.SplitContentLine(input) {
		indices := regnum.FindAllStringIndex(line, -1)
		numbers := regnum.FindAllString(line, -1)
		for whichnum, indece := range indices {
			haslookaround, _, _ := computeLookArounds(crmp, lineno, indece)
			if haslookaround {
				num, _ := strconv.Atoi(numbers[whichnum])
				total += num
			}
		}
	}
	return util.NewSolution(total, 1, time.Since(starttime))
}

func (d Day03) Part2(input string) util.Solution {
	starttime := time.Now()

	//get symbol cord map
	reg := regexp.MustCompile(`\*`)
	crmp := newCordGrid(input, *reg)

	//build result map
	resmap := map[int]map[int][]int{}
	regnum := regexp.MustCompile(`\d+`)
	for lineno, line := range util.SplitContentLine(input) {
		indices := regnum.FindAllStringIndex(line, -1)
		numbers := regnum.FindAllString(line, -1)
		for whichnum, indece := range indices {
			haslookaround, mline, mpos := computeLookArounds(crmp, lineno, indece)
			if haslookaround {
				_, haskey := resmap[mline]
				//init inner map if not already existing
				if !haskey {
					resmap[mline] = map[int][]int{}
				}
				num, _ := strconv.Atoi(numbers[whichnum])
				resmap[mline][mpos] = append(resmap[mline][mpos], num)
			}
		}
	}
	//calculate total
	total := 0
	for _, line := range resmap {
		for _, gearset := range line {
			if len(gearset) == 2 {
				total += gearset[0] * gearset[1]
			}
		}
	}
	return util.NewSolution(total, 2, time.Since(starttime))
}

func newCordGrid(input string, reg regexp.Regexp) CordMap {
	crmp := CordMap{}
	for index, line := range util.SplitContentLine(input) {
		//filter lines with no hits
		if !reg.MatchString(line) {
			continue
		}

		indicies := reg.FindAllStringIndex(line, -1)
		//init inner map of line
		crmp[index] = make(map[int]bool)
		for _, indice := range indicies {
			//place cords
			crmp[index][indice[0]] = true
		}

	}
	return crmp
}

func computeLookArounds(crmp CordMap, lineNo int, indices []int) (bool, int, int) {
	//expand indeces
	expandindeces, _ := util.Range(indices[0], indices[len(indices)-1])

	//check lefthand side
	if crmp[lineNo][expandindeces[0]-1] {
		return true, lineNo, expandindeces[0] - 1
	}

	//check righthand side
	if crmp[lineNo][expandindeces[len(expandindeces)-1]+1] {
		return true, lineNo, expandindeces[len(expandindeces)-1] + 1
	}

	//check top & bottom
	for _, indice := range expandindeces {
		//top
		if crmp[lineNo-1][indice] {
			return true, lineNo - 1, indice
		}
		//bottom
		if crmp[lineNo+1][indice] {
			return true, lineNo + 1, indice
		}
	}

	//check top diagonal
	//left
	if crmp[lineNo-1][expandindeces[0]-1] {
		return true, lineNo - 1, expandindeces[0] - 1
	}
	//right
	if crmp[lineNo-1][expandindeces[len(expandindeces)-1]+1] {
		return true, lineNo - 1, expandindeces[len(expandindeces)-1] + 1
	}

	//check bottom diagonal
	//left
	if crmp[lineNo+1][expandindeces[0]-1] {
		return true, lineNo + 1, expandindeces[0] - 1
	}
	//right
	if crmp[lineNo+1][expandindeces[len(expandindeces)-1]+1] {
		return true, lineNo + 1, expandindeces[len(expandindeces)-1] + 1
	}

	//no match
	nan := int(math.NaN())
	return false, nan, nan
}
