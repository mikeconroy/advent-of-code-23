package day18

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay18Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "62" {
		t.Fatal("Day 18 - Part 1 output should be 62")
	}
}

func TestDay18Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "0" {
		t.Fatal("Day 18 - Part 2 output should be xxx")
	}
}
