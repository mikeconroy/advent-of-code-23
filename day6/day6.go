package day6

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day6/input")
	return part1(input), part2(input)
}

/*
 *	Input: Time Allowed (ms) and Best Distance (mm) recorded in each race.
 *	To win must go further in each race.
 *  Holding the button charges the boat, releasing allows it to move.
 *	Boats move faster the longer the button was held.
 *	Can only hold button at start of the race.
 *	Boat starts at 0mm/ms. Each ms the button is held increases the ms/mm by 1.
 *	Output is the number of ways to win each race multiplied together.
 */
func part1(input []string) string {

	times, distances := loadInput(input)
	// Races can be represented by a quadratic equation:
	// y=-x^2 + (RaceTime * x)
	// Then find where x intercepts Y at target distance.
	// Integers between x1 & x2 is the number of winning options.
	result := 1
	for i, time := range times {
		a := float64(-1)
		b := float64(time)
		c := float64(-distances[i])
		low, high := solveQuadEquation(a, b, c)
		result *= countIntsBetweenRange(low, high)
	}

	return fmt.Sprint(result)
}

func part2(input []string) string {
	times, distances := loadInput(input)
	time := sliceToInt(times)
	distance := sliceToInt(distances)
	a := float64(-1)
	b := float64(time)
	c := float64(-distance)
	low, high := solveQuadEquation(a, b, c)
	result := countIntsBetweenRange(low, high)
	return fmt.Sprint(result)
}

func sliceToInt(in []int) int {
	str := ""
	for _, i := range in {
		str += strconv.Itoa(i)
	}
	result, _ := strconv.Atoi(str)
	return result
}

func countIntsBetweenRange(low, high float64) int {
	// We always want to round up for the low number to nearest int and down for the high number.
	floor := int(low) + 1
	ceil := int(high)
	if float64(ceil) == high {
		ceil--
	}
	return ceil - floor + 1
}

func solveQuadEquation(a, b, c float64) (float64, float64) {
	sqrt := math.Sqrt(math.Pow(b, 2) - (4 * a * c))
	x1 := (-b + sqrt) / (2 * a)
	x2 := (-b - sqrt) / (2 * a)
	return x1, x2
}

func loadInput(in []string) ([]int, []int) {
	timesStr := strings.Split(strings.TrimSpace(strings.Replace(in[0], "Time:", "", 1)), " ")
	distancesStr := strings.Split(strings.TrimSpace(strings.Replace(in[1], "Distance:", "", 1)), " ")

	times := sliceToInts(timesStr)
	distances := sliceToInts(distancesStr)
	return times, distances
}

func sliceToInts(slice []string) []int {
	var newSlice []int
	for _, str := range slice {
		num, err := strconv.Atoi(str)
		if err == nil {
			// panic("Failed to convert string to integer.")
			// Only process valid ints - ignore whitespace.
			newSlice = append(newSlice, int(num))
		}
	}
	return newSlice
}
