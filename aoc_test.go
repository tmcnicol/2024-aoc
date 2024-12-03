package aoc

import (
	"testing"

	"github.com/matryer/is"
)

func TestDay1(t *testing.T) {
	t.Run("pt1 - example", func(t *testing.T) {
		is := is.New(t)
		result := day1Part1("./data/day1_example.txt")
		is.Equal(11, result)
	})

	t.Run("pt1 input", func(t *testing.T) {
		is := is.New(t)
		result := day1Part1("./data/day1_input.txt")
		is.Equal(1830467, result)
	})

	t.Run("pt2 example", func(t *testing.T) {
		is := is.New(t)
		result := day1Part2("./data/day1_example.txt")
		is.Equal(31, result)
	})

	t.Run("pt2 input", func(t *testing.T) {
		is := is.New(t)
		result := day1Part2("./data/day1_input.txt")
		is.Equal(26674158, result)
	})
}

func TestDay2(t *testing.T) {
	t.Run("pt1 example", func(t *testing.T) {
		is := is.New(t)
		result := day2Part1("./data/day2_example.txt")
		is.Equal(2, result)
	})

	t.Run("pt1 input", func(t *testing.T) {
		is := is.New(t)
		result := day2Part1("./data/day2_input.txt")
		is.Equal(356, result)
	})

	t.Run("pt2 example", func(t *testing.T) {
		is := is.New(t)
		result := day2Part2("./data/day2_example.txt")
		is.Equal(4, result)
	})
	//
	t.Run("pt2 input", func(t *testing.T) {
		is := is.New(t)
		result := day2Part2("./data/day2_input.txt")
		is.Equal(413, result)
	})
}
