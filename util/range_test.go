package util

import (
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	expect := []int{2, 3, 4, 5}
	got, _ := Range(2, 6)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected,%v got: %v", expect, got)
	}
}

func TestRangeExpection(t *testing.T) {
	_, err := Range(3, 2)

	if err == nil {
		t.Errorf("Expected error but didn't get one")
	}
}

func TestRangeZero(t *testing.T) {
	expect := []int{2}
	got, _ := Range(2, 2)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected: %v, got: %v", expect, got)
	}
}
