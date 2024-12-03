package aoc

import (
	"bufio"
	"os"
)

func readfile(f string) []string {
	file := must(os.Open(f))
	defer file.Close()

	result := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}
