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
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "0" {
		t.Fatal("Day 10 - Part 2 output should be xxx")
	}
}
