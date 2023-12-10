package day2

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay2Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")

	if part1(input) != "8" {
		t.Fatal("Day 2 - Part 1 output should be 8")
	}
}

func TestDay2Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "2286" {
		t.Fatal()
	}
}
