package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := formatInput()
	if err != nil {
		return
	}
	sort.Slice(input[0], func(i, j int) bool {
		return input[0][i] < input[0][j]
	})
	sort.Slice(input[1], func(i, j int) bool {
		return input[1][i] < input[1][j]
	})

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input [2][]int) int {
	distance := 0
	for i := 0; i < len(input[0]); i++ {
		distance += int(
			math.Abs(float64(input[1][i]) - float64(input[0][i])))
	}

	return distance
}

func part2(input [2][]int) int {
	score := 0
	for _, leftNum := range input[0] {
		i := slices.Index(input[1], leftNum)
		if i == -1 {
			continue
		}

		nbOccurrences := 1
		for ; i+1 < len(input[1]); i++ {
			rightNum := input[1][i+1]
			if rightNum != leftNum {
				break
			}
			nbOccurrences++
		}
		score += leftNum * nbOccurrences
	}
	return score
}

func formatInput() ([2][]int, error) {
	input := parseInput()
	left := []int{}
	right := []int{}

	for _, line := range input {
		parts := strings.Split(line, "   ")
		leftPart, leftErr := strconv.Atoi(parts[0])
		rightPart, rightErr := strconv.Atoi(parts[1])
		if leftErr != nil || rightErr != nil {
			return [2][]int{}, errors.New("error during strconv.Atoi")
		}
		left = append(left, leftPart)
		right = append(right, rightPart)
	}

	return [2][]int{
		left, right,
	}, nil
}

func parseInput() []string {
	lines := []string{}
	file, err := os.Open("input.txt")
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
