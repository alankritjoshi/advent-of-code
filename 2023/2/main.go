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

	validIDTotal := 0

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

		game := game{
			id: gameID,
		}

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

			if !game.isValid() {
				break
			}

			game.reset()
		}

		if game.isValid() {
			validIDTotal += game.id
		}
	}

	fmt.Println(validIDTotal)
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
	id int

	red   int
	green int
	blue  int
}

func NewGame(id int) *game {
	return &game{
		id: id,
	}
}

func (g *game) addBall(color string, count int) error {
	c, ok := colorMap[color]
	if !ok {
		return fmt.Errorf("failed to find color: %s", c)
	}

	switch color {
	case blue:
		g.blue += count
	case green:
		g.green += count
	case red:
		g.red += count
	}

	return nil
}

func (g game) isValid() bool {
	if g.red > 12 || g.green > 13 || g.blue > 14 {
		return false
	}
	return true
}

func (g *game) reset() {
	g.red = 0
	g.green = 0
	g.blue = 0
}

func (g game) String() string {
	return fmt.Sprintf("Game %5d - red:%5d     green:%5d     blue:%5d", g.id, g.red, g.green, g.blue)
}
