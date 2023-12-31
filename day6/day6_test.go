package day6

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay6Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "288" {
		t.Fatal("Day 6 - Part 1 output should be 288")
	}
}

func TestDay6Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "71503" {
		t.Fatal("Day 6 - Part 2 output should be 71503")
	}
}
