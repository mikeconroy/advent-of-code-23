package day19

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day19/input")
	return part1(input), part2(input)
}

type Part struct {
	x int
	m int
	a int
	s int
}

type Category int

const (
	x = iota
	m
	a
	s
)

type Operation int

const (
	gt = iota
	lt = iota
)

type Rule struct {
	cat  Category
	op   Operation
	val  int
	wfId string
}

type Workflow struct {
	id    string
	rules []Rule
	def   string
}

/*
 *	Part - x,m,a,s
 *	Workflow - id, rules, default
 *  Rule - category (x,m,a,s), operator (<,>), value, workflow ID if rule is met.
 *  Start at workflow with id 'in'.
 *  Possible Optimization - Some workflows can be removed where the outcome is always A. E.g. lnx{m>1548:A,A} -> will always output A.
 *	This also then means workflows referencing those can be updated as well -> qs{s>3448:A,lnx} -> qs{s>3448:A,A} -> qs = A
 */
func part1(input []string) string {
	workflows, parts := readInput(input)
	totalAccepted := 0
	for _, part := range parts {

		wfId := "in"
		for wfId != "A" && wfId != "R" {
			wfId = processWorkflow(workflows[wfId], part)
		}

		if wfId == "A" {
			totalAccepted += part.x
			totalAccepted += part.m
			totalAccepted += part.a
			totalAccepted += part.s
		}
	}
	return fmt.Sprint(totalAccepted)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

func processWorkflow(wf Workflow, part Part) string {
	for _, rule := range wf.rules {
		var lh int
		rh := rule.val
		if rule.cat == x {
			lh = part.x
		} else if rule.cat == m {
			lh = part.m
		} else if rule.cat == a {
			lh = part.a
		} else if rule.cat == s {
			lh = part.s
		}

		if rule.op == gt {
			if lh > rh {
				return rule.wfId
			}
		} else if rule.op == lt {
			if lh < rh {
				return rule.wfId
			}
		}
	}
	return wf.def
}

func readInput(in []string) (workflows map[string]Workflow, parts []Part) {
	var partsStartLine int
	workflows = make(map[string]Workflow)
	for i, line := range in {
		if line == "" {
			partsStartLine = i + 1
			break
		}
		workflow := readWorkflow(line)
		workflows[workflow.id] = workflow
	}

	parts = readParts(in[partsStartLine:])
	return workflows, parts
}

func readWorkflow(line string) Workflow {
	split := strings.Split(line, "{")
	// split = px	a<2006:qkq,m>2090:A,rfg}
	id := split[0]
	split = strings.Split(split[1], ",")
	// split = a<2006:qkq	m>2090:A	rfg}
	var rules []Rule

	for _, ruleS := range split[:len(split)-1] {
		var cat Category
		var op Operation
		var val int
		var wfId string

		catS := ruleS[0]
		if catS == 'a' {
			cat = a
		} else if catS == 'm' {
			cat = m
		} else if catS == 'x' {
			cat = x
		} else if catS == 's' {
			cat = s
		}

		opS := ruleS[1]
		if opS == '>' {
			op = gt
		} else if opS == '<' {
			op = lt
		}

		ruleSplit := strings.Split(ruleS, ":")
		// ruleSplit = a<2006		qkq
		valS := ruleSplit[0][2:]
		val, _ = strconv.Atoi(valS)
		wfId = ruleSplit[1]

		rules = append(rules, Rule{cat: cat, op: op, val: val, wfId: wfId})
	}

	def := split[len(split)-1]
	def = def[:len(def)-1]

	return Workflow{id: id, rules: rules, def: def}
}

func readParts(partsIn []string) (parts []Part) {
	for _, line := range partsIn {
		split := strings.Split(line, ",")
		// Split = {x=787	m=2655		a=1222		s=2876}
		xS := strings.Split(split[0], "=")[1]
		x, _ := strconv.Atoi(xS)

		mS := strings.Split(split[1], "=")[1]
		m, _ := strconv.Atoi(mS)

		aS := strings.Split(split[2], "=")[1]
		a, _ := strconv.Atoi(aS)

		sS := strings.Split(split[3], "=")[1]
		sS = sS[:len(sS)-1]
		s, _ := strconv.Atoi(sS)

		parts = append(parts, Part{x: x, m: m, a: a, s: s})
	}
	return parts
}
