// Bloat the graph by SCALAR value to "loosen" the loop so that the loop can be distinguished from the rest of the graph
//
// e.g.,
//
// In this given graph, it's difficult to distinguish the 8 points that APPEAR to be within the loop
// ..........
// .S------7.
// .|F----7|.
// .||....||.
// .||....||.
// .|L-7F-J|.
// .|..||..|.
// .L--JL--J.
// ..........
//
// If bloated by SCALAR=2, it becomes this
//
// where,
// ~ -> new "bloated" locations
// x -> plugged gaps that are part of the loop
//
// # This allows confirming that the 8 points are NOT part of the loop
//
// ~.~.~.~.~.~.~.~.~.~.~
// ~~~~~~~~~~~~~~~~~~~~~
// ~.~Sx-x-x-x-x-x-x7~.~
// ~~~x~~~~~~~~~~~~~x~~~
// ~.~|~Fx-x-x-x-x7~|~.~
// ~~~x~x~~~~~~~~~x~x~~~
// ~.~|~|~.~.~.~.~|~|~.~
// ~~~x~x~~~~~~~~~x~x~~~
// ~.~|~|~.~.~.~.~|~|~.~
// ~~~x~x~~~~~~~~~x~x~~~
// ~.~|~Lx-x7~Fx-xJ~|~.~
// ~~~x~~~~~x~x~~~~~x~~~
// ~.~|~.~.~|~|~.~.~|~.~
// ~~~x~~~~~x~x~~~~~x~~~
// ~.~Lx-x-xJ~Lx-x-xJ~.~
// ~~~~~~~~~~~~~~~~~~~~~
// ~.~.~.~.~.~.~.~.~.~.~
// ~~~~~~~~~~~~~~~~~~~~~
//
// Why?
// Bloating creates "islands" where each island is composed of ~ symbols, groundSymbol and non-loop pipes
// The 8 points are connected to the outside of the loop via ~ symbols, thus, they can be searched using DFS/BFS
//
// Reasoning:
// 1. The islands which will be OUTSIDE the loop can contain tiles (groundSymbol and non-loop pipes) but these WILL NOT be part of the answer
// 2. The ONLY island that will be INSIDE the loop will contain the tiles that WILL part of the answer
//
// Steps:
// 1. perform DFS/BFS on all symbols that are NOT in the loop to find these islands
// 2. the islands which CAN reach out-of-bounds must be the islands OUTSIDE the loop
// 3. the ONLY remaining island INSIDE the loop will never be out-of-bounds and the points within WILL be the answer
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	SCALAR = 2
)

type direction struct {
	x, y int
}

var (
	north = direction{-SCALAR, 0}
	south = direction{SCALAR, 0}
	east  = direction{0, SCALAR}
	west  = direction{0, -SCALAR}
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
	String() string
}

type (
	ground string
	start  string
)

func (s start) getNeighbours() []location {
	return []location{
		location(north),
		location(south),
		location(west),
		location(east),
	}
}

func (s start) enterFrom(d direction) bool {
	return true
}

func (s start) String() string {
	return string(s)
}

type misc string

type pipe struct {
	symbol string
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

func (p pipe) String() string {
	return p.symbol
}

type pipeJoin string

func (p pipeJoin) getNeighbours() []location {
	return []location{}
}

func (p pipeJoin) enterFrom(d direction) bool {
	return false
}

func (p pipeJoin) String() string {
	return string(p)
}

func newPipe(pipeStr string) (*pipe, error) {
	switch pipeStr {
	case "-":
		return &pipe{
			pipeType: horizontal,
			symbol:   string(pipeStr),
		}, nil
	case "|":
		return &pipe{
			pipeType: vertical,
			symbol:   string(pipeStr),
		}, nil
	case "L":
		return &pipe{
			pipeType: lBend,
			symbol:   string(pipeStr),
		}, nil
	case "J":
		return &pipe{
			pipeType: jBend,
			symbol:   string(pipeStr),
		}, nil
	case "7":
		return &pipe{
			pipeType: qBend,
			symbol:   string(pipeStr),
		}, nil
	case "F":
		return &pipe{
			pipeType: fBend,
			symbol:   string(pipeStr),
		}, nil
	default:
		return nil, fmt.Errorf("invalid pipe type: %s", pipeStr)
	}
}

const (
	groundSymbol   = '.'
	scalarSymbol   = '~'
	startSymbol    = 'S'
	pipeJoinSymbol = 'x'
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
		startLocation location
		graph         = make(map[location]moveable)
		surface       = make(map[location]misc)
	)

	lineLength := 0
	lineNumber := 0

	// bloat the area by SCALAR and insert scalarSymbol for empty spaces in surface map
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("unable to read line: %w", err)
		}

		line = strings.TrimSuffix(line, "\n")

		// add line of scalarSymbols at the start
		if lineNumber == 0 {
			for i := 0; i < len(line)*SCALAR+1; i++ {
				surface[location{1, i}] = misc(scalarSymbol)
			}
			lineNumber++
		}

		// iterate over characters by assuming spaces at the start and end and between all characters
		for i := 0; i < len(line)*SCALAR+1; i++ {
			x := lineNumber
			y := i

			// insert scalarSymbol in between given characters
			if i%2 == 0 {
				surface[location{x, y}] = misc(scalarSymbol)
				continue
			}

			c := line[i/2]

			switch c {
			case groundSymbol:
				surface[location{x, y}] = misc(groundSymbol)
			case startSymbol:
				graph[location{x, y}] = start(startSymbol)
				startLocation = location{x, y}
			default:
				p, err := newPipe(string(c))
				if err != nil {
					log.Fatalf("unable to create pipe: %v", err)
				}

				graph[location{x, y}] = p
			}
		}

		lineLength = len(line)

		// increase after processing given line
		lineNumber++

		// add line of scalarSymbols
		for i := 0; i < lineLength*SCALAR+1; i++ {
			surface[location{lineNumber, i}] = misc(scalarSymbol)
		}

		// ready for next input line
		lineNumber++
	}

	var (
		steps   int
		current = startLocation
		loop    = map[location]bool{
			current: true,
		}
	)

	// find the loop and delete the scalarSymbol from surface map
	for {
		for _, n := range graph[current].getNeighbours() {
			var ok bool

			newLocation := location{current.x + n.x, current.y + n.y}

			if _, ok = graph[newLocation]; !ok {
				continue
			}

			if ok = graph[newLocation].enterFrom(direction{-n.x, -n.y}); ok {
				// delete scalarSymbol in between the two pipes
				fillerLocation := location{current.x + (n.x / SCALAR), current.y + (n.y / SCALAR)}
				delete(surface, fillerLocation)

				// add pipeJoinSymbol in between the two pipes to "plug" the gap
				graph[fillerLocation] = pipeJoin(pipeJoinSymbol)

				current = newLocation

				// add the pipeJoinSymbol and the new location to the loop
				loop[fillerLocation] = true
				loop[current] = true

				break
			}
		}

		steps++

		if graph[current].String() == "S" {
			break
		}
	}

	// TODO: might not work if the enclosed tiles are just pipes
	// if this happens, then change the loop to do it over graph instead of surface
	var numPointsInLoop int
	visited := make(map[location]bool)

	for surfaceLoc, surfaceSym := range surface {
		if surfaceSym == misc(scalarSymbol) {
			continue
		}

		// add all tiles that are inside loop
		points, areInsideLoop := dfs(surfaceLoc, &visited, &loop, &graph, &surface)
		if areInsideLoop && points > 0 {
			numPointsInLoop += points
		}
	}

	fmt.Println(numPointsInLoop)
}

// dfs to find all points inside the loop
func dfs(
	currentLoc location,
	visited *map[location]bool,
	loop *map[location]bool,
	graph *map[location]moveable,
	surface *map[location]misc,
) (int, bool) {
	// if already visited, ignore
	if _, ok := (*visited)[currentLoc]; ok {
		return 0, true
	}

	// if loop, ignore
	if _, ok := (*loop)[currentLoc]; ok {
		return 0, true
	}

	var (
		currentSym misc
		points     int
	)

	_, inPipes := (*graph)[currentLoc]

	currentSym, inSurface := (*surface)[currentLoc]

	// if neither pipe nor surface, then it is out of bounds
	// out of bounds is IMPORTANT
	// it means that everything connecting to this location is NOT inside the loop
	if !inPipes && !inSurface {
		(*visited)[currentLoc] = false
		return 0, false
	}

	(*visited)[currentLoc] = true

	// if pipe or a ground symbol
	if inPipes || (inSurface && currentSym == misc(groundSymbol)) {
		points++
	}

	var anyNeighborOutside bool

	// check adjacent tiles, including bloated ones or out of bounds
	for _, loc := range []location{
		{currentLoc.x + 1, currentLoc.y},
		{currentLoc.x - 1, currentLoc.y},
		{currentLoc.x, currentLoc.y + 1},
		{currentLoc.x, currentLoc.y - 1},
	} {
		newPoints, areInsideLoop := dfs(loc, visited, loop, graph, surface)
		if areInsideLoop {
			points += newPoints
		} else {
			anyNeighborOutside = true
		}
	}

	// if any neighbor reported to be outside, this means that this location is also outside
	if anyNeighborOutside {
		(*visited)[currentLoc] = false
		return 0, false
	}

	return points, true
}

// print bloated graph for debugging
func printGraph(
	numLines,
	lineLength int,
	graph *map[location]moveable,
	surface *map[location]misc,
) {
	for i := 1; i < numLines; i++ {
		for j := 0; j < lineLength*SCALAR+1; j++ {
			if sym, ok := (*surface)[location{i, j}]; ok {
				fmt.Print(string(sym))
			} else if m, ok := (*graph)[location{i, j}]; ok {
				fmt.Print(m)
			} else {
				log.Fatal("invalid symbol")
			}
		}
		fmt.Println()
	}
}
