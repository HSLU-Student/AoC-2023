package day05

import (
	"fmt"
	"regexp"
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day05 struct{}

func (d Day05) Part1(input string) util.Solution {
	starttime := time.Now()
	reg := regexp.MustCompile(`^\n`)
	fmt.Println(reg.Split(input, -1))
	return util.NewSolution(1, 1, time.Since(starttime))
}

func (d Day05) Part2(input string) util.Solution {
	starttime := time.Now()
	res := "Not implemeted yet..."
	return util.NewSolution(res, 2, time.Since(starttime))
}

// note to me implement with copys of parameter to prevent modification of underlying parameter data stucture
func remapToDest(input map[int]int, remapper [][]int) map[int]int {
	//for each remap entry calculate last element included
	for idx, entry := range remapper {
		remapper[idx] = append(entry, (entry[1] + entry[2] - 1))
	}

	//iterate over input & find remap target
	for key, src := range input {
		for _, entry := range remapper {
			// find corresponing remap entry
			offset := src - entry[1]
			if offset > -1 && entry[1]+offset < entry[3] {
				input[key] = entry[0] + offset
				break
			}
		}
	}
	return input
}
