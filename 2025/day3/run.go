package day3

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

func Part1() {
	lineIter := util.FileScanner("./day3/input.txt")

	res := uint64(0)
	for line := range lineIter {
		cmax := 0
		max1 := 0
		max2 := 0

		for i := range line {
			n := util.QuickAtoi(line[i : i+1])

			if n > max1 {
				lastCMax := max1*10 + max2
				if lastCMax > cmax {
					cmax = lastCMax
				}
				lastCMax = max1*10 + n
				if lastCMax > cmax {
					cmax = lastCMax
				}

				max1 = n
				max2 = 0
				n = 0
			}

			if n > max2 {
				max2 = n
				n = 0
			}
		}
		if max1 != 0 && max2 != 0 {
			res += uint64(max1*10 + max2)
		} else {
			res += uint64(cmax)
		}
	}

	fmt.Printf("Day3 Pt1 - Sum: %d\n", res)
}

func Part2() {
	lineIter := util.FileScanner("./day3/input.txt")

	res := uint64(0)
	for line := range lineIter {
		// Reverse bank and cast.
		bank := make([]int, len(line))
		for i, v := range strings.Split(line, "") {
			bank[len(bank)-i-1] = util.QuickAtoi(v)
		}

		bestJolt := make([]int, 12)
		if len(bank) < 12 {
			panic("bank < 12")
		}

		// Consider last 12 values as best initialy.
		for i := range 12 {
			bestJolt[i] = bank[11-i]
		}

		// Going back to front(because we iterate reversed bank) keep highest digits found as most significant and trickle replaces digits down.
		// By iterating backwards we can guarantee at any point that we selected the best 12 digits from the end until our point.
		bank = bank[12:]
		for i := range bank {
			n := bank[i]

			for j := 0; j < len(bestJolt) && n >= bestJolt[j]; j++ {
				bestJolt[j], n = n, bestJolt[j]
			}
		}

		res += uint64(util.NumSliceToNum(bestJolt))
	}

	fmt.Printf("Day3 Pt2 - Sum: %d\n", res)
}
