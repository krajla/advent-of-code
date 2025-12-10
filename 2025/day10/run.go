package day10

import (
	"adventofcode/util"
	"fmt"
	"math"
	"strings"
	"time"
)

type problem1 struct {
	goal    int
	options []int
}

func strToGoal(s string) int {
	s = s[1 : len(s)-1]
	goal := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			goal += 1 << i
		}
	}

	return goal
}

func strToOption(s string) int {
	s = s[1 : len(s)-1]
	option := 0
	for ns := range strings.SplitSeq(s, ",") {
		option += 1 << util.QuickAtoi(ns)
	}

	return option
}

func dfs1(goal int, current int, options []int) int {
	if goal == current {
		return 1
	}

	bestDepth := math.MaxInt
	for i := range options {
		depth := dfs1(goal, current^options[i], options[i+1:])
		if depth < bestDepth {
			bestDepth = depth
		}
	}

	if bestDepth == math.MaxInt {
		return bestDepth
	}

	return bestDepth + 1
}

func forAllCombinations1(goal int, options []int) int {
	res := dfs1(goal, 0, options) - 1
	return res
}

func Part1() time.Duration {
	lineIter := util.FileScanner("./day10/input.txt")

	problems := make([]problem1, 0, 200)
	for line := range lineIter {
		parts := strings.Split(line, " ")
		goal := strToGoal(parts[0])
		options := make([]int, 0, 10)
		for i := 1; i < len(parts)-1; i++ {
			options = append(options, strToOption(parts[i]))
		}
		problems = append(problems, problem1{
			goal:    goal,
			options: options,
		})
	}

	t := time.Now()
	res := 0
	for _, p := range problems {
		res += forAllCombinations1(p.goal, p.options)
	}

	fmt.Printf("Day10 Pt1 - Total: %d\n", res)
	return time.Since(t)
}
