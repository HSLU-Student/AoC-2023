package day02

import (
	"reflect"
	"testing"

	"github.com/HSLU-Student/AoC-2023/util"
)

var INPUT = util.GetContentRoot("day02")

func TestPart1(t *testing.T) {
	expect := 2593
	got := Day02{}.Part1(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}

func TestPart2(t *testing.T) {
	expect := 54699
	got := Day02{}.Part2(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}

// Helper function testing
func TestParseGame(t *testing.T) {
	expect := [][]string{{"3", "blue"}, {"10", "red"}, {"2", "green"}, {"2", "blue"}, {"8", "red"}}
	got := parseGame("Game 90: 3 blue, 10 red, 2 green; 2 blue; 8 red")

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected: %v, got: %v", expect, got)
	}
}
