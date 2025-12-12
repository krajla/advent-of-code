package day12

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

func Part1() {
	lineIter := util.FileScanner("./day12/input.txt")

	res := 0
	shapeSized := make([]int, 0, 5)
	currentShapeSize := 0
	for line := range lineIter {
		if line == "" {
			shapeSized = append(shapeSized, currentShapeSize)
			currentShapeSize = 0
			continue
		}
		if line[0] == '.' || line[0] == '#' {
			for _, c := range line {
				if c == '#' {
					currentShapeSize++
				}
			}
			continue
		}
		if line[1] == ':' {
			continue
		}

		parts := strings.Split(line, " ")
		sizeS := strings.Split(parts[0], "x")
		size := util.QuickAtoi(sizeS[0]) * util.QuickAtoi(sizeS[1][:len(sizeS[1])-1])

		for i, p := range parts[1:] {
			size -= util.QuickAtoi(p) * shapeSized[i]
		}

		if size >= 0 {
			res++
		}
	}

	fmt.Printf("Day12 - Result: %d\n", res)
}
