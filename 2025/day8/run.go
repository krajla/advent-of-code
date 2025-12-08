package day8

import (
	"adventofcode/util"
	"fmt"
	"math"
	"slices"
	"strings"
)

type node struct {
	x int
	y int
	z int
}

func (a node) cmp(b node) int {
	if a.x == b.x {
		if a.y == b.y {
			return b.z - a.z
		}

		return b.y - a.y
	}

	return b.x - a.x
}

func (a node) nodeDistance(b node) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	dz := float64(a.z - b.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type edge struct {
	n1       node
	n2       node
	distance float64
}

type group []node

func Part1() {
	lineIter := util.FileScanner("./day8/input.txt")

	groups := make([]group, 0, 1000)
	for line := range lineIter {
		parts := strings.Split(line, ",")
		node := node{
			x: util.QuickAtoi(parts[0]),
			y: util.QuickAtoi(parts[1]),
			z: util.QuickAtoi(parts[2]),
		}

		groups = append(groups, group{node})
	}

	edges := make([]edge, 0, 1000000)
	for i, g1 := range groups {
		for _, g2 := range groups[i+1:] {
			edges = append(edges, edge{
				n1:       g1[0],
				n2:       g2[0],
				distance: g1[0].nodeDistance(g2[0]),
			})
		}
	}
	slices.SortFunc(edges, func(a, b edge) int {
		return int(a.distance - b.distance)
	})

	for _, e := range edges[:1000] {
		group1, group2 := 0, 0
		for i, g := range groups {
			for _, n := range g {
				if e.n1.cmp(n) == 0 {
					group1 = i
				}

				if e.n2.cmp(n) == 0 {
					group2 = i
				}
			}
		}

		if group1 != group2 {
			groups[group1] = append(groups[group1], groups[group2]...)
			groups[group2] = nil
		}
	}

	lens := make([]int, 0, len(groups))
	for _, g := range groups {
		if g == nil {
			continue
		}
		lens = append(lens, len(g))
	}
	slices.Sort(lens)

	res := 1
	for i := range 3 {
		res *= lens[len(lens)-i-1]
	}

	fmt.Printf("Day8 Pt1 - Total: %d\n", res)
}

func Part2() {
	lineIter := util.FileScanner("./day8/input.txt")

	groups := make([]group, 0, 1000)
	for line := range lineIter {
		parts := strings.Split(line, ",")
		node := node{
			x: util.QuickAtoi(parts[0]),
			y: util.QuickAtoi(parts[1]),
			z: util.QuickAtoi(parts[2]),
		}

		groups = append(groups, group{node})
	}

	edges := make([]edge, 0, 1000000)
	for i, g1 := range groups {
		for _, g2 := range groups[i+1:] {
			edges = append(edges, edge{
				n1:       g1[0],
				n2:       g2[0],
				distance: g1[0].nodeDistance(g2[0]),
			})
		}
	}
	slices.SortFunc(edges, func(a, b edge) int {
		return int(a.distance - b.distance)
	})

	res := uint64(0)
	connectionsLeftUntilConnected := len(groups) - 1
	for i := 0; i < len(edges) && connectionsLeftUntilConnected > 0; i++ {
		e := edges[i]

		group1, group2 := 0, 0
		for i, g := range groups {
			for _, n := range g {
				if e.n1.cmp(n) == 0 {
					group1 = i
				}

				if e.n2.cmp(n) == 0 {
					group2 = i
				}
			}
		}

		if group1 != group2 {
			groups[group1] = append(groups[group1], groups[group2]...)
			groups[group2] = nil

			connectionsLeftUntilConnected--
		}

		if connectionsLeftUntilConnected == 0 {
			res = uint64(e.n1.x) * uint64(e.n2.x)
		}
	}

	fmt.Printf("Day8 Pt2 - Total: %d\n", res)
}
