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
	if part2(input) != "2" {
		t.Fatal("Day 9 - Part 2 output should be 2")
	}
}

func TestDay9AnalyzeSequence(t *testing.T) {
	actual := analyzeSequence([]int{0, 3, 6, 9, 12, 15}, false)
	expected := []int{0, 3, 6, 9, 12, 15, 18}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Fatal("Day 9 - Analyze Sequence Output is wrong. Expected:", expected, "Actual:", actual)
		}
	}

	actual = analyzeSequence([]int{10, 13, 16, 21, 30, 45}, true)
	expected = []int{5, 10, 13, 16, 21, 30, 45}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Fatal("Day 9 - Analyze Sequence Output is wrong. Expected:", expected, "Actual:", actual)
		}
	}

	actual = analyzeSequence([]int{0, 3, 6, 9, 12, 15}, true)
	expected = []int{-3, 0, 3, 6, 9, 12, 15}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Fatal("Day 9 - Analyze Sequence Output is wrong. Expected:", expected, "Actual:", actual)
		}
	}

	actual = analyzeSequence([]int{1, 3, 6, 10, 15, 21}, true)
	expected = []int{0, 1, 3, 6, 10, 15, 21}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Fatal("Day 9 - Analyze Sequence Output is wrong. Expected:", expected, "Actual:", actual)
		}
	}
}
