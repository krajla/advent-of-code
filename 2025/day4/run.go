package day4

import (
	"adventofcode/util"
	"fmt"
	"iter"
)

const cuttof = 4

func Part1() {
	lineIter := util.FileScanner("./day4/input.txt")

	m, mutation := initMatrix(lineIter)
	for i := range m {
		for j := range m[i] {
			determineNeighbors(i, j, m)
		}
	}

	res := reduceMatrix(m, mutation)

	fmt.Printf("Day4 Pt1 - Sum: %d\n", res)
}

func Part2() {
	lineIter := util.FileScanner("./day4/input.txt")

	m, mutation := initMatrix(lineIter)
	for i := range m {
		for j := range m[i] {
			determineNeighbors(i, j, m)
		}
	}

	res := uint64(0)
	for r := reduceMatrix(m, mutation); r != 0; r = reduceMatrix(m, mutation) {
		res += uint64(r)
		applyMutation(m, mutation)
	}

	fmt.Printf("Day4 Pt2 - Sum: %d\n", res)
}

func initMatrix(input iter.Seq[string]) ([][]int, [][]int) {
	m := make([][]int, 0, 140)
	mutation := make([][]int, 0, 140)
	for line := range input {
		ml := make([]int, len(line))
		mml := make([]int, len(line))
		m = append(m, ml)
		mutation = append(mutation, mml)
		for i, c := range line {
			if c == '.' {
				ml[i] = -1
			} else {
				ml[i] = 1
			}
		}
	}

	return m, mutation
}

func reduceMatrix(m, mutation [][]int) int {
	total := 0
	for i := range m {
		for j := range m[i] {
			if m[i][j] >= cuttof || m[i][j] < 0 {
				continue
			}

			mutation[i][j] = -200
			total++
			lowerNeighbors(i, j, m, mutation)
		}
	}

	return total
}

func lowerNeighbors(x, y int, m, mutation [][]int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i >= 0 && i < len(m) && j >= 0 && j < len(m[0]) {
				mutation[i][j]--
			}
		}
	}
}

func determineNeighbors(x, y int, m [][]int) {
	if m[x][y] < 0 {
		return
	}

	total := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i >= 0 && i < len(m) && j >= 0 && j < len(m[0]) {
				if (i != x || j != y) && m[i][j] > 0 {
					total++
				}
			}
		}
	}

	m[x][y] = total
}

func applyMutation(m, mutation [][]int) {
	for i := range m {
		for j := range m[0] {
			m[i][j] += mutation[i][j]
			mutation[i][j] = 0
		}
	}
}

func print(m [][]int, raw bool) {
	for i := range m {
		for j := range m[i] {
			if raw {
				fmt.Printf("%d ", m[i][j])
				continue
			}
			c := '.'
			if m[i][j] > 0 {
				c = '@'
			}
			if m[i][j] <= -200 {
				c = 'X'
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}
