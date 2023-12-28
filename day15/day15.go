package day15

import (
	"fmt"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day15/input")
	return part1(input), part2(input)
}

/*
 * HASH - Turn chars into a single number between 0-255.
 * Start with current value as 0.
 * Determine ASCII Code for the current char in the string.
 * Increase current value by ASCII code above.
 * Current Value = Current Value * 17
 * Current Value = Current Value % 256
 *
 */
func part1(input []string) string {
	strs := strings.Split(input[0], ",")
	result := 0
	for _, str := range strs {
		result += hash(str)
	}
	return fmt.Sprint(result)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

func hash(s string) int {
	currentVal := 0
	for _, char := range s {
		currentVal += int(char)
		currentVal *= 17
		currentVal %= 256
	}

	return currentVal
}
