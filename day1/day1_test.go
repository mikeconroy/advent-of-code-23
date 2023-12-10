package day1

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay1Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "142" {
		t.Fatal()
	}
}

func TestDay2Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test_2")
	if part2(input) != "281" {
		t.Fatal()
	}
}
