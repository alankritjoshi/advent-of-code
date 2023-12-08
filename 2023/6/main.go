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

	var race race

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("unable to read line: %w", err)
		}

		lineSplit := strings.Fields(line)

		var numStrBuilder strings.Builder

		for _, numPartStr := range lineSplit[1:] {
			numStrBuilder.WriteString(numPartStr)
		}

		num, err := strconv.Atoi(numStrBuilder.String())
		if err != nil {
			log.Fatalf("unable to convert string to int: %v", err)
		}

		if lineSplit[0] == timeKeyword {
			race.time = num
		} else {
			race.distance = num
		}
	}

	ways := 0

	for t := 1; t < race.time; t++ {
		if (race.time-t)*t > race.distance {
			ways++
		}
	}

	fmt.Println(ways)
}
