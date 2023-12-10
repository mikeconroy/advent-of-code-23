package day4

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay4Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	cards := inputToCards(input)
	if part1(cards) != "13" {
		t.Fatal("Day 4 - Part 1 output should be 13.")
	}
}

func TestDay4Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	cards := inputToCards(input)
	if part2(cards) != "0" {
		t.Fatal("Day X - Part 2 output should be 0")
	}
}
