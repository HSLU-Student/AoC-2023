package util

import (
	"regexp"
	"strconv"
)

func ParseNumbers(numstr string) []int {
	reg := regexp.MustCompile(`-?\d+`)
	numbersstr := reg.FindAllString(numstr, -1)

	numbers := []int{}
	for _, numstr := range numbersstr {
		numi, _ := strconv.Atoi(numstr)
		numbers = append(numbers, numi)
	}
	return numbers
}
