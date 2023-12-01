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
