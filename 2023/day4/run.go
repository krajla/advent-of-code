package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	reader, err := os.Open("./day4/input")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	sum := 0
	noCards := 0
	multiplier := make([]int, 100)
	cardNo := 1
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		winning := line[0]
		numbers := line[1]
		winning = winning[strings.Index(winning, ":")+2 : len(winning)-1]
		numbers = numbers[1:]

		winningSet := make(map[int]struct{}, 0)
		for i := 0; i < len(winning); i += 3 {
			num, err := strconv.Atoi(strings.TrimSpace(winning[i : i+2]))
			if err != nil {
				log.Fatal(err)
			}
			winningSet[num] = struct{}{}
		}

		hits := 0
		for i := 0; i < len(numbers); i += 3 {
			num, err := strconv.Atoi(strings.TrimSpace(numbers[i : i+2]))
			if err != nil {
				log.Fatal(err)
			}
			if _, ok := winningSet[num]; ok {
				hits++
			}
		}

		cardMultiplier := multiplier[cardNo]
		cardMultiplier++

		for i := 1; i <= hits; i++ {
			if cardNo+i > len(multiplier)-1 {
				newMultiplier := make([]int, len(multiplier)*2)
				copy(newMultiplier, multiplier)
				multiplier = newMultiplier
			}
			multiplier[cardNo+i] += cardMultiplier
		}

		if hits > 0 {
			sum += 1 << (hits - 1)
		}
		noCards += cardMultiplier
		cardNo++
	}

	fmt.Printf("Card point sum: %d\n", sum)
	fmt.Printf("Number of cards: %d\n", noCards)
}
