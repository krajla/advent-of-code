package day11

import (
	"adventofcode/util"
	"fmt"
	"strings"
	"time"
)

func dfs1(connections map[string][]string, memo map[string]int, current string) int {
	if current == "out" {
		return 1
	}
	if v, ok := memo[current]; ok {
		return v
	}

	sum := 0
	for _, c := range connections[current] {
		sum += dfs1(connections, memo, c)
	}

	memo[current] = sum
	return sum
}

func Part1() time.Duration {
	lineIter := util.FileScanner("./day11/input.txt")

	connections := make(map[string][]string, 600)
	for line := range lineIter {
		parts := strings.Split(line, " ")
		from := parts[0]
		from = from[:len(from)-1]
		connections[from] = parts[1:]
	}

	t := time.Now()

	res := dfs1(connections, make(map[string]int, len(connections)), "you")
	fmt.Printf("Day11 Pt1 - Result: %d\n", res)

	return time.Since(t)
}

type path struct {
	out    int
	fft    int
	dac    int
	fftdac int
}

func dfs2(connections map[string][]string, memo map[string]path, current string) path {
	if current == "out" {
		return path{1, 0, 0, 0}
	}
	if v, ok := memo[current]; ok {
		return v
	}

	pathSum := path{}
	for _, c := range connections[current] {
		path := dfs2(connections, memo, c)
		pathSum.out += path.out
		pathSum.dac += path.dac
		pathSum.fft += path.fft
		pathSum.fftdac += path.fftdac
	}

	switch current {
	case "fft":
		pathSum.fft = pathSum.out
		pathSum.fftdac = pathSum.dac
	case "dac":
		pathSum.dac = pathSum.out
		pathSum.fftdac = pathSum.fft
	}

	memo[current] = pathSum
	return pathSum
}

func Part2() time.Duration {
	lineIter := util.FileScanner("./day11/input.txt")

	connections := make(map[string][]string, 600)
	for line := range lineIter {
		parts := strings.Split(line, " ")
		from := parts[0]
		from = from[:len(from)-1]
		connections[from] = parts[1:]
	}

	t := time.Now()

	res := dfs2(connections, make(map[string]path, len(connections)), "svr")
	fmt.Printf("Day11 Pt2 - Result: %d\n", res.fftdac)

	return time.Since(t)
}
