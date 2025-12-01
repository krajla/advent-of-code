package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var digits3 = map[string]rune{
	"one": '1',
	"two": '2',
	"six": '6',
}

var digits4 = map[string]rune{
	"four": '4',
	"five": '5',
	"nine": '9',
}

var digits5 = map[string]rune{
	"three": '3',
	"seven": '7',
	"eight": '8',
}

func Run() {
	reader, err := os.Open("./day1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var first, last rune
		for i := 0; i <= len(line); i++ {
			r := rune(line[i])
			if unicode.IsDigit(r) {
				first = r
				break
			}

			if i >= 2 {
				window := line[i-2 : i+1]
				if digit, ok := digits3[window]; ok {
					first = digit
					break
				}
			}

			if i >= 3 {
				window := line[i-3 : i+1]
				if digit, ok := digits4[window]; ok {
					first = digit
					break
				}
			}

			if i >= 4 {
				window := line[i-4 : i+1]
				if digit, ok := digits5[window]; ok {
					first = digit
					break
				}
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			r := rune(line[i])
			if '0' <= r && r <= '9' {
				last = r
				break
			}

			if i <= len(line)-3 {
				window := line[i : i+3]
				if digit, ok := digits3[window]; ok {
					last = digit
					break
				}
			}

			if i <= len(line)-4 {
				window := line[i : i+4]
				if digit, ok := digits4[window]; ok {
					last = digit
					break
				}
			}

			if i <= len(line)-5 {
				window := line[i : i+5]
				if digit, ok := digits5[window]; ok {
					last = digit
					break
				}
			}
		}
		number, err := strconv.Atoi(string(first) + string(last))
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}

	fmt.Printf("Sum: %d\n", sum)
}
