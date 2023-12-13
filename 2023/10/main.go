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
	north = direction{-1, 0}
	south = direction{1, 0}
	east  = direction{0, 1}
	west  = direction{0, -1}
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
	getNeighbours() []location
	enterFrom(d direction) bool
}

type pipe struct {
	pipeType
	status [2]bool
}

func (p pipe) getNeighbours() []location {
	var neighbours []location

	for i, d := range p.ends {
		if !p.status[i] {
			neighbours = append(neighbours, location(d))
		}
	}

	return neighbours
}

func (p *pipe) enterFrom(d direction) bool {
	for i, e := range p.ends {
		if !p.status[i] && e.x == d.x && e.y == d.y {
			p.status[i] = true
			return true
		}
	}

	return false
}

type (
	ground string
	start  string
)

func (s start) getNeighbours() []location {
	return []location{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
}

func (s start) enterFrom(d direction) bool {
	return false
}

const (
	groundSymbol = '.'
	startSymbol  = 'S'
)

func newPipe(pipeStr string) (*pipe, error) {
	switch pipeStr {
	case "-":
		return &pipe{
			pipeType: horizontal,
		}, nil
	case "|":
		return &pipe{
			pipeType: vertical,
		}, nil
	case "L":
		return &pipe{
			pipeType: lBend,
		}, nil
	case "J":
		return &pipe{
			pipeType: jBend,
		}, nil
	case "7":
		return &pipe{
			pipeType: qBend,
		}, nil
	case "F":
		return &pipe{
			pipeType: fBend,
		}, nil
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

	for {
		var nextFound bool

		for _, n := range graph[current].getNeighbours() {
			var ok bool

			newLocation := location{current.x + n.x, current.y + n.y}

			if _, ok = graph[newLocation]; !ok {
				continue
			}

			if ok = graph[newLocation].enterFrom(direction{-n.x, -n.y}); ok {
				current = newLocation
				nextFound = true
				break
			}
		}

		steps++

		if !nextFound {
			break
		}
	}

	fmt.Println(int(steps / 2))
}
