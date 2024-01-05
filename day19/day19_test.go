package day19

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay19Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "19114" {
		t.Fatal("Day 19 - Part 1 output should be 19114")
	}
}

func TestDay19Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "0" {
		t.Fatal("Day 19 - Part 2 output should be xxx")
	}
}
