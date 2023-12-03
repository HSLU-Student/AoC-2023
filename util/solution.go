package util

import (
	"fmt"
	"time"
)

type Solution struct {
	Result   any
	RiddleNo int
	Duration time.Duration
}

func (s Solution) ToString() string {
	body := `
-------------------------------------------
The solution of todays puzzle #%v is: %v
Time elapsed: %v
-------------------------------------------
`
	res := fmt.Sprintf(body, s.RiddleNo, s.Result, s.Duration)
	return res
}

func NewSolution(res any, riddleNo int, duration time.Duration) Solution {
	return Solution{res, riddleNo, duration}
}

/* func NewSolution(result any, riddleNo int, took time.Duration) Solution {
	body := `
-------------------------------------------
The solution of todays puzzle #%v is: %v
Time elapsed: %v
-------------------------------------------
`
	res := fmt.Sprintf(body, riddleNo, result, took)
	return Solution(res)
} */
