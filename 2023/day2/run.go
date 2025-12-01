package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	reader, err := os.Open("./day2/input")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	counts := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		gameParts := strings.Split(line, ":")
		rounds := strings.Split(gameParts[1], ";")

		counts["red"] = 0
		counts["green"] = 0
		counts["blue"] = 0
		for _, round := range rounds {
			draws := strings.Split(strings.TrimSpace(round), ",")
			for _, draw := range draws {
				parts := strings.Split(strings.TrimSpace(draw), " ")
				count, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatal(err)
				}

				color := parts[1]
				if max := counts[color]; count > max {
					counts[color] = count
				}
			}
		}

		power := 1
		power *= counts["red"]
		power *= counts["green"]
		power *= counts["blue"]
		sum += power
	}

	fmt.Printf("Games sum: %d\n", sum)
}
