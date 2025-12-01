package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Run() {
	reader, err := os.Open("./day3/input")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var prev, curr, next = 0, 1, 2
	lineBuffer := make([]string, 3)

	padding := ""
	scanner.Scan()
	lineBuffer[next] = scanner.Text()
	for i := 0; i < len(lineBuffer[next]); i++ {
		padding += "."
	}
	lineBuffer[curr] = padding

	gearMap := make(map[string][]int)
	sum := 0
	height := 0
	for scanner.Scan() {
		lineBuffer[prev] = lineBuffer[curr]
		lineBuffer[curr] = lineBuffer[next]
		lineBuffer[next] = scanner.Text()

		sum += checkNumbers(lineBuffer[prev], lineBuffer[curr], lineBuffer[next], height, gearMap)
		height++
	}
	sum += checkNumbers(lineBuffer[curr], lineBuffer[next], padding, height, gearMap)

	gearRatioSum := 0
	for _, values := range gearMap {
		if len(values) == 2 {
			gearRatioSum += values[0] * values[1]
		}
	}

	fmt.Printf("Part numbers sum: %d\n", sum)
	fmt.Printf("Gear ratio sum: %d\n", gearRatioSum)
}

func checkNumbers(prev, curr, next string, height int, gearMap map[string][]int) int {
	sum := 0
	startIndex := -1
	numberBuffer := ""
	for i := 0; i < len(curr); i++ {
		r := rune(curr[i])
		endIndex := -1
		if unicode.IsDigit(r) {
			if startIndex == -1 {
				startIndex = i
			}

			numberBuffer += string(r)
			if i == len(curr)-1 {
				endIndex = i
			}
		}

		if !unicode.IsDigit(r) && startIndex != -1 {
			endIndex = i - 1
		}

		if endIndex != -1 {
			number, err := strconv.Atoi(numberBuffer)
			if err != nil {
				log.Fatal(err)
			}

			isPartNumber, gearCoordinates := checkNeighbors(prev, curr, next, startIndex, endIndex, height)
			if isPartNumber {
				sum += number
			}
			for _, coord := range gearCoordinates {
				if _, ok := gearMap[coord]; !ok {
					gearMap[coord] = make([]int, 0, 1)
				}
				gearMap[coord] = append(gearMap[coord], number)
			}

			startIndex = -1
			numberBuffer = ""
		}
	}

	return sum
}

func checkNeighbors(prev, curr, next string, start, end, height int) (bool, []string) {
	start = max(start-1, 0)
	end = min(end+1, len(prev)-1)

	isPartNumber := false
	gearCoordinates := make([]string, 0)

	for i := start; i <= end; i++ {
		pr, cr, nr := rune(prev[i]), rune(curr[i]), rune(next[i])
		if !unicode.IsDigit(pr) && pr != '.' {
			isPartNumber = true

		}
		if pr == '*' {
			gearCoordinates = append(gearCoordinates, fmt.Sprintf("%d-%d", height-1, i))
		}

		if !unicode.IsDigit(cr) && cr != '.' {
			isPartNumber = true

		}
		if cr == '*' {
			gearCoordinates = append(gearCoordinates, fmt.Sprintf("%d-%d", height, i))
		}

		if !unicode.IsDigit(nr) && nr != '.' {
			isPartNumber = true
		}
		if nr == '*' {
			gearCoordinates = append(gearCoordinates, fmt.Sprintf("%d-%d", height+1, i))
		}
	}

	return isPartNumber, gearCoordinates
}
