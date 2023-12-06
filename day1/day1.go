package day1

import (
	"fmt"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day1/input")
	return part1(input), part2(input)
}

func part1(input []string) string {
	total := 0
	for _, line := range input {
		firstNum, lastNum := -1, -1
		for _, value := range line {
			if isInt(value) {
				if firstNum == -1 {
					firstNum = int(value - '0')
				}
				lastNum = int(value - '0')
			}
		}
		total += (firstNum * 10) + lastNum
	}

	return fmt.Sprint(total)
}

func part2(input []string) string {

	numbers := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	total := 0
	for _, line := range input {
		firstNum, lastNum := -1, -1
		firstIdx, lastIdx := -1, -1
		for idx, value := range line {
			if isInt(value) {
				if firstNum == -1 {
					firstNum = int(value - '0')
					firstIdx = idx
				}
				lastNum = int(value - '0')
				lastIdx = idx
			}
		}

		for val, number := range numbers {
			firstIndex := strings.Index(line, number)
			if firstIndex != -1 {
				if firstIndex < firstIdx || firstIdx == -1 {
					firstNum = val + 1
					firstIdx = firstIndex
				}
			}
			lastIndex := strings.LastIndex(line, number)
			if lastIndex != -1 {
				if lastIndex > lastIdx {
					lastNum = val + 1
					lastIdx = lastIndex
				}
			}
		}

		total += (firstNum * 10) + lastNum
	}

	return fmt.Sprint(total)
}

func isInt(val rune) bool {
	if val >= 48 && val <= 58 {
		return true
	}
	return false
}
