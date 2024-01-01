package day17

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay17Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "102" {
		t.Fatal("Day 17 - Part 1 output should be 102")
	}
}

func TestDay17Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "0" {
		t.Fatal("Day 17 - Part 2 output should be xxx")
	}
}
