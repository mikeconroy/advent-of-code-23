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
		resultSeq := analyzeSequence(seq)
		result += resultSeq[len(resultSeq)-1]
	}
	return fmt.Sprint(result)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

// Takes a sequence
//
//	Create a new sequence with the values equal to the diff in each digit in the sequence.
//	If all values in the new sequene are 0 -> Add a 0 and return the sequence.
//	If not Pass the new Sequence to analyzeSequence and hold the response.
//	Take the last digit of the response and add it to the last digit of the new sequence calculated.
//	Return the newSequence with the new digit on the end.
func analyzeSequence(seq []int) []int {
	fmt.Println("Analyzing:", seq)
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

	if allZero {
		return append(seq, prevVal)
	} else {
		analyzedSeq := analyzeSequence(newSeq)
		addedVal := prevVal + analyzedSeq[len(analyzedSeq)-1]
		fmt.Println(seq, addedVal, analyzedSeq)
		newSeq = append(seq, addedVal)
		return newSeq
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
