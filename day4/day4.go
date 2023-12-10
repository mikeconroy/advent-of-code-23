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
 *	Cards will never make you copy a card past the end of the table.
 */
func part2(cards []ScratchCard) string {
	copies := make(map[int]int)
	var toProcess Stack

	for id := range cards {
		toProcess.Push(id)
		copies[id] = copies[id] + 1
	}

	for len(toProcess) > 0 {
		cardId := toProcess.Pop()
		matchCount := matchCount(cards[cardId])
		for i := 1; i <= matchCount; i++ {
			copies[cardId+i] = copies[cardId+i] + 1
			toProcess.Push(cardId + i)
		}
	}

	total := 0
	for _, val := range copies {
		total += val
	}

	return fmt.Sprint(total)
}

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	l := len(*s)
	if l == 0 {
		return -1
	}
	poppedVal := (*s)[l-1]
	*s = (*s)[:l-1]
	return poppedVal
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
