package day15

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay15Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "1320" {
		t.Fatal("Day 15 - Part 1 output should be 1320")
	}
}

func TestDay15Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "0" {
		t.Fatal("Day 15 - Part 2 output should be xxx")
	}
}
