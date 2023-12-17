package day10

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay10Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "8" {
		t.Fatal("Day 10 - Part 1 output should be 8")
	}
}

func TestDay10Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test_2")
	if part2(input) != "4" {
		t.Fatal("Day 10 - Part 2 input_test_2 output should be 4")
	}

	input = utils.ReadFileIntoSlice("input_test_3")
	if part2(input) != "4" {
		t.Fatal("Day 10 - Part 2 input_test_3 output should be 4")
	}

	input = utils.ReadFileIntoSlice("input_test_4")
	if part2(input) != "8" {
		t.Fatal("Day 10 - Part 2 input_test_4 output should be 8")
	}

	input = utils.ReadFileIntoSlice("input_test_5")
	if part2(input) != "10" {
		t.Fatal("Day 10 - Part 2 input_test_5 output should be 10")
	}
}
