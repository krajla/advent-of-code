package day6

import (
	"adventofcode/util"
	"bufio"
	"fmt"
	"log"
	"os"
)

type race struct {
	time     int
	distance int
}

func Run1() {
	file, err := os.Open("./day6/input")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	races := make([]race, 0)

	scan.Scan()
	times := util.SplitParseLine(scan.Text(), 1)
	scan.Scan()
	distances := util.SplitParseLine(scan.Text(), 1)
	for i := 0; i < len(times); i++ {
		time, distance := times[i], distances[i]
		races = append(races, race{time: time, distance: distance})
	}

	result := 1
	for _, race := range races {
		optimal := race.time / 2
		lowerLimit := bSearchLimit(0, optimal, race.time, race.distance)
		upperLimit := race.time - lowerLimit

		limitRange := upperLimit - lowerLimit + 1
		result *= limitRange
	}

	fmt.Printf("Day6 Pt1 - Limits results: %d\n", result)
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
