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
