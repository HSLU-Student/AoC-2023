package day07

import (
	"testing"

	"github.com/HSLU-Student/AoC-2023/util"
)

var INPUT = util.GetContentRoot("day07")

func TestPart1(t *testing.T) {

	expect := 251545216
	got := Day07{}.Part1(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}

func TestPart2(t *testing.T) {

	expect := 250384185
	got := Day07{}.Part2(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}
