package day14

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay14Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "136" {
		t.Fatal("Day 14 - Part 1 output should be 136")
	}
}

func TestDay14Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "0" {
		t.Fatal("Day 14 - Part 2 output should be xxx")
	}
}

func TestCalculateTotalLoad(t *testing.T) {
	plat := Platform{
		[]rune("OOOO.#.O.."),
		[]rune("OO..#....#"),
		[]rune("OO..O##..O"),
		[]rune("O..#.OO..."),
		[]rune("........#."),
		[]rune("..#....#.#"),
		[]rune("..O..#.O.O"),
		[]rune("..O......."),
		[]rune("#....###.."),
		[]rune("#....#...."),
	}
	actual := plat.CalculateTotalLoad()

	if actual != 136 {
		t.Fatal("Day 14 - Total Load is wrong. Should be 136. Got:", actual)
	}
}
