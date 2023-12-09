package day08

import (
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day08 struct{}

var NEXTPOSLEFT = map[string]string{}
var NEXTPOSRIGHT = map[string]string{}
var DIRECTIONS string

func (d Day08) Part1(input string) util.Solution {
	starttime := time.Now()

	//init maps
	buildMaps(input)

	//run
	steps := getNeededSteps("AAA", map[string]struct{}{"ZZZ": {}})

	return util.NewSolution(steps, 1, time.Since(starttime))
}

func (d Day08) Part2(input string) util.Solution {
	starttime := time.Now()

	//init maps
	buildMaps(input)

	//determine starting points
	reg := regexp.MustCompile("[A-Z]{2}A")
	currentpositions := reg.FindAllString(input, -1)

	//determine possible enpoints
	reg = regexp.MustCompile("[A-Z]{2}Z")
	endpoints := map[string]struct{}{}

	for _, endpoint := range reg.FindAllString(input, -1) {
		endpoints[endpoint] = struct{}{}
	}

	//get all step lenghts
	steplengths := []int{}
	for _, currentpos := range currentpositions {
		steplengths = append(steplengths, getNeededSteps(currentpos, endpoints))
	}

	//LCM
	var lcm int64 = 1
	for _, steps := range steplengths {
		lcm = util.Lcm(lcm, int64(steps))
	}

	return util.NewSolution(lcm, 2, time.Since(starttime))
}

func getNeededSteps(currentpos string, endpoints map[string]struct{}) int {
	/*hint: As pointed out by this post on elixir forum (https://elixirforum.com/t/advent-of-code-2023-day-8/60244/7) every startingpoint does ONLY lead to one endpoint.
	after reaching this point it's also worth pointing out we restart the search cycle which lead to fix steps between hits - meaning the solution must be LCM between all ghost hits*/
	steps := 0
	for idx := 0; true; idx = idx % len(DIRECTIONS) {
		if DIRECTIONS[idx] == 'L' {
			currentpos = NEXTPOSLEFT[currentpos]
		} else {
			currentpos = NEXTPOSRIGHT[currentpos]
		}

		steps++
		if _, hit := endpoints[currentpos]; hit {
			return steps
		}

		idx++
	}
	return int(math.NaN())
}

func buildMaps(input string) {

	splitstr := strings.SplitN(input, "\n\n", 2)

	//build direction map
	reg := regexp.MustCompile("[A-Z]{3}")
	for _, instruction := range util.SplitContentLine(splitstr[1]) {
		positons := reg.FindAllString(instruction, -1)

		//left
		NEXTPOSLEFT[positons[0]] = positons[1]

		//right
		NEXTPOSRIGHT[positons[0]] = positons[2]
	}

	//build direction map
	DIRECTIONS = splitstr[0]
}
