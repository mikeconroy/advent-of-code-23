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
	hands = sort(hands)

	result := 0

	for rank, hand := range hands {
		result += (rank + 1) * hand.bid
	}
	return fmt.Sprint(result)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

func loadInput(in []string) []Hand {
	var hands []Hand
	for _, line := range in {
		bid, _ := strconv.Atoi(strings.Split(line, " ")[1])
		cardsStr := strings.Split(line, " ")[0]
		var cards []int
		for _, cardStr := range cardsStr {
			var val int
			switch cardStr {
			case 'A':
				val = 14
			case 'K':
				val = 13
			case 'Q':
				val = 12
			case 'J':
				val = 11
			case 'T':
				val = 10
			default:
				val = int(cardStr - '0')
			}
			cards = append(cards, val)
		}
		newHand := Hand{
			cards: cards,
			bid:   bid,
		}

		newHand.hType = newHand.getType()
		hands = append(hands, newHand)
	}
	return hands
}
