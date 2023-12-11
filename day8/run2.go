package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Run2() {
	file, err := os.Open("./day8/input")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	scan.Scan()
	instructions := scan.Text()

	nodeMap := make(map[string]node)
	startLocations := make([]string, 0)

	scan.Scan()
	for scan.Scan() {
		line := scan.Text()
		location := line[0:3]
		nodeMap[location] = node{
			location: location,
			left:     line[7:10],
			right:    line[12:15],
		}
		if location[len(location)-1] == 'A' {
			startLocations = append(startLocations, location)
		}
	}

	// For the aoc specific input:
	// Obs: All traversals of starting nodes, and starting nodes after a loop, take a number of instructions equal to a multiple of 263(the length of the instruction).
	// {HJZ MQC BKF}   19199 / 263
	// {SBZ DVT DSH}   11309 / 263
	// {RFZ KKR SMR}   17621 / 263
	// {VPZ PLF VQV}   20777 / 263
	// {ZZZ VBQ PQP}   16043 / 263
	// {PQZ TJF HXM}   15517 / 263

	// Meaning that they all create a consistent loop of consistent length - because they take the same path of Lefts and Rights.
	// This code expects that of the input data and calculates the LCM on the traversal loop lengths.
	steps := findLoopLength(nodeMap[startLocations[0]], nodeMap, instructions)
	for i := 1; i < len(startLocations); i++ {
		n := nodeMap[startLocations[i]]
		steps = lcm(steps, findLoopLength(n, nodeMap, instructions))
	}

	fmt.Printf("Day8 Pt2 - Steps taken: %d\n", steps)
	fmt.Printf("%d", len(instructions))
}

func findLoopLength(n node, nodeMap map[string]node, instructions string) int {
	steps, currInstruction := 0, 0
	totalInstructions := 0
	for n.location[len(n.location)-1] != 'Z' {
		if currInstruction == len(instructions) {
			currInstruction = 0
		}

		if instructions[currInstruction] == 'L' {
			n = nodeMap[n.left]
		} else {
			n = nodeMap[n.right]
		}

		currInstruction++
		totalInstructions++
		steps++
	}
	fmt.Printf("%v\t", n)
	fmt.Printf("%d\n", totalInstructions)

	return steps
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
