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
		instructions []direction
		nodeMap      = make(map[string][2]string)
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

		nextNodesStr := lineSplit[1]
		nextNodesStr = strings.Trim(nextNodesStr, "()")
		nextNodes := strings.Split(nextNodesStr, ", ")

		nodeMap[node] = [2]string{nextNodes[0], nextNodes[1]}
	}

	var (
		currentNode      = "AAA"
		instructionIndex = -1
		steps            = 0
	)

	for currentNode != "ZZZ" {
		instructionIndex++
		steps++

		if instructionIndex == len(instructions) {
			instructionIndex = 0
		} else if instructionIndex > len(instructions) {
			log.Fatalf("instruction index out of bounds: %d", instructionIndex)
		}

		direction := instructions[instructionIndex]

		if _, ok := nodeMap[currentNode]; !ok {
			log.Fatalf("node %s not found in map", currentNode)
		}

		currentNode = nodeMap[currentNode][direction]
	}

	fmt.Println(steps)
}
