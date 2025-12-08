package util

import "iter"

type DSU[T comparable] struct {
	parent map[T]T
	rank   map[T]int
	size   map[T]int
}

func NewDSU[T comparable]() *DSU[T] {
	return &DSU[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
		size:   make(map[T]int),
	}
}

func (d *DSU[T]) Find(x T) T {
	if _, exists := d.parent[x]; !exists {
		d.parent[x] = x
		d.rank[x] = 0
		d.size[x] = 1
		return x
	}

	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU[T]) Union(x, y T) {
	rootX := d.Find(x)
	rootY := d.Find(y)

	if rootX == rootY {
		return
	}

	rankX := d.rank[rootX]
	rankY := d.rank[rootY]

	if rankX < rankY {
		rootX, rootY = rootY, rootX
		rankX, rankY = rankY, rankX
	}

	d.parent[rootY] = rootX
	d.size[rootX] += d.size[rootY]
	delete(d.rank, rootY)
	delete(d.size, rootY)

	if rankX == rankY {
		d.rank[rootX]++
	}
}

func (d *DSU[T]) SameSet(x, y T) bool {
	return d.Find(x) == d.Find(y)
}

func (d *DSU[T]) Size(x T) int {
	root := d.Find(x)
	return d.size[root]
}

func (d *DSU[T]) SetCount() int {
	return len(d.rank)
}

func (d *DSU[T]) Keys() []T {
	keys := make([]T, 0, len(d.parent))
	for k := range d.parent {
		keys = append(keys, k)
	}
	return keys
}

func (d *DSU[T]) Roots() []T {
	roots := make([]T, 0, len(d.rank))
	for root := range d.rank {
		roots = append(roots, root)
	}
	return roots
}

func (d *DSU[T]) Sizes() []int {
	sizes := make([]int, 0, len(d.rank))
	for root := range d.size {
		sizes = append(sizes, d.size[root])
	}
	return sizes
}

func (d *DSU[T]) GetSet(root T) []T {
	set := make([]T, 0)
	for k := range d.parent {
		if d.Find(k) == root {
			set = append(set, k)
		}
	}
	return set
}

func (d *DSU[T]) IterSets() iter.Seq2[T, []T] {
	return func(yield func(T, []T) bool) {
		for _, root := range d.Roots() {
			set := d.GetSet(root)
			if !yield(root, set) {
				return
			}
		}
	}
}
