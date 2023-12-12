package day5

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay5Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "35" {
		t.Fatal("Day 5 - Part 1 output should be 35")
	}
}

func TestDay5Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "46" {
		t.Fatal("Day 5 - Part 2 output should be 46")
	}
}
