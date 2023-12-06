package day06

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day06 struct{}

func (d Day06) Part1(input string) util.Solution {
	starttime := time.Now()
	inputs := util.ParseNumbers(input)

	total := 1
	for i := 0; i < len(inputs)/2; i++ {
		total *= numberOfWinningCharges(int64(inputs[i]), int64(inputs[i+len(inputs)/2]))
	}

	return util.NewSolution(total, 1, time.Since(starttime))
}

func (d Day06) Part2(input string) util.Solution {
	starttime := time.Now()

	inputs := strings.ReplaceAll(input, " ", "")

	reg := regexp.MustCompile(`\d+`)
	numbersstr := reg.FindAllString(inputs, -1)

	inputnumbers := []int64{}
	for _, numstr := range numbersstr {
		num, _ := strconv.ParseInt(numstr, 10, 64)
		inputnumbers = append(inputnumbers, num)
	}

	total := numberOfWinningCharges(inputnumbers[0], inputnumbers[1])
	return util.NewSolution(total, 2, time.Since(starttime))
}

func numberOfWinningCharges(ms, record int64) int {
	validcharges := 0
	for i := int64(0); i < ms; i++ {
		if (ms-i)*i > record {
			validcharges++
		}
	}

	return validcharges
}
