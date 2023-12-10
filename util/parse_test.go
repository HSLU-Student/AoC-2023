package util

import (
	"reflect"
	"testing"
)

func TestParseNumbers(t *testing.T) {
	input := "Test: 12 23 34 45"

	expect := []int{12, 23, 34, 45}
	got := ParseNumbers(input)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected: %v, got: %v", expect, got)
	}
}

func TestParseNumbersNeg(t *testing.T) {
	input := "Test: -12 23 34 -45"

	expect := []int{-12, 23, 34, -45}
	got := ParseNumbers(input)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected: %v, got: %v", expect, got)
	}
}
