package main

import (
	"slices"
	"testing"
)

func TestFormatInput(t *testing.T) {
	expected := [2][]int{
		[]int{
			1, 2, 3, 3, 3, 4,
		},
		[]int{
			3, 3, 3, 4, 5, 9,
		},
	}

	input, _ := formatInput("sample.txt")

	if slices.Compare(expected[0], input[0]) != 0 {
		t.Fatalf("Expected %d, got %d", expected[0], input[0])
	}
	if slices.Compare(expected[1], input[1]) != 0 {
		t.Fatalf("Expected %d, got %d", expected[1], input[1])
	}
}

func TestPart1(t *testing.T) {
	input, _ := formatInput("sample.txt")
	expected := 11

	res := part1(input)
	if res != expected {
		t.Fatalf("Expected %d, got %d", expected, res)
	}
}

func TestPart2(t *testing.T) {
	input, _ := formatInput("sample.txt")
	expected := 31

	res := part2(input)
	if res != expected {
		t.Fatalf("Expected %d, got %d", expected, res)
	}
}
