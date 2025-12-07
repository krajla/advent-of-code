package day7

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

func Part1() {
	lineIter := util.FileScanner("./day7/input.txt")

	res := uint64(0)
	prevLine := []byte{}
	for line := range lineIter {
		byteLine := []byte(line)

		if len(prevLine) == 0 {
			prevLine = byteLine
			continue
		}

		for i, c := range prevLine {
			switch c {
			case 'S':
				byteLine[i] = '|'
			case '|':
				if byteLine[i] == '^' {
					res++
					byteLine[i-1] = '|'
					byteLine[i+1] = '|'
				} else {
					byteLine[i] = '|'
				}
			}
		}

		prevLine = byteLine
	}

	fmt.Printf("Day6 Pt1 - Total: %d\n", res)
}

func Part2() {
	lineIter := util.FileScanner("./day7/input.txt")

	prevPaths := []int{}
	for line := range lineIter {
		paths := make([]int, len(line))

		if len(prevPaths) == 0 {
			paths[strings.Index(line, "S")] = 1
			prevPaths = paths
			continue
		}

		for i, v := range prevPaths {
			// Split paths.
			if line[i] == '^' {
				paths[i-1] += v
				paths[i+1] += v
			} else {
				paths[i] += v
			}

		}

		prevPaths = paths
	}

	res := uint64(0)
	for _, n := range prevPaths {
		res += uint64(n)
	}

	fmt.Printf("Day6 Pt1 - Total: %d\n", res)
}
