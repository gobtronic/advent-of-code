package main

import (
	"slices"
	"testing"
)

func TestFormatInput(t *testing.T) {
	expected := [2][]int{
		{1, 2, 3, 3, 3, 4},
		{3, 3, 3, 4, 5, 9},
	}
	input := parseInput("sample.txt")

	fmtInput, err := formatInput(input)

	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
	if !slices.Equal(expected[0], fmtInput[0]) {
		t.Fatalf("Expected %d, got %d", expected[0], fmtInput[0])
	}
	if !slices.Equal(expected[1], fmtInput[1]) {
		t.Fatalf("Expected %d, got %d", expected[1], fmtInput[1])
	}
}

func TestPart1(t *testing.T) {
	input := parseInput("sample.txt")
	expected := 11

	res := part1(input)

	if res != expected {
		t.Fatalf("Expected %d, got %d", expected, res)
	}
}

func TestPart2(t *testing.T) {
	input := parseInput("sample.txt")
	expected := 31

	res := part2(input)

	if res != expected {
		t.Fatalf("Expected %d, got %d", expected, res)
	}
}
