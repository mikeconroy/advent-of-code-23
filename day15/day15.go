package day15

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day15/input")
	return part1(input), part2(input)
}

/*
 * HASH - Turn chars into a single number between 0-255.
 * Start with current value as 0.
 * Determine ASCII Code for the current char in the string.
 * Increase current value by ASCII code above.
 * Current Value = Current Value * 17
 * Current Value = Current Value % 256
 *
 */
func part1(input []string) string {
	strs := strings.Split(input[0], ",")
	result := 0
	for _, str := range strs {
		result += hash(str)
	}
	return fmt.Sprint(result)
}

/*
 * rn=1
 * label = rn
 * hash(label) = 0 - This is the box the step belongs in.
 * operation is = or -
 * 	If dash - remove the lens with given label from the box if present. Maintain order of remaining lenses.
 *  If equals - Followed by a value which is the focal length.
 * 				If the label already exists in the box then replace it.
 *				If the label doesn't exist in the box then place the new lens behind any other lenses in the box.
 * Result is (Box Number + 1) * (Slot Number of the lens in the box) * Focal Length of the lens
 */
func part2(input []string) string {
	boxes := make([][]Lens, 256)

	for _, str := range strings.Split(input[0], ",") {
		if strings.Contains(str, "=") {
			parts := strings.Split(str, "=")
			label := parts[0]
			focalLength, _ := strconv.Atoi(parts[1])
			lens := Lens{
				label:       label,
				hash:        hash(label),
				focalLength: focalLength,
			}
			pos := findInBox(boxes[lens.hash], lens)
			if pos == -1 {
				boxes[lens.hash] = append(boxes[lens.hash], lens)
			} else {
				boxes[lens.hash][pos] = lens
			}

		} else {
			parts := strings.Split(str, "-")
			label := parts[0]
			hash := hash(label)
			removeFromBox(boxes[hash], label)
		}
	}

	result := 0
	for boxIdx, box := range boxes {
		pos := 0
		for _, lens := range box {
			if lens.focalLength != 0 {
				pos++
				power := (boxIdx + 1) * (pos) * (lens.focalLength)
				result += power
			}
		}
	}
	return fmt.Sprint(result)
}

func findInBox(box []Lens, lens Lens) int {
	for i, boxLens := range box {
		if lens.label == boxLens.label {
			return i
		}
	}
	return -1
}

func removeFromBox(box []Lens, label string) {
	for i, lens := range box {
		if lens.label == label {
			// removeFromSlice(box, i)
			box[i] = Lens{}
		}
	}
}

func removeFromSlice(slice []Lens, i int) []Lens {
	return append(slice[:i], slice[i+1:]...)
}

type Lens struct {
	label       string
	hash        int
	focalLength int
}

var hashCache = make(map[string]int)

func hash(s string) int {
	if cache, ok := hashCache[s]; ok {
		return cache
	}
	currentVal := 0
	for _, char := range s {
		currentVal += int(char)
		currentVal *= 17
		currentVal %= 256
	}
	hashCache[s] = currentVal
	return currentVal
}
