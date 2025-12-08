package day8

import (
	"adventofcode/util"
	"container/heap"
	"fmt"
	"math"
	"slices"
	"strings"
	"time"
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

type pq []edge

func (pq pq) Len() int           { return len(pq) }
func (pq pq) Less(i, j int) bool { return pq[i].distance < pq[j].distance }
func (pq pq) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *pq) Push(x any) {
	*pq = append(*pq, x.(edge))
}

func (pq *pq) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func Part1() time.Duration {
	lineIter := util.FileScanner("./day8/input.txt")

	groups := util.NewDSU[node]()
	h := &pq{}
	heap.Init(h)

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

	t := time.Now()
	for i, n1 := range nodes {
		for _, n2 := range nodes[i+1:] {
			edge := edge{
				n1:       n1,
				n2:       n2,
				distance: n1.nodeDistance(n2),
			}
			heap.Push(h, edge)
		}
	}

	for range 1000 {
		e := heap.Pop(h).(edge)
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

	elapsed := time.Since(t)
	fmt.Printf("Day8 Pt1 - Total: %d\n", res)
	return elapsed
}

func Part2() time.Duration {
	lineIter := util.FileScanner("./day8/input.txt")

	groups := util.NewDSU[node]()
	h := &pq{}
	heap.Init(h)

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

	t := time.Now()
	for i, n1 := range nodes {
		for _, n2 := range nodes[i+1:] {
			edge := edge{
				n1:       n1,
				n2:       n2,
				distance: n1.nodeDistance(n2),
			}
			heap.Push(h, edge)
		}
	}

	res := uint64(0)
	for {
		e := heap.Pop(h).(edge)

		if !groups.SameSet(e.n1, e.n2) {
			groups.Union(e.n1, e.n2)

			if groups.SetCount() == 1 {
				res = uint64(e.n1.x) * uint64(e.n2.x)
				break
			}
		}
	}

	elapsed := time.Since(t)
	fmt.Printf("Day8 Pt2 - Total: %d\n", res)
	return elapsed
}
