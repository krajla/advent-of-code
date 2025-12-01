package day11

import (
	"adventofcode/util"
	"bufio"
	"fmt"
	"log"
	"os"
)

func Run1() {
	file, err := os.Open("./day11/input")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	x, y := 0, 0
	galaxyX, galaxyY := make(map[int]struct{}), make(map[int]struct{})
	galaxyLocations := make(map[int]map[int]struct{})
	for scan.Scan() {
		line := scan.Text()
		x++

		for i, r := range line {
			if r == '#' {
				galaxyX[x-1] = struct{}{}
				galaxyY[i] = struct{}{}
				if _, ok := galaxyLocations[x-1]; !ok {
					galaxyLocations[x-1] = make(map[int]struct{})
				}
				galaxyLocations[x-1][i] = struct{}{}
			}
		}
		y = len(line)
	}

	// Expand on Y, going backwards to not overwrite
	expansions := 0
	for i := 0; i < y; i++ {
		if _, ok := galaxyY[i]; !ok {
			expansions++
		}
	}

	for i := y - 1; i > 0; i-- {
		if _, ok := galaxyY[i]; !ok {
			expansions--
			y++
		}

		for j := 0; j < x; j++ {
			if val, ok := galaxyLocations[j][i]; ok {
				delete(galaxyLocations[j], i)
				galaxyLocations[j][i+expansions] = val
			}
		}
	}

	// Expand on X, going backwards to not overwrite
	expansions = 0
	for i := 0; i < x; i++ {
		if _, ok := galaxyX[i]; !ok {
			expansions++
		}
	}

	for i := x - 1; i > 0; i-- {
		if _, ok := galaxyX[i]; !ok {
			expansions--
			x++
		}

		if val, ok := galaxyLocations[i]; ok {
			delete(galaxyLocations, i)
			galaxyLocations[i+expansions] = val
		}
	}

	distances := 0
	pairs := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if _, ok := galaxyLocations[i][j]; ok {
				startJ := j + 1
				for searchI := i; searchI < x; searchI++ {
					for searchJ := startJ; searchJ < y; searchJ++ {
						if _, ok := galaxyLocations[searchI][searchJ]; ok {
							distances += distance(i, j, searchI, searchJ)
							pairs++
						}
					}
					startJ = 0
				}
			}
		}
	}

	fmt.Printf("Total galaxies distance: %d\n", distances)
}

func distance(ax, ay, bx, by int) int {
	return util.IntAbs(ax-bx) + util.IntAbs(ay-by)
}
