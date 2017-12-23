package main

import "testing"

func CompareFrequencyMaps(t *testing.T, actual, expected map[int]int) {
	for k, expectedCount := range expected {
		if actualCount, ok := actual[k]; !ok || actualCount != expectedCount {
			t.Logf("Number %v did not have expected frequency %v. Got %v instead...", k, expectedCount, actualCount)
			t.Fail()
		}
	}
}

func CompareNumberArrays(t *testing.T, actual, expected []int) {
	for i := range expected {
		if actual[i] != expected[i] {
			t.Logf("Expected %v and got %v at index %v", expected[i], actual[i], i)
			t.Fail()
		}
	}
}

func TestGetNumbers(t *testing.T) {
	nums := GetNumbers("123 124 125")
	expected := map[int]int{123: 1, 124: 1, 125: 1}
	if len(nums) < len(expected) {
		t.Log("GetNumbers() did not get the right count")
		t.Fail()
	}
	CompareFrequencyMaps(t, nums, expected)
}

func TestOrderMissingNumbers(t *testing.T) {
	arrayA := []int{123, 124, 125}
	setB := map[int]int{120: 1, 121: 1, 122: 2, 123: 1, 124: 2, 125: 1, 126: 1}
	expected := []int{120, 121, 122, 122, 124, 126}
	actual := OrderMissingNumbers(arrayA, setB)
	if len(actual) != len(expected) {
		t.Logf("OrderMissingNumbers() returned %v numbers, but expected %v", len(actual), len(expected))
		t.Fail()
	}
	CompareNumberArrays(t, actual, expected)
}

func TestParseNumbers(t *testing.T) {
	nums := ParseNumbers("123 456 100001")
	expected := []int{123, 456, 100001}
	if len(nums) != len(expected) {
		t.Log("ParseNumbers() did not get the right count")
		t.Fail()
	}
	CompareNumberArrays(t, nums, expected)
}

func TestCountNumbers(t *testing.T) {
	counts := CountNumbers([]int{123, 124, 125, 125})
	expected := map[int]int{123: 1, 124: 1, 125: 2}
	if len(counts) != len(expected) {
		t.Log("CountNumbers() did not return the expected length of counts")
		t.Fail()
	}
	CompareFrequencyMaps(t, counts, expected)
}
