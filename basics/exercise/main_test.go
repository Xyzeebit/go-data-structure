package main

import (
	"reflect"
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

func TestBinarySearch(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var find int = 5
	var exp int = 4
	var result int = BinarySearch(find, data)

	if exp != result {
		t.Logf("Expected: %d\n", exp)
		t.Logf("Got: %d\n", result)
		t.Error("Result does not match expectation")
	}

	find = 9
	exp = 8
	result = BinarySearch(find, data)

	if exp != result {
		t.Logf("Expected: %d\n", exp)
		t.Logf("Got: %d\n", result)
		t.Error("Result does not match expectation")
	}

	find = 1
	exp = 0
	result = BinarySearch(find, data)

	if exp != result {
		t.Logf("Expected: %d\n", exp)
		t.Logf("Got: %d\n", result)
		t.Error("Result does not match expectation")
	}

	find = 10
	exp = -1
	result = BinarySearch(find, data)

	if exp != result {
		t.Logf("Expected: %d\n", exp)
		t.Logf("Got: %d\n", result)
		t.Error("Result does not match expectation")
	}

	for i, v := range data {
		n := BinarySearch(v, data)
		if i != n {
			t.Error("Expected:", i, "Got:", n, "\n")
		}
	}
}

func TestRotateList(t *testing.T) {
	data := []int{10, 20, 30, 40, 50, 60}
	k := 2
	exp := []int{30, 40, 50, 60, 10, 20}
	result := RotateList(k, data)

	if !reflect.DeepEqual(exp, result) {
		t.Logf("Expected: %v\n", exp)
		t.Logf("Got: %v\n", result)
		t.Error("Result does not match expectation")
	}
}

func TestBinaryRecursiveSearch(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range data {
		n := BinaryRecursiveSearch(v, data)
		if i != n {
			t.Logf("Expected: %d\n", i)
			t.Logf("Got: %d\n", n)
			t.Error("Result does not match expectation")
		}
	}
}

func TestSum2DimensionalArray(t *testing.T) {
	data := [2][2]int{
		{1, 2},
		{5, 6},
	}
	exp := 14
	result := Sum2DimensionalArray(data)
	if exp != result {
		t.Logf("Expected: %d\n", exp)
		t.Logf("Got: %d\n", result)
		t.Error("Result does not match expectation")
	}

}

func TestSecondLargestInList(t *testing.T) {
	data := []int{1, 6, 7, 3, 9, 2}
	exp := 7
	result := SecondLargestInList(data)

	if exp != result {
		t.Logf("Expected: %d\n", exp)
		t.Logf("Got: %d\n", result)
		t.Error("Result doen not match expectation")
	}
}

func TestSum(t *testing.T) {
	v := 5
	exp := 15
	r := Sum(v)

	if exp != r {
		t.Logf("Expected: %d\n", exp)
		t.Logf("Got: %d\n", r)
		t.Error("Result does not match expectation")
	}
}

func TestReverseList(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	exp := []int{5, 4, 3, 2, 1}
	result := ReverseList(data)

	if !reflect.DeepEqual(exp, result) {
		t.Logf("Expected: %v\n", exp)
		t.Logf("Got: %v\n", result)
		t.Error("Result does not match expectation")
	}
}

func TestSort01(t *testing.T) {
	data := []int{1, 0, 1, 0, 0, 1, 0, 1, 1}
	exp := []int{0, 0, 0, 0, 1, 1, 1, 1, 1}
	r := Sort01(data)

	if !reflect.DeepEqual(exp, r) {
		t.Logf("Expected %v\n", exp)
		t.Logf("Got: %v\n", r)
		t.Error("Result does not match expectation")
	}

	data = []int{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0}
	exp = []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}
	r = Sort01(data)

	if !reflect.DeepEqual(exp, r) {
		t.Logf("Expected: %v\n", exp)
		t.Logf("Got: %v\n", r)
		t.Error("Result does not match expectation")
	}
}

func TestFindSumInArray(t *testing.T) {
	data := []int{2, 3, 5, 1, 6, 9}
	v := 3
	exp := []int{2, 1}
	r := FindSumInArray(v, data)

	if !reflect.DeepEqual(exp, r) {
		t.Logf("Expected: %v\n", exp)
		t.Logf("Got: %v\n", r)
		t.Error("Result does not match expectation")
	}
}
