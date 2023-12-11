package day10

import (
	"testing"

	"github.com/HSLU-Student/AoC-2023/util"
)

var INPUT = util.GetContentRoot("day10")

func TestPart1(t *testing.T) {

	expect := 6907
	got := Day10{}.Part1(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}

func TestPart2(t *testing.T) {

	expect := 1112
	got := Day10{}.Part2(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}
