package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type direction int

const (
	left  direction = 0
	right direction = 1
)

type nodeMarker struct {
	node   string
	marker int
}

func main() {
	input := flag.String("i", "", "Input File Name")

	flag.Parse()

	if *input == "" {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(*input)
	if err != nil {
		log.Fatal("unable to open file: %w", err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	var (
		instructions   []direction
		nodeMap        = make(map[string][2]string)
		startingNodes  = make([]string, 0)
		endingNodesMap = make(map[string]bool)
	)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("unable to read line: %w", err)
		}

		line = strings.TrimSuffix(line, "\n")

		if !strings.Contains(line, "=") {
			for _, c := range line {
				switch c {
				case 'L':
					instructions = append(instructions, left)
				case 'R':
					instructions = append(instructions, right)
				}
			}

			continue
		}

		lineSplit := strings.Split(line, " = ")

		node := lineSplit[0]

		if strings.HasSuffix(node, "A") {
			startingNodes = append(startingNodes, node)
		}

		if strings.HasSuffix(node, "Z") {
			endingNodesMap[node] = true
		}

		nextNodesStr := lineSplit[1]
		nextNodesStr = strings.Trim(nextNodesStr, "()")
		nextNodes := strings.Split(nextNodesStr, ", ")

		nodeMap[node] = [2]string{nextNodes[0], nextNodes[1]}
	}

	endings := make([]int, len(startingNodes))

	// for each node, find the steps required to reach ending node until we starting loop for the same instruction
	for i, startingNode := range startingNodes {
		var (
			currentNode = startingNode
			visitedMap  = make(map[string]bool)

			steps            = 1
			instructionIndex int
		)

		for {
			// reset instruction index if we reach the end of the instructions
			if instructionIndex == len(instructions) {
				instructionIndex = 0
			}

			direction := instructions[instructionIndex]

			currentNode = nodeMap[currentNode][direction]

			// node-instruction to check if we've visited it already
			_, ok := visitedMap[fmt.Sprintf("%s-%d", currentNode, instructionIndex)]
			if ok {
				break
			}

			if !ok {
				visitedMap[fmt.Sprintf("%s-%d", currentNode, instructionIndex)] = true
			}

			// record the ending when found
			_, ok = endingNodesMap[currentNode]
			if ok {
				endings[i] = steps
			}

			instructionIndex++
			steps++
		}
	}

	// lcm of all the steps required to reach the ending nodes
	fmt.Println(lcm(endings[0], endings[1], endings[2:]...))
}

// find Greatest Common Divisor (gcd) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (lcm) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
