package day11

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay11Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "374" {
		t.Fatal("Day 11 - Part 1 output should be 374")
	}
}

func TestDay11Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input, 2) != "374" {
		t.Fatal("Day 11 - Part 2 output should be 374")
	}
	if part2(input, 10) != "1030" {
		t.Fatal("Day 11 - Part 2 output should be 1030")
	}

	if part2(input, 100) != "8410" {
		t.Fatal("Day 11 - Part 2 output should be 8410")
	}
}

func TestGetDistancesBetweenPoints(t *testing.T) {
	a := Point{x: 1, y: 6}
	b := Point{x: 5, y: 11}
	if getDistanceBetweenPoints(a, b) != 9 {
		t.Fatal("Day 11 - Distance calculation should be 9.")
	}

	a = Point{x: 5, y: 11}
	b = Point{x: 0, y: 11}
	if getDistanceBetweenPoints(a, b) != 5 {
		t.Fatal("Day 11 - Distance calculation should be 9.")
	}
}
