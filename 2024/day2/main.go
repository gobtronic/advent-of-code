package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	reports, err := formatInput(input)
	if err != nil {
		return 0
	}

	nbSafe := 0
	for _, levels := range reports {
		levelsOrdering := sliceOrdering(levels)
		if levelsOrdering == OrderingUnordered {
			continue
		}

		previousLevel := -999
		for i, level := range levels {
			if previousLevel == -999 {
				previousLevel = level
				continue
			}

			distance := math.Abs(float64(level) - float64(previousLevel))
			if !(distance >= 1 && distance <= 3) {
				break
			}

			if i == len(levels)-1 {
				nbSafe++
			} else {
				previousLevel = level
			}
		}
	}
	return nbSafe
}

const (
	OrderingReverseSorted = iota - 1
	OrderingUnordered
	OrderingSorted
)

func sliceOrdering(slice []int) int {
	if slices.IsSorted(slice) {
		return OrderingSorted
	}

	sliceCopy := append(make([]int, 0, len(slice)), slice...)
	slices.Sort(sliceCopy)
	slices.Reverse(sliceCopy)
	if slices.Equal(slice, sliceCopy) {
		return OrderingReverseSorted // Reverse sorted
	}

	return OrderingUnordered // Non sorted
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

func formatInput(input []string) ([][]int, error) {
	reports := [][]int{}
	for _, line := range input {
		levels := []int{}
		for _, part := range strings.Split(line, " ") {
			level, err := strconv.Atoi(part)
			if err != nil {
				return [][]int{}, errors.New("error during strconv.Atoi")
			}
			levels = append(levels, level)
		}
		reports = append(reports, levels)
	}
	return reports, nil
}
