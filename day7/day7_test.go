package day7

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay7Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "6440" {
		t.Fatal("Day 7 - Part 1 output should be 6440")
	}
}

func TestDay7Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "0" {
		t.Fatal("Day 7 - Part 2 output should be xxx")
	}
}

func TestDay7GetType(t *testing.T) {
	hand := Hand{
		cards: []int{1, 2, 3, 4, 5},
		bid:   0,
	}
	if hand.getType() != HIGH_CARD {
		t.Fatal("Day 7 - High Card Hand Type is wrong.")
	}
	hand.cards = []int{1, 1, 2, 3, 4}
	if hand.getType() != ONE_PAIR {
		t.Fatal("Day 7 - One Pair Hand Type is wrong.")
	}
	hand.cards = []int{1, 1, 2, 3, 3}
	if hand.getType() != TWO_PAIR {
		t.Fatal("Day 7 - Two Pair Hand Type is wrong.")
	}
	hand.cards = []int{1, 2, 3, 3, 3}
	if hand.getType() != THREE_OF_A_KIND {
		t.Fatal("Day 7 - Three of a Kind Hand Type is wrong.")
	}
	hand.cards = []int{2, 2, 3, 3, 3}
	if hand.getType() != FULL_HOUSE {
		t.Fatal("Day 7 - Full House Hand Type is wrong.")
	}
	hand.cards = []int{2, 3, 3, 3, 3}
	if hand.getType() != FOUR_OF_A_KIND {
		t.Fatal("Day 7 - Four of a Kind Hand Type is wrong.")
	}
	hand.cards = []int{3, 3, 3, 3, 3}
	if hand.getType() != FIVE_OF_A_KIND {
		t.Fatal("Day 7 - Five of a Kind Hand Type is wrong.")
	}
}
