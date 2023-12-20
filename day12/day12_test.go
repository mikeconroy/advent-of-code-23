package day12

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay12Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "21" {
		t.Fatal("Day 12 - Part 1 output should be 21")
	}
}

func TestDay12Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "525152" {
		t.Fatal("Day 12 - Part 2 output should be 525152")
	}
}
