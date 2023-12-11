package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type node struct {
	location string
	left     string
	right    string
}

var start, end = "AAA", "ZZZ"

func Run1() {
	file, err := os.Open("./day8/input")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	scan.Scan()
	instructions := scan.Text()

	nodeMap := make(map[string]node)

	scan.Scan()
	for scan.Scan() {
		line := scan.Text()
		location := line[0:3]
		nodeMap[location] = node{
			location: location,
			left:     line[7:10],
			right:    line[12:15],
		}
	}

	steps, currInstruction := 0, 0
	currNode := nodeMap[start]
	for currNode.location != end {
		if currInstruction == len(instructions) {
			currInstruction = 0
		}

		if instructions[currInstruction] == 'L' {
			currNode = nodeMap[currNode.left]
		} else {
			currNode = nodeMap[currNode.right]
		}

		currInstruction++
		steps++
	}

	fmt.Printf("Day8 Pt1 - Steps taken: %d", steps)
}
