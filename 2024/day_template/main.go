package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := parseInput("input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	return 0
}

func part2(input []string) int {
	return 0
}

func parseInput(fileName string) []string {
	lines := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return lines
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
