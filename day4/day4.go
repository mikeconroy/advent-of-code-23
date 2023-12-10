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
		var matchingNumberCount int
		for _, myNum := range card.myNums {
			if card.winMap[myNum] == true {
				matchingNumberCount++
			}
		}
		// 1 Matching Number = 1 Point		2^0
		// 2 Matching Numbers = 2 Points	2^1
		// 3 Matching Numbers = 4 Points	2^2
		// 2^(matchingNumbers - 1)
		totalPoints += int(math.Pow(2, float64((matchingNumberCount - 1))))
	}
	return fmt.Sprint(totalPoints)
}

func part2(cards []ScratchCard) string {
	return fmt.Sprint(0)
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
		num, _ := strconv.Atoi(val)
		out = append(out, num)
	}
	return out
}
