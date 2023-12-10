package day4

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day4/input")
	cards := inputToCards(input)
	return part1(cards), part2(cards)
}

type ScratchCard struct {
	myNums  []int
	winNums []int
	winMap  map[int]bool
}

func part1(cards []ScratchCard) string {
	var totalPoints int
	for _, card := range cards {
		matchingNumberCount := matchCount(card)

		// 1 Matching Number = 1 Point		2^0
		// 2 Matching Numbers = 2 Points	2^1
		// 3 Matching Numbers = 4 Points	2^2
		// 2^(matchingNumbers - 1)
		totalPoints += int(math.Pow(2, float64((matchingNumberCount - 1))))
	}
	return fmt.Sprint(totalPoints)
}

func matchCount(card ScratchCard) (matchingNumberCount int) {
	for _, myNum := range card.myNums {
		if card.winMap[myNum] == true {
			matchingNumberCount++
		}
	}
	return matchingNumberCount
}

/*
 *	The performance of this Part has been improved since first implementation.
 *  The main improvement was recognising that the Stack wasn't required. Previously, a stack was used to track cards yet to be processed,
 *  each copy of a card would be added to the stack and be processed. This meant each copy of a card was processed multiple times despite always giving the same results.
 *  Instead we can process the cards in order (1, 2, 3...) and multiply the counts added to future cards by the number of copies of the current card.
 *  This would work because we never add copies of previous cards.
 *  E.g. we would get to Card 3 and we may have 3 instances of Card 3 to process (1 original + 1 copy from Card 1 + 1 Copy from Card 2)
 *  Then if our matchCount was 4 we can add 3 copies of Cards 4, 5, 6 and 7 to the copies map.
 *  Previously we were just adding 1 copy and repeating the processing 3 times.
 */
func part2(cards []ScratchCard) string {
	copies := make(map[int]int)
	totalCards := len(cards)

	for id := range cards {
		copies[id] = copies[id] + 1
		matches := matchCount(cards[id])
		for i := 1; i <= matches; i++ {
			copies[id+i] = copies[id+i] + copies[id]
			totalCards += copies[id]
		}
	}

	return fmt.Sprint(totalCards)
}

func inputToCards(input []string) (cards []ScratchCard) {
	for _, card := range input {
		nums := strings.Split(card, " | ")
		mine := strings.Split(nums[1], " ")
		winning := strings.Split(strings.Split(nums[0], ": ")[1], " ")
		winMap := make(map[int]bool)
		for _, win := range winning {
			if win != "" {
				winNum, err := strconv.Atoi(win)
				if err != nil {
					fmt.Println("Error converting to int:", err, winNum, win)
					panic("Error converting to int.")
				}
				winMap[winNum] = true
			}
		}

		cards = append(cards, ScratchCard{
			myNums:  convStrToInts(mine),
			winNums: convStrToInts(winning),
			winMap:  winMap,
		})
	}
	return cards
}

func convStrToInts(in []string) (out []int) {
	for _, val := range in {
		if val != "" {
			num, _ := strconv.Atoi(val)
			out = append(out, num)
		}
	}
	return out
}
