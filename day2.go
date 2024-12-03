package aoc

import (
	"strconv"
	"strings"
)

func day2Part1(input string) int {
	data := readfile(input)
	tests := parseRows(data)

	result := 0
	for _, test := range tests {
		if checkSafe(test) {
			result++
		}
	}

	return result
}

func day2Part2(input string) int {
	data := readfile(input)
	tests := parseRows(data)

	result := 0
	for _, test := range tests {
		if modifyAndCheckSafe(test) {
			result++
		}
	}

	return result
}

func modifyAndCheckSafe(test []int) bool {
	if checkSafe(test) {
		return true
	}

	for i := 0; i < len(test); i++ {
		modified := []int{}
		modified = append(modified, test[:i]...)
		modified = append(modified, test[i+1:]...)
		if checkSafe(modified) {
			return true
		}
	}
	return false
}

func checkSafe(test []int) bool {
	accending := true
	decending := true

	for i := 0; i < len(test)-1; i++ {
		diff := test[i] - test[i+1]
		if diff == 0 {
			return false
		}
		if diff < 0 {
			accending = false
		}
		if diff > 0 {
			decending = false
		}

		if diff > 3 || diff < -3 {
			return false
		}
	}

	if !accending && !decending {
		return false
	}

	return true
}

func parseRows(data []string) [][]int {
	result := [][]int{}

	for _, row := range data {
		result = append(result, parseRow(row))
	}

	return result
}

func parseRow(raw string) []int {
	fields := strings.Fields(raw)
	row := []int{}
	for _, f := range fields {
		row = append(row, must(strconv.Atoi(f)))
	}

	return row
}
