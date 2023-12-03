package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	powerTotal := 0

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("unable to read line: %w", err)
		}

		line = strings.TrimSuffix(line, "\n")

		idBallsSplit := strings.Split(line, ": ")

		// get game ID
		gameWord := idBallsSplit[0]
		gameIDStr := strings.Split(gameWord, " ")[1]
		gameID, _ := strconv.Atoi(gameIDStr)

		game := NewGame(gameID)

		// get balls
		ballsLine := idBallsSplit[1]
		ballsSets := strings.Split(ballsLine, "; ")

		for _, ballsSet := range ballsSets {
			balls := strings.Split(ballsSet, ", ")

			for _, ball := range balls {
				ballSplit := strings.Split(ball, " ")
				numStr := ballSplit[0]
				numInt, _ := strconv.Atoi(numStr)
				color := ballSplit[1]

				err := game.addBall(color, numInt)
				if err != nil {
					log.Fatalf("unable to add ball: %v", err)
				}
			}

			game.reset()
		}

		powerTotal += game.power()
	}

	fmt.Println(powerTotal)
}

const (
	blue  = "blue"
	green = "green"
	red   = "red"
)

var colorMap = map[string]string{
	"blue":  blue,
	"green": green,
	"red":   red,
}

type game struct {
	red   ball
	green ball
	blue  ball
	id    int
}

func NewGame(id int) *game {
	return &game{
		id: id,
	}
}

func (g *game) addBall(colorStr string, count int) error {
	colorType, ok := colorMap[colorStr]
	if !ok {
		return fmt.Errorf("failed to find color: %s", colorStr)
	}

	switch colorType {
	case blue:
		g.blue.add(count)
	case green:
		g.green.add(count)
	case red:
		g.red.add(count)
	}

	return nil
}

func (g game) isValid() bool {
	return g.red.count <= 12 && g.green.count <= 13 && g.blue.count <= 14
}

func (g *game) reset() {
	g.red.reset()
	g.green.reset()
	g.blue.reset()
}

func (g game) power() int {
	if g.red.min == 0 && g.green.min == 0 && g.blue.min == 0 {
		return 0
	}

	var (
		red   = g.red.min
		green = g.green.min
		blue  = g.blue.min
	)

	if g.red.min == 0 {
		red = 1
	}

	if g.green.min == 0 {
		green = 1
	}

	if g.blue.min == 0 {
		blue = 1
	}

	return red * green * blue
}

func (g game) String() string {
	return fmt.Sprintf("Game %5d - red:%5d     green:%5d     blue:%5d     -     power:%5d", g.id, g.red, g.green, g.blue, g.power())
}

type ball struct {
	count int
	min   int
}

func (b *ball) add(count int) {
	b.count += count
	b.min = max(b.min, count)
}

func (b *ball) reset() {
	b.count = 0
}
