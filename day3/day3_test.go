package day3

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay3Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	eng := LoadSchematic(input)
	if part1(eng) != "4361" {
		t.Fatal("Day 3 - Part 1 output should be 4361")
	}
}

func TestDay3Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	eng := LoadSchematic(input)
	if part2(eng) != "0" {
		t.Fatal("Day 3 - Part 2 output should be xxx")
	}
}

func TestLoadSchematic(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	eng := LoadSchematic(input)
	if len(eng.schematic) != 10 {
		t.Fatal("Load Schematic Failed - Wrong number of rows.")
	}
	if len(eng.schematic[0]) != 10 {
		t.Fatal("Load Schematic Failed - Wrong number of columns.")
	}
}

func TestFindSymbols(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	eng := LoadSchematic(input)
	symbols := eng.FindSymbols()
	if len(symbols) != 6 {
		t.Fatal("Finding Symbols failed. Wrong number of symbols found.")
	}
	if symbols[0].char != '*' && symbols[0].x != 3 && symbols[0].y != 1 {
		t.Fatal("Finding Symbols failed. The Symbol position/value is wrong.")
	}
}
