package day9

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func area(a, b point) int {
	return (abs(b.x-a.x) + 1) * (abs(b.y-a.y) + 1)
}

type point struct {
	x int
	y int
}

func Part1() {
	lineIter := util.FileScanner("./day9/input.txt")

	allPoints := make([]point, 0, 1000)
	for line := range lineIter {
		parts := strings.Split(line, ",")
		x := util.QuickAtoi(parts[0])
		y := util.QuickAtoi(parts[1])
		p := point{
			x: x,
			y: y,
		}

		allPoints = append(allPoints, p)
	}

	res := 0
	for i, p1 := range allPoints {
		for _, p2 := range allPoints[i+1:] {
			area := area(p1, p2)
			if area > res {
				res = area
			}
		}
	}

	fmt.Printf("Day9 Pt1 - Total: %d\n", res)
}

func pointOnPolygonEdge(p point, polygon []point) bool {
	for i := range polygon {
		p1 := polygon[i]
		p2 := polygon[(i+1)%len(polygon)]

		if p1.x == p2.x && p.x == p1.x {
			if p.y >= util.Min(p1.y, p2.y) && p.y <= util.Max(p1.y, p2.y) {
				return true
			}
		} else if p1.y == p2.y && p.y == p1.y {
			if p.x >= util.Min(p1.x, p2.x) && p.x <= util.Max(p1.x, p2.x) {
				return true
			}
		} else {
			cross := (p2.x-p1.x)*(p.y-p1.y) - (p2.y-p1.y)*(p.x-p1.x)
			if cross == 0 {
				if p.x >= util.Min(p1.x, p2.x) && p.x <= util.Max(p1.x, p2.x) &&
					p.y >= util.Min(p1.y, p2.y) && p.y <= util.Max(p1.y, p2.y) {
					return true
				}
			}
		}
	}
	return false
}

func pointInPolygon(p point, polygon []point) bool {
	if pointOnPolygonEdge(p, polygon) {
		return true
	}

	if len(polygon) < 3 {
		return false
	}

	inside := false
	j := len(polygon) - 1

	for i := range polygon {
		p1 := polygon[i]
		p2 := polygon[j]

		if p1.y != p2.y {
			if (p1.y > p.y) != (p2.y > p.y) {
				xIntersect := float64(p1.x) + float64(p2.x-p1.x)*float64(p.y-p1.y)/float64(p2.y-p1.y)
				if float64(p.x) < xIntersect {
					inside = !inside
				}
			}
		}
		j = i
	}

	return inside
}

func crosses(a1, a2, b1, b2 point) bool {
	ccw := func(p1, p2, p3 point) int {
		val := (p2.x-p1.x)*(p3.y-p1.y) - (p2.y-p1.y)*(p3.x-p1.x)
		if val > 0 {
			return 1
		}
		if val < 0 {
			return -1
		}
		return 0
	}

	d1 := ccw(b1, b2, a1)
	d2 := ccw(b1, b2, a2)
	d3 := ccw(a1, a2, b1)
	d4 := ccw(a1, a2, b2)

	return ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) &&
		((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0))
}

func polygonEdgeCrossesRectangle(rectP1, rectP2, segP1, segP2 point) bool {
	minX, maxX := util.Min(rectP1.x, rectP2.x), util.Max(rectP1.x, rectP2.x)
	minY, maxY := util.Min(rectP1.y, rectP2.y), util.Max(rectP1.y, rectP2.y)

	if (segP1.x < minX && segP2.x < minX) || (segP1.x > maxX && segP2.x > maxX) {
		return false
	}
	if (segP1.y < minY && segP2.y < minY) || (segP1.y > maxY && segP2.y > maxY) {
		return false
	}

	onBoundary := func(p point) bool {
		return ((p.x == minX || p.x == maxX) && p.y >= minY && p.y <= maxY) ||
			((p.y == minY || p.y == maxY) && p.x >= minX && p.x <= maxX)
	}

	if onBoundary(segP1) && onBoundary(segP2) {
		if (segP1.x == segP2.x && (segP1.x == minX || segP1.x == maxX)) ||
			(segP1.y == segP2.y && (segP1.y == minY || segP1.y == maxY)) {
			return false
		}
	}

	rectEdges := [][]point{
		{{minX, minY}, {maxX, minY}}, // bottom
		{{maxX, minY}, {maxX, maxY}}, // right
		{{maxX, maxY}, {minX, maxY}}, // top
		{{minX, maxY}, {minX, minY}}, // left
	}

	for _, edge := range rectEdges {
		if crosses(segP1, segP2, edge[0], edge[1]) {
			return true
		}
	}

	return false
}

func rectangleInsidePolygon(p1, p2 point, polygon []point) bool {
	if p1.x == p2.x || p1.y == p2.y {
		return false
	}

	corners := []point{
		{p1.x, p1.y},
		{p1.x, p2.y},
		{p2.x, p1.y},
		{p2.x, p2.y},
	}

	// Check all corners are inside polygon
	for _, corner := range corners {
		if !pointInPolygon(corner, polygon) {
			return false
		}
	}

	// Check if any polygon edge crosses the rectangle
	for i := range polygon {
		next := (i + 1) % len(polygon)
		if polygonEdgeCrossesRectangle(p1, p2, polygon[i], polygon[next]) {
			return false
		}
	}

	return true
}

func Part2() {
	lineIter := util.FileScanner("./day9/input.txt")

	allPoints := make([]point, 0, 1000)
	for line := range lineIter {
		parts := strings.Split(line, ",")
		x := util.QuickAtoi(parts[0])
		y := util.QuickAtoi(parts[1])
		p := point{
			x: x,
			y: y,
		}

		allPoints = append(allPoints, p)
	}

	res := 0
	var bestP1, bestP2 point
	for i, p1 := range allPoints {
		for _, p2 := range allPoints[i+1:] {
			if rectangleInsidePolygon(p1, p2, allPoints) {
				rectArea := area(p1, p2)
				if rectArea > res {
					res = rectArea
					bestP1, bestP2 = p1, p2
				}
			}
		}
	}

	fmt.Printf("Day9 Pt2 - Total: %d (from {%d %d} to {%d %d})\n", res, bestP1.x, bestP1.y, bestP2.x, bestP2.y)
}
