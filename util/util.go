package util

import (
	"log"
	"strconv"
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
