package day09

import (
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day09 struct{}

func (d Day09) Part1(input string) util.Solution {
	starttime := time.Now()
	total := 0
	for _, series := range util.SplitContentLine(input) {
		seriesint := util.ParseNumbers(series)
		total += extrapoliteRight(seriesint)
	}
	return util.NewSolution(total, 1, time.Since(starttime))
}

func (d Day09) Part2(input string) util.Solution {
	starttime := time.Now()
	total := 0
	for _, series := range util.SplitContentLine(input) {
		seriesint := util.ParseNumbers(series)
		total += extrapoliteLeft(seriesint)
	}
	return util.NewSolution(total, 2, time.Since(starttime))
}

func buildExpTbl(series []int) [][]int {
	seriesexp := [][]int{series}

	//build extrapolation table
	isarithmetic := false
	for nextseries := 0; !isarithmetic; nextseries++ {

		isarithmetic = true

		newseries := []int{}
		arithmetic := map[int]struct{}{}

		for idx := 0; idx < len(seriesexp[nextseries])-1; idx++ {
			num := (seriesexp[nextseries][idx+1] - seriesexp[nextseries][idx])

			newseries = append(newseries, num)
			arithmetic[num] = struct{}{}

			//check if we can stop processing
			if !isarithmetic || len(arithmetic) > 1 {
				isarithmetic = false
			}
		}
		seriesexp = append(seriesexp, newseries)
	}
	return seriesexp
}

func extrapoliteRight(series []int) int {
	seriesexp := buildExpTbl(series)

	//solve extrapolation table
	prev := 0
	for idx := len(seriesexp) - 1; idx > -1; idx-- {
		prev += seriesexp[idx][len(seriesexp[idx])-1]
	}
	return prev
}

func extrapoliteLeft(series []int) int {
	seriesexp := buildExpTbl(series)

	//solve extrapolation table
	prev := 0
	for idx := len(seriesexp) - 1; idx > -1; idx-- {
		prev = seriesexp[idx][0] - prev
	}
	return prev

}
