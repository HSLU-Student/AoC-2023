package day10

import (
	"strings"
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day10 struct{}

// Directions & there representation in the coordinate field
type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

// Pipes & where the start & end
var PIPEMAPPING = map[string][]Direction{
	"|": {NORTH, SOUTH},
	"-": {EAST, WEST},
	"L": {NORTH, EAST},
	"J": {NORTH, WEST},
	"7": {SOUTH, WEST},
	"F": {SOUTH, EAST},
}

type Pipe struct {
	x    int
	y    int
	pipe string
	from Direction
}

func (d Day10) Part1(input string) util.Solution {
	starttime := time.Now()

	//build pipe plot & get starting point
	pipeplot, startingpipe := buildPlot(input)

	//find pipe connect to starting pipe
	currentpipe := findStartConnectedPipe(startingpipe, pipeplot)

	//run
	steps := 1 //init with 1 because finding the connected pipe of starting pipe we already took a step
	for currentpipe.pipe != "S" {
		currentpipe = nextPipe(currentpipe, pipeplot)

		steps++
	}

	//solution must be steps/2
	return util.NewSolution(steps/2, 1, time.Since(starttime))
}

func (d Day10) Part2(input string) util.Solution {
	starttime := time.Now()
	res := "Not implemeted yet..."
	return util.NewSolution(res, 2, time.Since(starttime))
}

func buildPlot(input string) ([][]string, Pipe) {
	//build pipe plot & get starting point (inefficent way to search the startpoint)
	pipeplot := [][]string{}
	startingpipe := Pipe{}

	for idx, line := range util.SplitContentLine(input) {
		coordline := strings.Split(line, "")

		pipeplot = append(pipeplot, coordline)

		if cord := util.ContainsAtIndex[string](coordline, "S"); cord != -1 {
			startingpipe = Pipe{y: idx, x: cord, pipe: "S"}
		}
	}
	return pipeplot, startingpipe
}

func nextPipe(current Pipe, pipeplot [][]string) Pipe {
	//select outgoing direction of current pipe (direction which we did not come from)
	endings := PIPEMAPPING[current.pipe]

	var outgoing Direction
	for _, ending := range endings {
		if ending != current.from {
			outgoing = ending
		}
	}

	//determine new coord & return new pipe
	if outgoing == NORTH {
		return Pipe{x: current.x, y: current.y - 1, from: SOUTH, pipe: pipeplot[current.y-1][current.x]}
	}

	if outgoing == EAST {
		return Pipe{x: current.x + 1, y: current.y, from: WEST, pipe: pipeplot[current.y][current.x+1]}
	}

	if outgoing == SOUTH {
		return Pipe{x: current.x, y: current.y + 1, from: NORTH, pipe: pipeplot[current.y+1][current.x]}
	}

	if outgoing == WEST {
		return Pipe{x: current.x - 1, y: current.y, from: EAST, pipe: pipeplot[current.y][current.x-1]}
	}

	return Pipe{}
}

func findStartConnectedPipe(startpipe Pipe, pipeplot [][]string) Pipe {

	//as soon as we found the first possibility - return it & only need to check 3 options as we will find a path for sure

	//check top
	if startpipe.y != 0 {
		//check if top pipe is a pipe with ending in south direction
		if containsEnding(pipeplot[startpipe.y-1][startpipe.x], SOUTH) {
			return Pipe{x: startpipe.x, y: startpipe.y - 1, pipe: pipeplot[startpipe.y-1][startpipe.x]}
		}
	}

	//check bottom
	if startpipe.y != len(pipeplot)-1 {
		if containsEnding(pipeplot[startpipe.y+1][startpipe.x], NORTH) {
			return Pipe{x: startpipe.x, y: startpipe.y + 1, pipe: pipeplot[startpipe.y+1][startpipe.x]}
		}
	}

	//check left
	if startpipe.x != 0 {
		if containsEnding(pipeplot[startpipe.y][startpipe.x-1], EAST) {
			return Pipe{x: startpipe.x - 1, y: startpipe.y, pipe: pipeplot[startpipe.y][startpipe.x-1]}
		}
	}

	return Pipe{}
}

func containsEnding(pipechar string, direction Direction) bool {
	pipeendings := PIPEMAPPING[pipechar]

	for _, ending := range pipeendings {
		if ending == direction {
			return true
		}
	}

	return false
}
