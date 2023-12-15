package day9

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay9Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "114" {
		t.Fatal("Day 9 - Part 1 output should be 114")
	}
}

func TestDay9Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "0" {
		t.Fatal("Day 9 - Part 2 output should be xxx")
	}
}

func TestDay9AnalyzeSequence(t *testing.T) {
	actual := analyzeSequence([]int{0, 3, 6, 9, 12, 15})
	expected := []int{0, 3, 6, 9, 12, 15, 18}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Fatal("Day 9 - Analyze Sequence Output is wrong. Expected:", expected, "Actual:", actual)
		}
	}
}
