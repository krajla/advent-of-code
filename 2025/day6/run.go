package day6

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

type compute struct {
	mult uint64
	sum  uint64
}

func Part1() {
	lineIter := util.FileScanner("./day6/input.txt")

	res := uint64(0)
	computes := make([]compute, 3774)
	for i := range computes {
		computes[i].mult = 1
	}
	for line := range lineIter {
		i := 0
		for _, col := range strings.Split(line, " ") {
			switch col {
			case " ", "":
				break
			case "+":
				res += computes[i].sum
				i++
			case "*":
				res += computes[i].mult
				i++
			default:
				v := uint64(util.QuickAtoi(col))
				computes[i].mult *= v
				computes[i].sum += v
				i++
			}
		}
	}

	fmt.Printf("Day6 Pt1 - Total: %d\n", res)
}

type problem struct {
	nums       []int
	operator   string
	startIndex int
	endIndex   int
}

func Part2() {
	lineIter := util.FileScanner("./day6/input.txt")

	problems := make([]problem, 0, 3774)
	for line := range lineIter {
		if line[0] != '*' && line[0] != '+' {
			continue
		}

		// Easier to determine where problems start and end by looking at operators
		// otherwise left and right padding confuses us.
		currProblem := problem{}
		for i, c := range line {
			if c == '*' || c == '+' {
				currProblem.endIndex = i - 2
				problems = append(problems, currProblem)
				currProblem = problem{
					nums:       make([]int, 10),
					operator:   fmt.Sprintf("%c", c),
					startIndex: i,
				}
			}
		}

		currProblem.endIndex = len(line) - 1
		problems = append(problems, currProblem)
	}

	for line := range lineIter {
		if line[0] == '*' || line[0] == '+' {
			continue
		}

		for _, p := range problems {
			for i := 0; i <= p.endIndex-p.startIndex; i++ {
				if line[p.startIndex+i] == ' ' {
					continue
				}

				v := util.QuickAtoi(fmt.Sprintf("%c", line[i+p.startIndex]))
				p.nums[i] = p.nums[i]*10 + v
			}
		}
	}

	res := uint64(0)
	for _, p := range problems {
		res += computeProblem(p)
	}

	fmt.Printf("Day6 Pt2 - Total: %d\n", res)
}

func computeProblem(p problem) uint64 {
	total := uint64(0)
	if p.operator == "+" {
		for i := 0; i < p.endIndex-p.startIndex+1; i++ {
			total += uint64(p.nums[i])
		}
	}

	if p.operator == "*" {
		total = 1

		for i := 0; i < p.endIndex-p.startIndex+1; i++ {
			total *= uint64(p.nums[i])
		}
	}

	return total
}
