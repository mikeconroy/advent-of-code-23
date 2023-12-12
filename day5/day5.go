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

func part2(input []string) string {
	return fmt.Sprint(0)
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
	destinationStart int
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
				destinationStart: dest,
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
