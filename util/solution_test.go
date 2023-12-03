package util

import (
	"testing"
	"time"
)

func TestNewSolution(t *testing.T) {
	t.Run("Test new int solution", func(t *testing.T) {
		sol := NewSolution(420, 1, time.Duration(10000000))

		body := `
-------------------------------------------
The solution of todays puzzle #1 is: 420
Time elapsed: 10ms
-------------------------------------------
`
		assertSolution(t, body, string(sol.ToString()))
	})

	t.Run("Test new string Solution", func(t *testing.T) {
		sol := NewSolution("Answer", 1, time.Duration(10000000))

		body := `
-------------------------------------------
The solution of todays puzzle #1 is: Answer
Time elapsed: 10ms
-------------------------------------------
`
		assertSolution(t, body, string(sol.ToString()))
	})
}

func assertSolution(t testing.TB, expect, got string) {
	t.Helper()
	if expect != got {
		t.Errorf("Expected %v, got %v", expect, got)
	}
}
