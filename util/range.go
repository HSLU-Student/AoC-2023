package util

import "errors"

func Range(start, end int) ([]int, error) {
	if start > end {
		return []int{}, errors.New("start index can't be greater than end index")
	}

	if start == end {
		return []int{start}, nil
	}

	rng := []int{}
	for i := start; i < end; i++ {
		rng = append(rng, i)
	}
	return rng, nil
}
