package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
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

	total := 0

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("unable to read line: %w", err)
		}

		line = strings.TrimSuffix(line, "\n")

		lineSplit := strings.Split(line, ":")

		// cardNum, _ := strconv.Atoi(strings.Split(lineSplit[0], " ")[1])

		cardsSplit := strings.Split(lineSplit[1], "|")

		hand := make(map[int]int)

		for _, cardInHand := range strings.Fields(cardsSplit[1]) {
			cardVal, _ := strconv.Atoi(cardInHand)

			hand[cardVal] += 1
		}

		var cardTotal int

		for _, winningCard := range strings.Split(strings.TrimSpace(cardsSplit[0]), " ") {
			cardVal, _ := strconv.Atoi(winningCard)

			if count, ok := hand[cardVal]; ok {
				if cardTotal == 0 {
					cardTotal = int(math.Pow(2, float64(count-1)))
				} else {
					cardTotal *= int(math.Pow(2, float64(count)))
				}
			}
		}

		total += cardTotal
	}

	fmt.Println(total)
}
