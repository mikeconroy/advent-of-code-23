package day13

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay13Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "405" {
		t.Fatal("Day 13 - Part 1 output should be 405")
	}
}

func TestDay13Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "0" {
		t.Fatal("Day 13 - Part 2 output should be xxx")
	}
}
