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
	return fmt.Sprint(0)
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
