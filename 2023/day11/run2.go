package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const magnitude = 10

type cord struct {
	x int
	y int
}

func Run2() {
	file, err := os.Open("./day11/input2")
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
			expansions += magnitude
		}
	}

	for i := y - 1; i > 0; i-- {
		if _, ok := galaxyY[i]; !ok {
			expansions -= magnitude
			y += magnitude
		}

		for j := 0; j < x; j++ {
			if val, ok := galaxyLocations[j][i]; ok && expansions > 0 {
				galaxyLocations[j][i+expansions] = val
				delete(galaxyLocations[j], i)
			}
		}
	}

	// Expand on X, going backwards to not overwrite
	expansions = 0
	for i := 0; i < x; i++ {
		if _, ok := galaxyX[i]; !ok {
			expansions += magnitude
		}
	}

	for i := x - 1; i > 0; i-- {
		if _, ok := galaxyX[i]; !ok {
			expansions -= magnitude
			x += magnitude
		}

		if val, ok := galaxyLocations[i]; ok && expansions > 0 {
			galaxyLocations[i+expansions] = val
			delete(galaxyLocations, i)
		}
	}

	fmt.Printf(" %v\n", galaxyLocations)

	galaxyArray := make([]cord, 0)
	for i, v := range galaxyLocations {
		for j, _ := range v {
			galaxyArray = append(galaxyArray, cord{
				x: i,
				y: j,
			})
		}
	}

	distances := 0
	pairs := 0
	for i := 0; i < len(galaxyArray); i++ {
		a := galaxyArray[i]
		for j := i + 1; j < len(galaxyArray); j++ {
			b := galaxyArray[j]
			distances += distance(a.x, a.y, b.x, b.y)
			pairs++
		}
	}

	fmt.Printf("Total galaxies distance: %d\n", distances)
	fmt.Printf("Total galaxies pairs: %d\n", pairs)
}
