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

const timeKeyword = "Time:"

type race struct {
	time     int
	distance int
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

	races := make([]race, 0)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("unable to read line: %w", err)
		}

		lineSplit := strings.Fields(line)

		for numIndex, numStr := range lineSplit[1:] {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatalf("unable to convert string to int: %v", err)
			}

			if lineSplit[0] == timeKeyword {
				races = append(races, race{time: num})
			} else {
				races[numIndex].distance = num
			}
		}
	}

	ways := 1

	for _, race := range races {
		raceWays := 0
		for t := 1; t < race.time; t++ {
			if (race.time-t)*t > race.distance {
				raceWays++
			}
		}
		ways *= raceWays
	}

	fmt.Println(ways)
}
