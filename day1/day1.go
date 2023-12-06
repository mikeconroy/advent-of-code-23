package day1

import (
	"fmt"

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
	return "Part 2"
}

func isInt(val rune) bool {
	if val >= 48 && val <= 58 {
		return true
	}
	return false
}
