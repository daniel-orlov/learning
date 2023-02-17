package main

import "testing"

func TestMultiply(t *testing.T) {
	var v int
	v = multiply([]int{1, 4, 8, 6}...)
	exp := 192
	if v != exp {
		t.Errorf("Expected: %v\nGot: %v\n", exp, v)
	}
	var w int
	input := []int{1, 4, 8, 0}
	w = multiply(input...)
	exp = 0
	if w != exp {
		t.Errorf("Expected: %v\nGot: %v\n", exp, w)
	}
}
