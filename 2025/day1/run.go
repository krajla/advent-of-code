package day1

import (
	"adventofcode/util"
	"fmt"
)

func Part1() {
	lineIter := util.FileScanner("./day1/input.txt")

	dial := 50
	hits := 0
	for line := range lineIter {
		direction := line[0]
		distance := util.QuickAtoi(line[1:])
		switch direction {
		case 'L':
			dial -= distance
		case 'R':
			dial += distance
		}

		dial = mod(dial, 100)
		if dial == 0 {
			hits++
		}
	}

	fmt.Printf("Day1 Pt1 - Hits: %d\n", hits)
}

func Part2() {
	lineIter := util.FileScanner("./day1/input.txt")

	max := 100
	dial := 50
	hits := 0
	for line := range lineIter {
		direction := line[0]
		distance := util.QuickAtoi(line[1:])
		switch direction {
		case 'L':
			// Rotating left from 0 double counts.
			if dial == 0 {
				hits--
			}
			dial -= distance
		case 'R':
			dial += distance
		}

		hits += spins(dial, max)
		dial = mod(dial, max)
	}

	fmt.Printf("Day1 Pt2 - Hits: %d\n", hits)
}

func mod(dial, max int) int {
	dial %= max
	if dial < 0 {
		dial += max
	}
	return dial
}

func spins(dial, max int) int {
	s := dial / max
	if dial <= 0 {
		s *= -1
		s++
	}
	return s
}
