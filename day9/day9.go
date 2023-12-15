package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day9/input")
	return part1(input), part2(input)
}

/*
 *	INPUT: Report of many values and how they are changing over time. Each line is history of a single value.
 *	OUTPUT: Prediction of next value in each history.
 * 	PROCESS: Make a ew sequence from the difference at each step of the history.
 *		If that sequence is not all zeroes, repeat this process, using the sequence you just generated as the input sequence.
 *		Once all of the values in your latest sequence are zeroes, you can extrapolate what the next value of the original history should be.
 *	0   3   6   9  12  15			0   3   6   9  12  15   B			0   3   6   9  12  15  18
 * 	  3   3   3   3   3		--> 	  3   3   3   3   3   A		-->		  3   3   3   3   3   3
 *		0   0   0   0					0   0   0   0   0   			 	0   0   0   0   0
 */
func part1(input []string) string {
	sequences := getSequences(input)
	result := 0
	for _, seq := range sequences {
		resultSeq := analyzeSequence(seq, false)
		result += resultSeq[len(resultSeq)-1]
	}
	return fmt.Sprint(result)
}

/*
 *	10  13  16  21  30  45		5  10  13  16  21  30  45
 *    3   3   5   9  15			  5   3   3   5   9  15
 *      0   2   4   6		-->		-2   0   2   4   6
 *     	  2   2   2					   2   2   2   2
 *        	0   0        				 0   0   0
 */
func part2(input []string) string {
	sequences := getSequences(input)
	result := 0
	for _, seq := range sequences {
		resultSeq := analyzeSequence(seq, true)
		result += resultSeq[0]
	}
	return fmt.Sprint(result)
}

// Takes a sequence
//
//	Create a new sequence with the values equal to the diff in each digit in the sequence.
//	If all values in the new sequene are 0 -> Add a 0 and return the sequence.
//	If not Pass the new Sequence to analyzeSequence and hold the response.
//	Take the last digit of the response and add it to the last digit of the new sequence calculated.
//	Return the newSequence with the new digit on the end.
func analyzeSequence(seq []int, backwards bool) []int {
	var newSeq []int

	prevVal := seq[0]
	allZero := true
	for i := 1; i < len(seq); i++ {
		diff := seq[i] - prevVal
		prevVal = seq[i]
		newSeq = append(newSeq, diff)
		if diff != 0 {
			allZero = false
		}
	}

	if !backwards {
		if allZero {
			return append(seq, prevVal)
		} else {
			analyzedSeq := analyzeSequence(newSeq, backwards)
			addedVal := prevVal + analyzedSeq[len(analyzedSeq)-1]
			newSeq = append(seq, addedVal)
			return newSeq
		}
	} else {
		if allZero {
			return append([]int{seq[0]}, seq...)
		} else {
			analyzedSeq := analyzeSequence(newSeq, backwards)
			addedVal := seq[0] - analyzedSeq[0]
			newSeq = append([]int{addedVal}, seq...)
			return newSeq
		}
	}
}

func getSequences(in []string) [][]int {
	var seqs [][]int
	for _, line := range in {
		var seq []int
		vals := strings.Split(line, " ")
		for _, valStr := range vals {
			val, _ := strconv.Atoi(valStr)
			seq = append(seq, val)
		}
		seqs = append(seqs, seq)
	}
	return seqs
}
