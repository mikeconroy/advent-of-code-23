package day8

import (
	"fmt"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day8/input")
	return part1(input), part2(input)
}

func part1(input []string) string {
	directions, tree := parseInput(input)
	currNode := tree["AAA"]

	steps := 0
	endFound := false
	i := 0
	for !endFound {
		if directions[i] == 'L' {
			currNode = tree[currNode.l]
		} else if directions[i] == 'R' {
			currNode = tree[currNode.r]
		}
		if currNode.id == "ZZZ" {
			endFound = true
		}
		steps++
		i = (i + 1) % len(directions)
	}

	return fmt.Sprint(steps)
}

func part2(input []string) string {
	directions, tree := parseInput(input)
	// Order of currNodes will be different every time due to looping over the tree Map.
	var currNodes []string
	for _, node := range tree {
		if node.id[2] == 'A' {
			currNodes = append(currNodes, node.id)
		}
	}

	/*
	 * First attempt worked by going through the paths step by step.
	 * This doesn't scale to larger path sizes & nodes.
	 *
	 * We need to identify the loops each path takes and then find the LCM between the loops (steps it takes to get to Z).
	 * A loop is recognised when a path reaches back to the start and the current direction is the same.
	 */

	var loopLengths []int

	for _, nodeId := range currNodes {
		node := tree[nodeId]

		// Check for loops
		steps := 0
		endFound := 0
		i := 0
		firstEndFound := -1
		secondEndFound := -1
		for endFound < 2 {
			if directions[i] == 'L' {
				node = tree[node.l]
			} else if directions[i] == 'R' {
				node = tree[node.r]
			}
			steps++
			i = (i + 1) % len(directions)

			if node.id[2] == 'Z' {
				endFound++
				if firstEndFound == -1 {
					firstEndFound = steps
				} else {
					secondEndFound = steps
				}
			}
		}
		loopLength := secondEndFound - firstEndFound
		loopLengths = append(loopLengths, loopLength)
	}

	lcmLoops := LCM(loopLengths[0], loopLengths[1], loopLengths[2:]...)
	return fmt.Sprint(lcmLoops)
}

// greatest common divisor (GCD) via Euclidean algorithm
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

type Node struct {
	id string
	l  string
	r  string
}

type Tree map[string]Node

func parseInput(in []string) ([]rune, Tree) {
	var directions []rune
	tree := make(Tree)

	for _, char := range in[0] {
		directions = append(directions, char)
	}

	for i := 2; i < len(in); i++ {
		line := strings.Split(in[i], " ")
		id := line[0]
		l := line[2][1:4]
		r := line[3][:3]
		node := Node{
			id: id,
			l:  l,
			r:  r,
		}
		tree[id] = node
	}

	return directions, tree
}

/*
 * Part 2 first attempt.
 * Algorithm is too slow for larger path sizes & several nodes.
 *	steps := 0
 *	i := 0
 *	endFound := false
 * 	for !endFound {
 *		if steps%1_000_000 == 0 {
 *			fmt.Println(steps)
 *		}
 *		var newNodeIds []string
 *		endFound = true
 *		for _, currNodeId := range currNodes {
 *			currNode := tree[currNodeId]
 *			if currNode.id[2] != 'Z' {
 *				endFound = false
 *			}
 *			if directions[i] == 'L' {
 *				newNodeIds = append(newNodeIds, tree[currNode.l].id)
 *			} else if directions[i] == 'R' {
 *				newNodeIds = append(newNodeIds, tree[currNode.r].id)
 *			}
 *		}
 *
 *		currNodes = newNodeIds
 *		steps++
 *		i = (i + 1) % len(directions)
 *	}
 */
