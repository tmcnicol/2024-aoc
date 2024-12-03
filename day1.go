package aoc

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func day1Part1(input string) int {
	data := readfile(input)
	left, right := parseInput(data)
	slices.Sort(left)
	slices.Sort(right)

	result := 0
	for i := 0; i < len(left); i++ {
		l, r := left[i], right[i]
		diff := l - r
		if diff < 0 {
			diff = r - l
		}
		result += diff
	}

	return result
}

func day1Part2(input string) int {
	data := readfile(input)
	left, right := parseInput(data)

	result := 0
	for _, v := range left {
		result += similarity(v, right)

	}
	return result
}

func similarity(key int, vals []int) int {
	instances := 0
	for _, v := range vals {
		if key == v {
			instances++
		}

	}
	return instances * key
}

func parseInput(data []string) ([]int, []int) {
	left, right := []int{}, []int{}

	for _, row := range data {
		values := strings.Fields(row)
		if len(values) != 2 {
			input := fmt.Sprintf("bad input: %v", values)
			panic(input)
		}
		l, r := strings.TrimSpace(values[0]), strings.TrimSpace(values[1])

		left = append(left, must(strconv.Atoi(l)))
		right = append(right, must(strconv.Atoi(r)))
	}

	return left, right
}
