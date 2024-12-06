package aoc

import (
	"regexp"
	"strconv"
)

const CMD = "mul"

func day3Part1(input string) int {
	data := readfile(input)
	raw := parseData(data)

	pattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := pattern.FindAllString(raw, -1)

	result := 0
	for _, match := range matches {
		op := parseCandidate(match)
		result += op.mul()
	}
	return result
}

func parseData(data []string) string {
	chars := []rune{}
	for _, row := range data {
		for _, v := range row {
			chars = append(chars, v)
		}
	}
	return string(chars)
}

type OP struct {
	arg1 int
	arg2 int
}

func (o OP) mul() int {
	return o.arg1 * o.arg2
}

func parseCandidate(candidate string) OP {
	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	match := pattern.FindStringSubmatch(candidate)
	arg1 := must(strconv.Atoi(match[1]))
	arg2 := must(strconv.Atoi(match[2]))

	return OP{
		arg1: arg1,
		arg2: arg2,
	}
}
