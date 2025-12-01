package day9

import (
	"adventofcode/util"
	"bufio"
	"fmt"
	"log"
	"os"
)

func Run1() {
	file, err := os.Open("./day9/input")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	predictions := 0
	for scan.Scan() {
		values := util.SplitParseLine(scan.Text(), 0)

		penultimate := extrapolate(values)
		lastValue := penultimate + values[len(values)-1]

		predictions += lastValue
	}

	fmt.Printf("Day9 Pt1 - Predictions: %d\n", predictions)
}

func extrapolate(values []int) int {
	allZ := true
	diffValues := make([]int, len(values)-1)
	for i := 0; i < len(diffValues); i++ {
		diffValues[i] = values[i+1] - values[i]
		if diffValues[i] != 0 {
			allZ = false
		}
	}

	if allZ {
		return 0
	} else {
		return diffValues[len(diffValues)-1] + extrapolate(diffValues)
	}
}
