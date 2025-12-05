package day5

import (
	"adventofcode/util"
	"fmt"
	"slices"
	"strings"
)

func Part1() {
	lineIter := util.FileScanner("./day5/input.txt")

	res := uint64(0)
	readingRanges := true
	ranges := make([]util.Range, 0, 100)
	for line := range lineIter {
		if line == "" {
			readingRanges = false
			continue
		}

		if readingRanges {
			parts := strings.Split(line, "-")
			r := util.Range{
				Start: util.QuickAtoi(parts[0]),
				End:   util.QuickAtoi(parts[1]),
			}

			ranges = append(ranges, r)
		} else {
			v := util.QuickAtoi(line)
			for _, r := range ranges {
				if r.Start <= v && r.End >= v {
					res++
					break
				}
			}
		}
	}

	fmt.Printf("Day5 Pt1 - Total: %d\n", res)
}

func Part2() {
	lineIter := util.FileScanner("./day5/input.txt")

	res := uint64(0)
	ranges := make([]util.Range, 0, 100)
	for line := range lineIter {
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		r := util.Range{
			Start: util.QuickAtoi(parts[0]),
			End:   util.QuickAtoi(parts[1]),
		}

		ranges, r = popOverlaps(ranges, r)

		index, _ := slices.BinarySearchFunc(ranges, r, util.CmpRange)
		ranges = slices.Insert(ranges, index, r)
	}

	for _, r := range ranges {
		res += uint64(r.End - r.Start + 1)
	}

	fmt.Printf("Day5 Pt2 - Total: %d\n", res)
}

func popOverlaps(ranges []util.Range, newRange util.Range) ([]util.Range, util.Range) {
	newRanges := make([]util.Range, 0, len(ranges))
	minMax := newRange
	for _, r := range ranges {
		if r.Overlaps(newRange) {
			minMax.Start = util.Min(minMax.Start, r.Start)
			minMax.End = util.Max(minMax.End, r.End)
		} else {
			newRanges = append(newRanges, r)
		}
	}

	return newRanges, minMax
}
