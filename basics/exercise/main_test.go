package main

import (
	"testing"
)

func TestArraySum(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := 15
	result := ArraySum(data)

	if expected != result {
		t.Logf("Expected: %d\n", expected)
		t.Logf("Got: %d\n", result)
		t.Error("result does not match expectation")
	}
}

func TestSearchList(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := 2
	v := 3
	result := SearchList(v, data)

	if expected != result {
		t.Logf("Expected: %d\n", expected)
		t.Logf("Result: %d\n", result)
		t.Error("Result does not match expectation")
	}
}
