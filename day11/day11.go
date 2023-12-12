package day11

import (
	"regexp"
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day11 struct{}

type Galaxy struct {
	x int
	y int
}

func (d Day11) Part1(input string) util.Solution {
	starttime := time.Now()

	//parse games
	galaxies := parseGalaxies(input, 2)

	//run
	total := processGalaxies(galaxies)

	return util.NewSolution(total, 1, time.Since(starttime))
}

func (d Day11) Part2(input string) util.Solution {
	starttime := time.Now()

	//parse games
	galaxies := parseGalaxies(input, 1000000)

	//run
	total := processGalaxies(galaxies)

	return util.NewSolution(total, 2, time.Since(starttime))
}

func parseGalaxies(input string, shiftfactor int) []Galaxy {
	//prettier shift factor
	shiftfactor--

	//find x-shifts
	xshiftsmap := map[int]int{}
	xshifts := 0

	reg := regexp.MustCompile(".*#.*")

	for c, column := range util.SplitContentRow(input) {
		if !reg.MatchString(column) {
			xshifts += shiftfactor
		}
		xshiftsmap[c] = xshifts
	}

	//find galaxies & y shifts
	galaxies := []Galaxy{}
	yshift := 0

	name := 1
	for y, line := range util.SplitContentLine(input) {

		//chars
		hasnogalaxy := true
		for x, char := range line {
			if char != '.' {
				hasnogalaxy = false
				galaxies = append(galaxies, Galaxy{x + xshiftsmap[x], y + yshift})
				name++
			}
		}

		//add yshift 1 if has no galaxis is still true
		if hasnogalaxy {
			yshift += shiftfactor
		}

	}
	return galaxies
}

func processGalaxies(galaxies []Galaxy) int {
	total := 0

	for i := len(galaxies) - 1; i >= 0; i-- {
		if len(galaxies) > 1 {
			//remove processing node
			current := galaxies[i]
			galaxies = append(galaxies[:i], galaxies[i+1:]...)

			total += pathToOtherGalaxies(current, galaxies)

		}
	}
	return total
}

func pathToOtherGalaxies(from Galaxy, to []Galaxy) int {
	total := 0

	for _, galaxy := range to {
		total += util.Abs(from.x-galaxy.x) + util.Abs(from.y-galaxy.y)
	}

	return total
}
