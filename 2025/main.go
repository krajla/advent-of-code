package main

import (
	day8 "adventofcode/day8"
	day8_dsu "adventofcode/day8_dsu"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	day8.Part1()
	elapsed := time.Since(t)
	fmt.Printf("Part1 arrays time: %v\n", elapsed)

	t = time.Now()
	day8_dsu.Part1()
	elapsed = time.Since(t)
	fmt.Printf("Part1 DSU time: %v\n", elapsed)

	fmt.Println()

	t = time.Now()
	day8.Part2()
	elapsed = time.Since(t)
	fmt.Printf("Part2 arrays time: %v\n", elapsed)

	t = time.Now()
	day8_dsu.Part2()
	elapsed = time.Since(t)
	fmt.Printf("Part2 DSU time: %v\n", elapsed)
}
