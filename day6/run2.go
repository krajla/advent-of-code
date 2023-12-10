package day6

import (
	"adventofcode/util"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func Run2() {
	file, err := os.Open("./day6/input")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	race := race{}

	scan.Scan()
	times := splitParseLine(scan.Text())
	scan.Scan()
	distances := splitParseLine(scan.Text())
	for i := 0; i < len(times); i++ {
		time, distance := times[i], distances[i]

		race.time *= int(math.Pow10(util.LenInt(time)))
		race.time += time
		race.distance *= int(math.Pow10(util.LenInt(distance)))
		race.distance += distance
	}

	optimal := race.time / 2
	lowerLimit := bSearchLimit(0, optimal, race.time, race.distance)
	upperLimit := race.time - lowerLimit
	limitRange := upperLimit - lowerLimit + 1

	fmt.Printf("Limit full race: %d\n", limitRange)
}
