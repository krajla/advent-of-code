package day7

import (
	"adventofcode/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var cards1 = []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

type handBid struct {
	hand  string
	bid   int
	power int
}

func Run1() {
	file, err := os.Open("./day7/input")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	handBids := make([]handBid, 0)
	for scan.Scan() {
		line := strings.Split(scan.Text(), " ")
		hand := line[0]
		bid := util.QuickAtoi(line[1])
		power := calcHandPower1(hand)

		handBids = append(handBids, handBid{
			hand:  hand,
			bid:   bid,
			power: power,
		})
	}

	slices.SortFunc(handBids, func(a, b handBid) int {
		if a.power != b.power {
			return a.power - b.power
		}
		return cmpFirstCard1(a.hand, b.hand)
	})

	totalEarning := 0
	for i, handBid := range handBids {
		totalEarning += handBid.bid * (len(handBids) - i)
	}
	fmt.Printf("Day7 Pt1 - Total earnings: %d\n", totalEarning)
}

func calcHandPower1(hand string) int {
	// Poor man's max heap.
	max1, max2 := 0, 0
	maxCard := byte('.')
	freq := make(map[byte]int, len(cards1))
	for i := 0; i < len(hand); i++ {
		freq[hand[i]]++
		if freq[hand[i]] >= max1 {
			// Avoids max1 and max2 referring to the same card in hand
			if maxCard != hand[i] {
				max2 = max1
			}

			max1 = freq[hand[i]]
			maxCard = hand[i]
		} else if freq[hand[i]] > max2 {
			max2 = freq[hand[i]]
		}
	}

	switch {
	case max1 == 5:
		return 1
	case max1 == 4:
		return 2
	case max1 == 3 && max2 == 2:
		return 3
	case max1 == 3:
		return 4
	case max1 == 2 && max2 == 2:
		return 5
	case max1 == 2:
		return 6
	default:
		return 7
	}
}

func cmpFirstCard1(a, b string) int {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return slices.Index(cards1, a[i]) - slices.Index(cards1, b[i])
		}
	}

	log.Fatal("HoHoHo")
	return 0
}
