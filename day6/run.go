package day6

import (
	"adventofcode/util"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

type race struct {
	time     int
	distance int
}

func Run() {
	file, err := os.Open("./day6/input")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	racesP1 := make([]race, 0)
	raceP2 := race{}

	scan.Scan()
	times := splitParseLine(scan.Text())
	scan.Scan()
	distances := splitParseLine(scan.Text())
	for i := 0; i < len(times); i++ {
		time, distance := times[i], distances[i]
		racesP1 = append(racesP1, race{time: time, distance: distance})

		raceP2.time *= int(math.Pow10(util.LenInt(time)))
		raceP2.time += time
		raceP2.distance *= int(math.Pow10(util.LenInt(distance)))
		raceP2.distance += distance
	}

	result := 1
	for _, race := range racesP1 {
		optimal := race.time / 2
		lowerLimit := bSearchLimit(0, optimal, race.time, race.distance)
		upperLimit := race.time - lowerLimit

		limitRange := upperLimit - lowerLimit + 1
		result *= limitRange
	}

	optimal := raceP2.time / 2
	lowerLimit := bSearchLimit(0, optimal, raceP2.time, raceP2.distance)
	upperLimit := raceP2.time - lowerLimit
	limitRange := upperLimit - lowerLimit + 1

	fmt.Printf("Limits results: %d\n", result)
	fmt.Printf("Limit full race: %d\n", limitRange)
}

func bSearchLimit(start, end, time, distance int) int {
	if start == end {
		return start
	}

	mid := (start + end) / 2
	midDistance := calcDistance(time, mid)
	if midDistance == distance {
		return mid + 1
	}

	if distance < midDistance {
		return bSearchLimit(start, mid, time, distance)
	} else {
		return bSearchLimit(mid+1, end, time, distance)
	}
}

func calcDistance(time, windup int) int {
	return (time - windup) * windup
}

func splitParseLine(line string) []int {
	parts := strings.Split(line, " ")
	parts = slices.DeleteFunc(parts, func(s string) bool {
		return s == ""
	})

	elements := make([]int, 0, len(parts))
	for i := 1; i < len(parts); i++ {
		el := util.QuickAtoi(strings.TrimSpace(parts[i]))
		elements = append(elements, el)
	}
	return elements
}
