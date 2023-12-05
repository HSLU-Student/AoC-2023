package day05

import (
	"reflect"
	"testing"

	"github.com/HSLU-Student/AoC-2023/util"
)

var INPUT = util.GetContentRoot("day05")

func TestPart1(t *testing.T) {
	expect := 35
	got := Day05{}.Part1(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got)
	}
}

// Helper testing
func TestRemapToDest(t *testing.T) {
	seedsmap := map[int]int{
		79: 79,
		14: 14,
		55: 55,
		13: 13,
	}
	soilremap := [][]int{{50, 98, 2}, {52, 50, 48}}
	expect := map[int]int{
		79: 81,
		14: 14,
		55: 57,
		13: 13,
	}

	got := remapToDest(seedsmap, soilremap)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected: %v, got: %v", expect, got)
	}
}
