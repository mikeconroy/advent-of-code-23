package day12

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay12Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "21" {
		t.Fatal("Day 12 - Part 1 output should be 21")
	}
}

func TestCountValidArrangements(t *testing.T) {
	field := Field{
		springs:       []rune{'?', '?', '?', '.', '#', '#', '#'},
		damagedGroups: []int{1, 1, 3},
		unknowns:      3,
	}
	actual := countValidArrangements(field, 0)
	if actual != 1 {
		t.Fatal("Day 12 - Count Valid Arrangements is not correct. Output should be 1. Instead got:", actual)
	}

	field = Field{
		springs:       []rune{'.', '?', '?', '.', '.', '?', '?', '.', '.', '.', '?', '#', '#', '.'},
		damagedGroups: []int{1, 1, 3},
		unknowns:      5,
	}
	actual = countValidArrangements(field, 0)
	if actual != 4 {
		t.Fatal("Day 12 - Count Valid Arrangements is not correct. Output should be 4. Instead got:", actual)
	}

	var springs []rune
	fieldStr := "???.###????.###????.###????.###????.###"
	for _, spring := range fieldStr {
		springs = append(springs, spring)
	}
	field = Field{
		springs:       springs,
		damagedGroups: []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3},
		unknowns:      19,
	}
	actual = countValidArrangements(field, 0)
	if actual != 1 {
		t.Fatal("Day 12 - Count Valid Arrangements is not correct. Output should be 1. Instead got:", actual)
	}

}

func TestDay12Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	actual := part2(input)
	if actual != "525152" {
		t.Fatal("Day 12 - Part 2 output should be 525152. Instead got:", actual)
	}
}
