package day22

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay22Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "5" {
		t.Fatal("Day 22 - Part 1 output should be 5")
	}
}

func TestDay22Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "7" {
		t.Fatal("Day 22 - Part 2 output should be 7")
	}
}
