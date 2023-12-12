package day5

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day5/input")
	return part1(input), part2(input)
}

/*
 *	Almanac (input) lists:
 * 		Seeds to be planted
 *  	Seed To Soil			- The type of soil required by each seed.
 *  	Soil To Fertilizer
 *		Fertilizer To Water
 *  	Water To Light
 *  	Light to Temperature
 *  	Temperature To Humidity
 *  	Humidity To Location
 * 	Each line within a map contains three numbers:
 * 		The destination range start
 *		The source range start
 *		The range length
 */
func part1(input []string) string {
	alm := loadAlmanac(input)
	smallestLoc := math.MaxInt
	for _, seed := range alm.seeds {
		soil := findSourceInRanges(seed, alm.seedToSoil)
		fertilizer := findSourceInRanges(soil, alm.soilToFertilizer)
		water := findSourceInRanges(fertilizer, alm.fertilizerToWater)
		light := findSourceInRanges(water, alm.waterToLight)
		temp := findSourceInRanges(light, alm.lightToTemperature)
		humidity := findSourceInRanges(temp, alm.temperatureToHumidity)
		location := findSourceInRanges(humidity, alm.humidityToLocation)
		if location < smallestLoc {
			smallestLoc = location
		}
	}
	return fmt.Sprint(smallestLoc)
}

// Seeds are input as pairs:
//
//	First value is the start of a seed range
//	Second value is the length of the range.
//
// Now calculate the lowest location possible based on all these seeds.
func part2(input []string) string {
	// Not feasible to run Part 1 logic on every single seed possible (9 digit ranges)

	// Can we work backwards - from Locations to Seeds?
	// Identify the lowest possible location and check if feasible? Still large ranges of Locations.

	// Sets / Algorithm that recognises Range Overlaps.
	// May be partial overlaps and multiple overlaps. Won't tell us the destination value or whether it's higher or lower.

	// Reduce inputs/ranges based on once a higher value is found - remove it as an option.

	// We only care about numbers on the edge of ranges?
	// Numbers not in a range are themselves.
	// Numbers in a range would be the same (n+1) for most others...

	// Work through each type and treat ranges as a single unit.
	// At the end we should end up with the possible ranges of Locations.
	// Select the range with the lowest start number.
	alm := loadAlmanac(input)

	// Holds the ranges possible at the end of each step.
	var currentRanges []RangeMap
	// Load the seed ranges
	for i := 0; i < len(alm.seeds); i++ {
		if i%2 == 0 {
			seeds := RangeMap{
				sourceStart: alm.seeds[i],
				sourceEnd:   alm.seeds[i] + alm.seeds[i+1] - 1,
				rangeLength: alm.seeds[i+1],
			}
			currentRanges = append(currentRanges, seeds)
		}
	}
	currentRanges = findPossibleRanges(currentRanges, alm.seedToSoil)
	currentRanges = findPossibleRanges(currentRanges, alm.soilToFertilizer)
	currentRanges = findPossibleRanges(currentRanges, alm.fertilizerToWater)
	currentRanges = findPossibleRanges(currentRanges, alm.waterToLight)
	currentRanges = findPossibleRanges(currentRanges, alm.lightToTemperature)
	currentRanges = findPossibleRanges(currentRanges, alm.temperatureToHumidity)
	currentRanges = findPossibleRanges(currentRanges, alm.humidityToLocation)

	lowestStart := math.MaxInt
	for _, possibleRange := range currentRanges {
		if possibleRange.sourceStart < lowestStart {
			lowestStart = possibleRange.sourceStart
		}
	}

	return fmt.Sprint(lowestStart)
}

// Need to store the ranges as the size of them is big.
type Almanac struct {
	seeds                 []int
	seedToSoil            []RangeMap
	soilToFertilizer      []RangeMap
	fertilizerToWater     []RangeMap
	waterToLight          []RangeMap
	lightToTemperature    []RangeMap
	temperatureToHumidity []RangeMap
	humidityToLocation    []RangeMap
}

type RangeMap struct {
	sourceStart      int
	sourceEnd        int
	destinationStart int
	destinationEnd   int
	rangeLength      int
}

const (
	SEED_TO_SOIL            = 0
	SOIL_TO_FERTILIZER      = 1
	FERTILIZER_TO_WATER     = 2
	WATER_TO_LIGHT          = 3
	LIGHT_TO_TEMPERATURE    = 4
	TEMPERATURE_TO_HUMIDITY = 5
	HUMIDITY_TO_LOCATION    = 6
)

func loadAlmanac(input []string) Almanac {
	alm := Almanac{}
	seeds := sliceToInts(strings.Split(input[0], " ")[1:])

	input = input[3:]

	currentMapType := SEED_TO_SOIL

	for _, line := range input {
		if line != "" && !strings.Contains(line, "map") {
			strs := strings.Split(line, " ")
			source, _ := strconv.Atoi(strs[1])
			dest, _ := strconv.Atoi(strs[0])
			length, _ := strconv.Atoi(strs[2])
			rangeMapping := RangeMap{
				sourceStart:      source,
				sourceEnd:        source + length - 1,
				destinationStart: dest,
				destinationEnd:   dest + length - 1,
				rangeLength:      length,
			}
			switch currentMapType {
			case SEED_TO_SOIL:
				alm.seedToSoil = append(alm.seedToSoil, rangeMapping)
			case SOIL_TO_FERTILIZER:
				alm.soilToFertilizer = append(alm.soilToFertilizer, rangeMapping)
			case FERTILIZER_TO_WATER:
				alm.fertilizerToWater = append(alm.fertilizerToWater, rangeMapping)
			case WATER_TO_LIGHT:
				alm.waterToLight = append(alm.waterToLight, rangeMapping)
			case LIGHT_TO_TEMPERATURE:
				alm.lightToTemperature = append(alm.lightToTemperature, rangeMapping)
			case TEMPERATURE_TO_HUMIDITY:
				alm.temperatureToHumidity = append(alm.temperatureToHumidity, rangeMapping)
			case HUMIDITY_TO_LOCATION:
				alm.humidityToLocation = append(alm.humidityToLocation, rangeMapping)
			}
		} else if strings.Contains(line, "map:") {
			currentMapType++
		}
	}

	alm.seeds = seeds
	return alm
}

func sliceToInts(slice []string) []int {
	var newSlice []int
	for _, str := range slice {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic("Failed to convert string to integer.")
		}
		newSlice = append(newSlice, int(num))
	}
	return newSlice
}

// May be possible to apply caching here to reduce searches for the same IDs. Not necesarry for Part 1.
func findSourceInRanges(source int, ranges []RangeMap) int {
	for _, rngeMap := range ranges {
		if source >= rngeMap.sourceStart && source <= rngeMap.sourceStart+rngeMap.rangeLength {
			diff := source - rngeMap.sourceStart
			return rngeMap.destinationStart + diff
		}
	}
	return source
}

// VERY MESSY CODE!!!!
// Takes a current set of ranges & returns new ranges based on ranges in destRanges.
// E.g. currRanges: [1-3, 7-9, 20-21] destRanges; [2-3, 19-22]
// Becomes: [1-1, 2-3, 7-9, 20-21]
//
// Adds currRanges to a list of toProcess
// Loops whilst there are items in toProcess list
// for each range in toProcess
//
//	check if the range intersects any of the destRanges.
//	As soon as an intersection is found stop looping and:
//		If the currRange exists within a range entirely - use the destiation values to map to the new range
//		If the currRange start or end exists within a range -
//			Add the overlap of ranges to possibleRanges (mapped to destination values) AND add the remainder of the range into the toProcess list to check for further overlaps.
//	If no intersection is found - keep the current range in possible ranges.
func findPossibleRanges(currRanges, destRanges []RangeMap) []RangeMap {
	var possibleRanges []RangeMap
	toProcess := currRanges
	for len(toProcess) > 0 {
		// fmt.Println("To Process:", toProcess)
		var newRangesToProcess []RangeMap
		for _, currRange := range toProcess {
			rangeFound := false
			for _, destRange := range destRanges {
				if currRange.sourceStart >= destRange.sourceStart && currRange.sourceEnd <= destRange.sourceEnd {
					// Range fits completeley in the new range.
					possibleRanges = append(possibleRanges, getNewRange(currRange, destRange))
					// fmt.Println("Range fits, adding to possibleRanges:", currRange, destRange, possibleRanges)
					rangeFound = true
					break
				} else if currRange.sourceStart < destRange.sourceStart && currRange.sourceEnd >= destRange.sourceStart {
					// The end of the range is within the range.
					newPossibleRange := RangeMap{
						sourceStart: destRange.sourceStart,
						sourceEnd:   currRange.sourceEnd,
						rangeLength: currRange.sourceEnd - destRange.sourceStart,
					}
					possibleRanges = append(possibleRanges, getNewRange(newPossibleRange, destRange))
					newRangeToProcess := RangeMap{
						sourceStart: currRange.sourceStart,
						sourceEnd:   destRange.sourceStart - 1,
						rangeLength: currRange.sourceStart - 1 - currRange.sourceStart,
					}
					newRangesToProcess = append(newRangesToProcess, newRangeToProcess)
					rangeFound = true
					// fmt.Println("End of Range is within Range, adding to possibleRanges:", currRange, destRange, newPossibleRange, newRangeToProcess)
					break
				} else if currRange.sourceStart < destRange.sourceEnd && currRange.sourceEnd > destRange.sourceEnd {
					// The start of the range is within the range.
					// TODO: These conditions dont account for our currRange supersetting the dest range.
					// The end of the range is within the range.
					newPossibleRange := RangeMap{
						sourceStart: currRange.sourceStart,
						sourceEnd:   destRange.sourceEnd,
						rangeLength: destRange.sourceEnd - currRange.sourceStart,
					}
					possibleRanges = append(possibleRanges, getNewRange(newPossibleRange, destRange))
					newRangeToProcess := RangeMap{
						sourceStart: destRange.sourceEnd + 1,
						sourceEnd:   currRange.sourceEnd,
						rangeLength: currRange.sourceEnd - destRange.sourceEnd,
					}
					newRangesToProcess = append(newRangesToProcess, newRangeToProcess)
					rangeFound = true
					// fmt.Println("Start of Range is within Range, adding to possibleRanges:", currRange, destRange, newPossibleRange, newRangeToProcess)
					break
				} else {
					// Range is entirely out of the range.
					rangeFound = false
				}
			}
			if !rangeFound {
				// fmt.Println("Range is not found:", currRange, destRanges)
				possibleRanges = append(possibleRanges, currRange)
			}
		}
		// fmt.Println("New Ranges to Process:", newRangesToProcess)
		toProcess = newRangesToProcess
	}

	return possibleRanges
}

func getNewRange(source, dest RangeMap) RangeMap {
	diff := dest.sourceStart - dest.destinationStart
	// fmt.Println("NEW RANGE:", source, dest, diff)

	return RangeMap{
		sourceStart: source.sourceStart - diff,
		sourceEnd:   source.sourceEnd - diff,
		rangeLength: source.rangeLength,
	}
}
