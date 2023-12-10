package dayX

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDayXPart1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "0" {
		t.Fatal("Day X - Part 1 output should be xxx")
	}
}

func TestDayXPart2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "0" {
		t.Fatal("Day X - Part 2 output should be xxx")
	}
}
