package day05

import (
	"math"
	"strings"
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day05 struct{}

func (d Day05) Part1(input string) util.Solution {
	starttime := time.Now()
	splitstr := strings.Split(input, "\n\n")

	//build remap table
	remaps := [][][]int{}
	for _, remap := range splitstr[1:] {
		remaps = append(remaps, buildRemap(remap))
	}

	//run
	lowest := math.MaxInt32
	for _, seed := range util.ParseNumbers(splitstr[0]) {
		location := getLocation(seed, remaps)
		if location < lowest {
			lowest = location
		}
	}

	return util.NewSolution(lowest, 1, time.Since(starttime))
}

func (d Day05) Part2(input string) util.Solution {
	starttime := time.Now()
	splitstr := strings.Split(input, "\n\n")

	//build seed map
	remaps := [][][]int{}
	for _, remap := range splitstr[1:] {
		remaps = append(remaps, buildRemap(remap))
	}

	//run
	lowest := math.MaxInt32

	//already calcuated values no need for recalcuation
	lowerbounderies := 0
	upperbounderies := 0

	seeds := util.ParseNumbers(splitstr[0])
	for i := 0; i < len(seeds); i += 2 {
		lowerbound := seeds[i]
		upperbound := seeds[i] + seeds[i+1]
		for j := lowerbound; j < upperbound; j++ {
			//elements calcuated in round before
			if j > lowerbounderies && j < upperbounderies {
				continue
			}
			location := getLocation(j, remaps)
			if location < lowest {
				lowest = location
			}
		}
		lowerbounderies = lowerbound
		upperbounderies = upperbound
	}

	return util.NewSolution(lowest, 2, time.Since(starttime))
}

func buildRemap(remap string) [][]int {
	remaping := [][]int{}
	for _, line := range strings.Split(remap, "\n")[1:] {
		parsednum := util.ParseNumbers(line)

		//append last item so we dont need to recalcuate
		parsednum = append(parsednum, parsednum[1]+parsednum[2])

		//append to remap
		remaping = append(remaping, parsednum)
	}
	return remaping
}

func getLocation(seed int, remaps [][][]int) int {
	for _, remap := range remaps {
		for _, entry := range remap {
			offset := seed - entry[1]
			if offset > -1 && entry[1]+offset < entry[3] {
				seed = entry[0] + offset
				break
			}
		}
	}
	return seed
}
