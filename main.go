package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/HSLU-Student/AoC-2023/day01"
	"github.com/HSLU-Student/AoC-2023/day02"
	"github.com/HSLU-Student/AoC-2023/day03"
	"github.com/HSLU-Student/AoC-2023/day04"
	"github.com/HSLU-Student/AoC-2023/day05"
	"github.com/HSLU-Student/AoC-2023/day06"
	"github.com/HSLU-Student/AoC-2023/day07"
	"github.com/HSLU-Student/AoC-2023/day08"
	"github.com/HSLU-Student/AoC-2023/day09"
	"github.com/HSLU-Student/AoC-2023/util"
)

const AOC = `
  __    __    ___    ____   __  ____  ____ 
 / _\  /  \  / __)  (___ \ /  \(___ \( __ \
/    \(  O )( (__    / __/(  0 )/ __/ (__ (
\_/\_/ \__/  \___)  (____) \__/(____)(____/
===========================================

`

func main() {
	fmt.Print(AOC)

	if len(os.Args[1:]) != 1 {
		fmt.Println("❌  Invalid number of arguments passed. Expected: 1 got:", len(os.Args[1:]))
		fmt.Println("Type `go run . <Day>` to execute the puzzle")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("❌  Invalid day argument passed.")
		return
	}

	//the puzzle registry with all puzzles
	puzzles := map[int]util.Day{
		1: day01.Day01{},
		2: day02.Day02{},
		3: day03.Day03{},
		4: day04.Day04{},
		5: day05.Day05{},
		6: day06.Day06{},
		7: day07.Day07{},
		8: day08.Day08{},
		9: day09.Day09{},
	}

	selectedPuzzle, exist := puzzles[day]

	if !exist {
		fmt.Println("⚠️  Solution not implemented.")
		return
	}

	runDay(day, selectedPuzzle)
}

func runDay(dayNo int, day util.Day) {
	input := fmt.Sprintf("day%02d", dayNo)
	solutions := []util.Solution{day.Part1(util.GetContent(input)), day.Part2(util.GetContent(input))}

	for _, s := range solutions {
		fmt.Print(s.ToString())
	}
}
