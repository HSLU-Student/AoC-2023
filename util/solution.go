package util

import (
	"fmt"
	"time"
)

type Solution string

func NewSolution(result any, riddleNo int, took time.Duration) Solution {
	body := `
-------------------------------------------
The solution of todays puzzle #%v is: %v
Time elapsed: %v
-------------------------------------------
`
	res := fmt.Sprintf(body, riddleNo, result, took.Milliseconds())
	return Solution(res)
}
