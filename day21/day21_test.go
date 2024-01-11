package day21

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay21Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input, 6, 11) != "16" {
		t.Fatal("Day 21 - Part 1 output should be 16")
	}
}

// func TestDay21Part2(t *testing.T) {
// 	input := utils.ReadFileIntoSlice("input_test")
// 	if part2(input, 11) != "0" {
// 		t.Fatal("Day 21 - Part 2 output should be xxx")
// 	}
// }
