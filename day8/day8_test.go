package day8

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay8Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test_1")
	if part1(input) != "2" {
		t.Fatal("Day 8 - Part 1 output should be 2")
	}

	input = utils.ReadFileIntoSlice("input_test_2")
	if part1(input) != "6" {
		t.Fatal("Day 8 - Part 1 output should be 6")
	}
}

func TestDay8Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test_3")
	if part2(input) != "6" {
		t.Fatal("Day 8 - Part 2 output should be 6")
	}
}
