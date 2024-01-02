package main

import (
	"flag"
	"fmt"

	"github.com/mikeconroy/advent-of-code-23/day1"
	"github.com/mikeconroy/advent-of-code-23/day2"
	"github.com/mikeconroy/advent-of-code-23/day3"
	"github.com/mikeconroy/advent-of-code-23/day4"
	"github.com/mikeconroy/advent-of-code-23/day5"
	"github.com/mikeconroy/advent-of-code-23/day6"
	"github.com/mikeconroy/advent-of-code-23/day7"
	"github.com/mikeconroy/advent-of-code-23/day8"
	"github.com/mikeconroy/advent-of-code-23/day9"
	"github.com/mikeconroy/advent-of-code-23/day10"
	"github.com/mikeconroy/advent-of-code-23/day11"
	"github.com/mikeconroy/advent-of-code-23/day12"
	"github.com/mikeconroy/advent-of-code-23/day13"
	"github.com/mikeconroy/advent-of-code-23/day14"
	"github.com/mikeconroy/advent-of-code-23/day15"
	"github.com/mikeconroy/advent-of-code-23/day16"
	"github.com/mikeconroy/advent-of-code-23/day17"
	"github.com/mikeconroy/advent-of-code-23/day18"
)

func main() {
	dayToRun := flag.Int("day", 0, "Which Day to run? Defaults to 0 (all)")
	flag.Parse()

	days := []func() (string, string){
		day1.Run,
		day2.Run,
		day3.Run,
		day4.Run,
		day5.Run,
		day6.Run,
		day7.Run,
		day8.Run,
		day9.Run,
		day10.Run,
		day11.Run,
		day12.Run,
		day13.Run,
		day14.Run,
		day15.Run,
		day16.Run,
		day17.Run,
		day18.Run,
	}

	if *dayToRun == 0 {
		for day, run := range days {
			runDay(day+1, run)
		}
	} else {
		runDay(*dayToRun, days[*dayToRun-1])
	}
}

func runDay(dayNum int, run func() (string, string)) {
	p1, p2 := run()
	fmt.Printf("Day %d\n\tP1: %s\n\tP2: %s\n", dayNum, p1, p2)
}
