package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day7/input")
	return part1(input), part2(input)
}

/*
 *	Camel Cards - Get a list of hands. Goal is to order based on strength.
 *	A Hand = 5 cards. A is highest. 2 is lowest.
 *	Strongest to Weakest Hands:
 *		Five of a Kind
 *		Four of a Kind
 *		Full House
 *		Three of a Kind
 *		Two Pair
 *		One Pair
 *		High Card
 *	If two hands have the same type, a second ordering rule takes effect.
 *	Start by comparing the first card in each hand. If these cards are different,
 *	the hand with the stronger first card is considered stronger.
 *	If the first card in each hand have the same label, however,
 *	then move on to considering the second card in each hand etc.
 *
 *	Input is a list of hands and their bid.
 *	Each hand wins it's bid * it's rank
 *	Weakest hand is rank 1, 2nd Weakest is rank 2 etc.
 *	Answer is the sum of the wins
 */
func part1(input []string) string {
	hands := loadInput(input)
	for i, hand := range hands {
		hands[i].hType = hand.getType()
	}
	hands = sort(hands)

	result := 0

	for rank, hand := range hands {
		result += (rank + 1) * hand.bid
	}
	return fmt.Sprint(result)
}

var valMap = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'1': 1,
}

// 249114596 is too low
// 247514536 is too low
// 249837714 is too low
// 250807190 is wrong
// 250506580
func part2(input []string) string {
	valMap['J'] = 0
	hands := loadInput(input)

	for i, hand := range hands {
		// The getType now needs to workout the best type by making Js lowest value & a wildcard.
		hands[i].hType = hand.getTypeWithJoker()
	}

	hands = sort(hands)
	result := 0

	for rank, hand := range hands {
		result += (rank + 1) * hand.bid
	}
	return fmt.Sprint(result)
}

func loadInput(in []string) []Hand {
	var hands []Hand
	for _, line := range in {
		bid, _ := strconv.Atoi(strings.Split(line, " ")[1])
		cardsStr := strings.Split(line, " ")[0]
		var cards []int
		for _, cardChar := range cardsStr {
			cards = append(cards, valMap[cardChar])
		}
		newHand := Hand{
			cards: cards,
			bid:   bid,
		}

		hands = append(hands, newHand)
	}
	return hands
}
