package main

import (
	"adventofcode/day11"
	"fmt"
)

func main() {
	t := day11.Part1()
	fmt.Printf("Part1 time: %v\n", t.Nanoseconds())

	t = day11.Part2()
	fmt.Printf("Part time: %v\n", t.Nanoseconds())
}
