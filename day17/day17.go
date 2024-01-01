package day17

import (
	"container/heap"
	"fmt"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day17/input")
	return part1(input), part2(input)
}

/*
 * Input is a [][]int representing city blocks and heat loss associated to that block.
 * Top Left ([0][0]) is the starting block - heat loss is not incurred unless the block is left & re-entered (which there shouldn't be a reason to do).
 * Bottom Right is the destination block.
 * Must turn left or right after 3 blocks of going straight.
 * Can't reverse direction.
 */
func part1(input []string) string {
	city := loadBlocks(input)
	start := Point{x: 0, y: 0}
	end := Point{x: len(city[len(city)-1]) - 1, y: len(city) - 1}
	minDir := 0
	maxDir := 3
	return fmt.Sprint(minHeatLoss(city, start, end, minDir, maxDir))
}

type Point struct {
	x int
	y int
}

// Based on:
//
//	https://www.youtube.com/watch?v=2pDSooPLLkI
func minHeatLoss(grid [][]int, start Point, end Point, minDir int, maxDir int) int {
	seen := make(map[string]bool)
	startBlock := &CityBlock{
		pos: start,
		hl:  0,
		dY:  0,
		dX:  0,
		n:   0,
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, startBlock)

	for pq.Len() > 0 {
		currBlock := heap.Pop(&pq).(*CityBlock)
		if currBlock.pos == end && currBlock.n >= minDir {
			return currBlock.hl
		}

		if seen[currBlock.key()] == true {
			continue
		}

		seen[currBlock.key()] = true

		// Continue moving in the same diretion if less than 3 and not stationary (start).
		if currBlock.n < maxDir && !(currBlock.dX != 0 && currBlock.dY != 0) {
			newY := currBlock.pos.y + currBlock.dY
			newX := currBlock.pos.x + currBlock.dX
			if newY >= 0 && newY < len(grid) && newX >= 0 && newX < len(grid[newY]) {
				newCB := &CityBlock{
					pos: Point{x: newX, y: newY},
					hl:  currBlock.hl + grid[newY][newX],
					dY:  currBlock.dY,
					dX:  currBlock.dX,
					n:   currBlock.n + 1,
				}
				heap.Push(&pq, newCB)
			}
		}

		// Add valid neighbours to PQ
		if currBlock.n >= minDir || currBlock.pos == start {
			for _, neighbour := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				// Make sure we don't handle the same direction again
				if neighbour[0] == currBlock.dY && neighbour[1] == currBlock.dX {
					continue
				}
				// Make sure we don't handle the reverse direction
				if neighbour[0] == -(currBlock.dY) && neighbour[1] == -(currBlock.dX) {
					continue
				}

				newY := currBlock.pos.y + neighbour[0]
				newX := currBlock.pos.x + neighbour[1]
				if newY >= 0 && newY < len(grid) && newX >= 0 && newX < len(grid[newY]) {
					neighbourBlock := &CityBlock{
						pos: Point{x: newX, y: newY},
						hl:  currBlock.hl + grid[newY][newX],
						dY:  neighbour[0],
						dX:  neighbour[1],
						n:   1,
					}
					heap.Push(&pq, neighbourBlock)
				}
			}
		}

	}
	return -1
}

func (cb *CityBlock) key() string {
	return fmt.Sprintf("%d,%d,%d,%d,%d", cb.pos.x, cb.pos.y, cb.dX, cb.dY, cb.n)
}

func part2(input []string) string {
	city := loadBlocks(input)
	start := Point{x: 0, y: 0}
	end := Point{x: len(city[len(city)-1]) - 1, y: len(city) - 1}
	minDir := 4
	maxDir := 10
	return fmt.Sprint(minHeatLoss(city, start, end, minDir, maxDir))
}

func loadBlocks(in []string) (blocks [][]int) {
	for _, row := range in {
		var newRow []int
		for _, val := range row {
			newRow = append(newRow, int(val-'0'))
		}
		blocks = append(blocks, newRow)
	}
	return blocks
}
