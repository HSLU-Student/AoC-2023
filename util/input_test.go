package util

import (
	"reflect"
	"testing"
)

func TestSplitContentLine(t *testing.T) {
	input := `This
is
a
multiline string`

	expect := []string{"This", "is", "a", "multiline string"}
	got := SplitContentLine(input)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected: %v, got: %v", expect, got)
	}
}

func TestSplitContentRow(t *testing.T) {
	input := `This
is
a
string`

	expect := []string{"Tias", "hst", "ir", "si", "n", "g"}
	got := SplitContentRow(input)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected: %v, got: %v", expect, got)
	}

}
