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
	fmtInput, err := formatInput(input)
	if err != nil {
		return 0
	}

	distance := 0
	for i := 0; i < len(fmtInput[0]); i++ {
		distance += int(math.Abs(float64(fmtInput[0][i]) - float64(fmtInput[1][i])))
	}

	return distance
}

func part2(input []string) int {
	fmtInput, err := formatInput(input)
	if err != nil {
		return 0
	}

	score := 0
	for _, leftNum := range fmtInput[0] {
		i := slices.Index(fmtInput[1], leftNum)
		if i == -1 {
			continue
		}

		nbOccurrences := 1
		for ; i+1 < len(fmtInput[1]); i++ {
			rightNum := fmtInput[1][i+1]
			if rightNum != leftNum {
				break
			}
			nbOccurrences++
		}
		score += leftNum * nbOccurrences
	}
	return score
}

func formatInput(input []string) ([2][]int, error) {
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

	slices.Sort(left)
	slices.Sort(right)

	return [2][]int{
		left, right,
	}, nil
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
