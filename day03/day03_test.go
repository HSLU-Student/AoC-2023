package day03

import (
	"testing"

	"github.com/HSLU-Student/AoC-2023/util"
)

var INPUT = util.GetContentRoot("day03")

func TestPart1(t *testing.T) {
	expect := 527369
	got := Day03{}.Part1(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}

func TestPart2(t *testing.T) {
	expect := 73074886
	got := Day03{}.Part2(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}
