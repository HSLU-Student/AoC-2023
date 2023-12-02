package day02

import (
	"reflect"
	"testing"
)

// Helper function testing
func TestParseGame(t *testing.T) {
	expect := [][]string{{"3", "blue"}, {"10", "red"}, {"2", "green"}, {"2", "blue"}, {"8", "red"}}
	got := ParseGame("Game 90: 3 blue, 10 red, 2 green; 2 blue; 8 red")

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected: %v, got: %v", expect, got)
	}
}
