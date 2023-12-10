package day09

import (
	"testing"

	"github.com/HSLU-Student/AoC-2023/util"
)

var INPUT = util.GetContentRoot("day09")

func TestPart1(t *testing.T) {

	expect := 1938800261
	got := Day09{}.Part1(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}

func TestPart2(t *testing.T) {

	expect := 1112
	got := Day09{}.Part2(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}
