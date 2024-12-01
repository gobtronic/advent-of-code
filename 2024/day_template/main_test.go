package main

import "testing"

func TestPart1(t *testing.T) {
	input := parseInput("sample.txt")
	expected := 999

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
