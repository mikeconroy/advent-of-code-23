package day20

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay20Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "32000000" {
		t.Fatal("Day 20 - Part 1 output should be 32000000")
	}

	input = utils.ReadFileIntoSlice("input_test_2")
	if part1(input) != "11687500" {
		t.Fatal("Day 20 - Part 1 output should be 11687500")
	}
}

// No test data provided for Day 20 P2
// func TestDay20Part2(t *testing.T) {
// 	input := utils.ReadFileIntoSlice("input_test")
// 	if part2(input) != "0" {
// 		t.Fatal("Day 20 - Part 2 output should be xxx")
// 	}
// }
