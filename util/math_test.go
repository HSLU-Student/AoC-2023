package util

import "testing"

func TestGcd(t *testing.T) {
	expect := 13
	got := Gcd(12389, 34567)

	if expect != int(got) {
		t.Errorf("Expect: %v, got: %v", expect, got)
	}
}

func TestLcm(t *testing.T) {
	expect := 1783130
	got := Lcm(1234, 1445)

	if expect != int(got) {
		t.Errorf("Expected: %v, got: %v", expect, got)
	}
}

func TestAbs(t *testing.T) {
	expect := 1456
	got := Abs(-1456)

	if expect != got {
		t.Errorf("Excepted: %v, got: %v", expect, got)
	}
}
