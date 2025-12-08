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

func Part1() {
	lineIter := util.FileScanner("./day8/input.txt")

	groups := util.NewDSU[node]()
	edges := make([]edge, 0, 1000000)
	{
		nodes := make([]node, 0, 1000)
		for line := range lineIter {
			parts := strings.Split(line, ",")
			node := node{
				x: util.QuickAtoi(parts[0]),
				y: util.QuickAtoi(parts[1]),
				z: util.QuickAtoi(parts[2]),
			}

			groups.Find(node)
			nodes = append(nodes, node)
		}

		for i, n1 := range nodes {
			for _, n2 := range nodes[i+1:] {
				edges = append(edges, edge{
					n1:       n1,
					n2:       n2,
					distance: n1.nodeDistance(n2),
				})
			}
		}
	}
	slices.SortFunc(edges, func(a, b edge) int {
		return int(a.distance - b.distance)
	})

	for _, e := range edges[:1000] {
		if !groups.SameSet(e.n1, e.n2) {
			groups.Union(e.n1, e.n2)
		}
	}

	lens := groups.Sizes()
	slices.Sort(lens)

	res := 1
	for i := range 3 {
		res *= lens[len(lens)-i-1]
	}

	fmt.Printf("Day8 Pt1 - Total: %d\n", res)
}

func Part2() {
	lineIter := util.FileScanner("./day8/input.txt")

	groups := util.NewDSU[node]()
	edges := make([]edge, 0, 1000000)
	{
		nodes := make([]node, 0, 1000)
		for line := range lineIter {
			parts := strings.Split(line, ",")
			node := node{
				x: util.QuickAtoi(parts[0]),
				y: util.QuickAtoi(parts[1]),
				z: util.QuickAtoi(parts[2]),
			}

			groups.Find(node)
			nodes = append(nodes, node)
		}

		for i, n1 := range nodes {
			for _, n2 := range nodes[i+1:] {
				edges = append(edges, edge{
					n1:       n1,
					n2:       n2,
					distance: n1.nodeDistance(n2),
				})
			}
		}
	}
	slices.SortFunc(edges, func(a, b edge) int {
		return int(a.distance - b.distance)
	})

	res := uint64(0)
	for _, e := range edges {
		if !groups.SameSet(e.n1, e.n2) {
			groups.Union(e.n1, e.n2)

			if groups.SetCount() == 1 {
				res = uint64(e.n1.x) * uint64(e.n2.x)
				break
			}
		}
	}

	fmt.Printf("Day8 Pt2 - Total: %d\n", res)
}
