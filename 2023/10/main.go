package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type direction struct {
	x, y int
}

var (
	north = direction{0, 1}
	south = direction{0, -1}
	east  = direction{1, 0}
	west  = direction{-1, 0}
)

type pipeType struct {
	ends [2]direction
}

var (
	horizontal = pipeType{[2]direction{east, west}}
	vertical   = pipeType{[2]direction{north, south}}
	lBend      = pipeType{[2]direction{north, east}}
	jBend      = pipeType{[2]direction{north, west}}
	qBend      = pipeType{[2]direction{south, west}}
	fBend      = pipeType{[2]direction{south, east}}
)

type location struct {
	x, y int
}

type moveable interface {
	getNeighbours(x, y int) []location
	canEnterFrom(d direction) bool
}

type pipe struct {
	pipeType
}

func (p pipe) getNeighbours(x, y int) []location {
	var neighbours []location

	for _, d := range p.ends {
		neighbours = append(neighbours, location{x + d.x, y + d.y})
	}

	return neighbours
}

func (p pipe) canEnterFrom(d direction) bool {
	for _, e := range p.ends {
		if e == d {
			return true
		}
	}

	return false
}

type (
	ground string
	start  string
)

func (s start) getNeighbours(x, y int) []location {
	return []location{
		{x + 1, y},
		{x - 1, y},
		{x, y + 1},
		{x, y - 1},
	}
}

func (s start) canEnterFrom(d direction) bool {
	return true
}

const (
	groundSymbol = '.'
	startSymbol  = 'S'
)

func newPipe(pipeStr string) (*pipe, error) {
	switch pipeStr {
	case "-":
		return &pipe{horizontal}, nil
	case "|":
		return &pipe{vertical}, nil
	case "L":
		return &pipe{lBend}, nil
	case "J":
		return &pipe{jBend}, nil
	case "7":
		return &pipe{qBend}, nil
	case "F":
		return &pipe{fBend}, nil
	default:
		return nil, fmt.Errorf("invalid pipe type: %s", pipeStr)
	}
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
		startLocation location
		graph         = make(map[location]moveable)
	)

	lineNumber := 0

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("unable to read line: %w", err)
		}

		line = strings.TrimSuffix(line, "\n")

		for i, c := range line {
			switch c {
			case groundSymbol:
				continue
			case startSymbol:
				graph[location{lineNumber, i}] = start(startSymbol)
				startLocation = location{lineNumber, i}
			default:
				p, err := newPipe(string(c))
				if err != nil {
					log.Fatalf("unable to create pipe: %v", err)
				}

				graph[location{lineNumber, i}] = p
			}
		}

		lineNumber++
	}

	var (
		steps   int
		current = startLocation
	)

	for steps == 0 || current != startLocation {
		for _, n := range graph[current].getNeighbours(current.x, current.y) {
			var o moveable
			var ok bool

			if o, ok = graph[n]; !ok {
				continue
			}

			fmt.Println(direction{n.x - current.x, n.y - current.y})

			if ok && o.canEnterFrom(direction{n.x - current.x, n.y - current.y}) {
				current = n
				steps++
			}
		}
	}

	fmt.Println(steps)
}
