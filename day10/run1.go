package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type cord struct {
	x int
	y int
}

func (a cord) Equals(b cord) bool {
	return a.x == b.x && a.y == b.y
}

func (a cord) InLimits(limit cord) bool {
	return a.x >= 0 && a.x < limit.x && a.y >= 0 && a.y < limit.y
}

func (a cord) Sum(b cord) cord {
	return cord{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

type pipeLayout struct {
	a cord
	b cord
}

var pipes = map[byte]pipeLayout{
	'|': {
		a: cord{x: 1},
		b: cord{x: -1},
	},
	'-': {
		a: cord{y: 1},
		b: cord{y: -1},
	},
	'L': {
		a: cord{x: -1},
		b: cord{y: 1},
	},
	'J': {
		a: cord{x: -1},
		b: cord{y: -1},
	},
	'7': {
		a: cord{x: 1},
		b: cord{y: -1},
	},
	'F': {
		a: cord{x: 1},
		b: cord{y: 1},
	},
}

func Run1() {
	file, err := os.Open("./day10/input")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	maze := make([][]byte, 0)
	for scan.Scan() {
		maze = append(maze, []byte(scan.Text()))
	}

	start := cord{}
outLoop:
	for i, ba := range maze {
		for j, b := range ba {
			if b == byte('S') {
				start.x = i
				start.y = j
				break outLoop
			}
		}
	}

	// Determine the 2 starting paths.
	limit := cord{x: len(maze), y: len(maze[0])}
	prev := make([]cord, 0, 2)
	paths := make([]cord, 0, 2)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			curr := start.Sum(cord{x: i, y: j})
			if (i == 0 || j == 0) && curr.InLimits(limit) {
				next, ok := nextCord(maze[curr.x][curr.y], start, curr, limit)
				if ok {
					prev = append(prev, curr)
					paths = append(paths, next)
				}
			}
		}
	}

	distance := 2
	for !paths[0].Equals(paths[1]) {
		next, _ := nextCord(maze[paths[0].x][paths[0].y], prev[0], paths[0], limit)
		prev[0] = paths[0]
		paths[0] = next

		next, _ = nextCord(maze[paths[1].x][paths[1].y], prev[1], paths[1], limit)
		prev[1] = paths[1]
		paths[1] = next
		distance++
	}

	fmt.Printf("Day10 Pt1 - Furthest distance in maze: %d\n", distance)
}

func nextCord(tile byte, prev, curr, limit cord) (cord, bool) {
	layout, ok := pipes[tile]
	if !ok {
		return cord{}, false
	}

	if curr.Sum(layout.a).Equals(prev) && curr.Sum(layout.b).InLimits(limit) {
		return curr.Sum(layout.b), true
	}
	if curr.Sum(layout.b).Equals(prev) && curr.Sum(layout.a).InLimits(limit) {
		return curr.Sum(layout.a), true
	}

	return cord{}, false
}
