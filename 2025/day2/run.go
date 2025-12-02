package day2

import (
	"adventofcode/util"
	"fmt"
	"math"
	"strings"
)

var multiplesTable = map[int][]int{
	2:  {1},
	3:  {1},
	4:  {1, 2},
	5:  {1},
	6:  {1, 2, 3},
	7:  {1},
	8:  {1, 2, 4},
	9:  {1, 3},
	10: {1, 2, 5},
	11: {1},
	12: {1, 2, 3, 4, 6},
}

func Part1() {
	lineIter := util.FileScanner("./day2/input.txt")

	res := uint64(0)
	for line := range lineIter {
		for valueRange := range strings.SplitSeq(line, ",") {
			parts := strings.Split(valueRange, "-")
			lower := util.QuickAtoi(parts[0])
			upper := util.QuickAtoi(parts[1])

			if upper < lower {
				panic("upper < lower")
			}

			lenLower := util.LenInt(lower)
			lenUpper := util.LenInt(upper)

			// Extract significant digits(those duplicated) for limits of our ranges
			// and make sure they start from a valid target, skipping over odd lenghts.

			sLower := 0
			if lenLower%2 == 0 {
				sLower = significantLowerBound(lower, lenLower/2)
			} else {
				sLower = int(math.Pow10(lenLower / 2))
			}

			sUpper := 0
			if lenUpper%2 == 0 {
				sUpper = significantUpperBound(upper, lenUpper/2)
			} else {
				sUpper = int(math.Pow10(lenUpper/2)) - 1
			}

			// Iterative sum good enough for Pt1.
			res += iterativeSum(sLower, sUpper)
		}
	}

	fmt.Printf("Day2 Pt1 - Sum: %d\n", res)
}

func Part2() {
	lineIter := util.FileScanner("./day2/input.txt")

	res := uint64(0)
	for line := range lineIter {
		for valueRange := range strings.SplitSeq(line, ",") {
			parts := strings.Split(valueRange, "-")
			lower := util.QuickAtoi(parts[0])
			upper := util.QuickAtoi(parts[1])
			if upper < lower {
				panic("upper < that lower")
			}

			// Split problem into sub-ranges of the same legth and sum them up.
			lowerLen := util.LenInt(lower)
			upperLen := util.LenInt(upper)
			for i := lowerLen; i <= upperLen; i++ {
				// If lower bound is smaller in length that i
				// assume smallest number of length i.
				a := lower
				if lowerLen < i {
					a = int(math.Pow10(i - 1))
				}

				// If upper bound is larger in length than i
				// assume largest number of length i.
				b := upper
				if i < upperLen {
					b = int(math.Pow10(i)) - 1
				}
				res += sumForBoundLength(a, b, i)
			}
		}
	}

	fmt.Printf("Day2 Pt2 - Sum: %d\n", res)
}

func iterativeSum(a, b int) uint64 {
	res := uint64(0)
	for ; a <= b; a++ {
		res += mplicate(a, util.LenInt(a), 2)
	}

	return res
}

func sumForBoundLength(a, b, n int) uint64 {
	currentSet, ok := multiplesTable[n]
	if !ok {
		return 0
	}

	totals := make(map[int]uint64, len(currentSet))
	// Split problem into subdivisions of size m and sum the up.
	for _, seq := range currentSet {

		// Instead of summing mplicated digit sequences, we calculate the sum of the series
		// then mplicate with what would have been the digit sequence length.
		totals[seq] = mplicate(sumForDigitSequence(a, b, seq), seq, n/seq)
	}

	// Since we don't count every element and only do operations on sequences
	// we must delete values counted multiples times, we do this by subtracting
	// totals from subsequences that divide our subsequence.
	total := uint64(0)
	for i := range totals {
		dedupTotal := totals[i]
		for _, dupI := range multiplesTable[i] {
			dedupTotal -= totals[dupI]
		}

		//fmt.Printf("%d %d length: %d seq: %d = %d\n", a, b, n, i, dedupTotal)
		total += dedupTotal
	}
	return total
}

func sumForDigitSequence(a, b, seq int) int {
	sLower := float64(significantLowerBound(a, seq))
	sUpper := float64(significantUpperBound(b, seq))
	if sUpper < sLower {
		return 0
	}

	return int(((sUpper - sLower + 1) / 2) * (sLower + sUpper))
}

func mplicate(n, seq, multiples int) uint64 {
	complete := uint64(0)
	pow := uint64(math.Pow10(seq))
	for range multiples {
		complete = complete*pow + uint64(n)
	}
	return complete
}

func significantLowerBound(v, significantLen int) int {
	len := util.LenInt(v)
	significantValue := v / int(math.Pow10(len-significantLen))

	if v > int(mplicate(significantValue, significantLen, len/significantLen)) {
		return significantValue + 1
	}

	return significantValue
}

func significantUpperBound(v, significantLen int) int {
	len := util.LenInt(v)
	significantValue := v / int(math.Pow10(len-significantLen))

	if v < int(mplicate(significantValue, significantLen, len/significantLen)) {
		return significantValue - 1
	}

	return significantValue
}
