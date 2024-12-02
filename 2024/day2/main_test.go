package main

import (
	"slices"
	"testing"
)

func TestFormatInput(t *testing.T) {
	expected := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	input := parseInput("sample.txt")

	fmtInput, err := formatInput(input)

	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
	for i, levels := range fmtInput {
		if slices.Compare(expected[i], levels) != 0 {
			t.Fatalf("Expected %d, got %d", expected[i], levels)
		}
	}
}

func TestSliceOrdering(t *testing.T) {
	expected := []int{OrderingReverseSorted, OrderingUnordered, OrderingSorted, OrderingSorted}
	examples := [][]int{
		{3, 2, 1},
		{1, 3, 2},
		{1, 2, 2, 3},
		{1, 2, 3},
	}

	for i, example := range examples {
		ordering := sliceOrdering(example)
		if ordering != expected[i] {
			t.Fatalf("expected ordering %d, got %d for slice %v", expected[i], ordering, example)
		}
	}
}

func TestPart1(t *testing.T) {
	input := parseInput("sample.txt")
	expected := 2

	res := part1(input)
	if res != expected {
		t.Fatalf("Expected %d, got %d", expected, res)
	}
}

func TestPart2(t *testing.T) {
	input := parseInput("sample.txt")
	expected := 999

	res := part2(input)
	if res != expected {
		t.Fatalf("Expected %d, got %d", expected, res)
	}
}
