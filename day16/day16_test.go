package day16

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay16Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "46" {
		t.Fatal("Day 16 - Part 1 output should be 46")
	}
}

func TestDay16Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "51" {
		t.Fatal("Day 16 - Part 2 output should be 51")
	}
}
