package util

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

func LenInt(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func QuickAtoi(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func SplitParseLine(line string, skip int) []int {
	parts := strings.Split(line, " ")
	parts = slices.DeleteFunc(parts, func(s string) bool {
		return s == ""
	})

	elements := make([]int, 0, len(parts))
	for i := skip; i < len(parts); i++ {
		el := QuickAtoi(strings.TrimSpace(parts[i]))
		elements = append(elements, el)
	}
	return elements
}

func IntAbs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
