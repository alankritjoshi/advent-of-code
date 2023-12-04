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

	hand := make(map[int]int)

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

		cardNum, _ := strconv.Atoi(strings.Fields(lineSplit[0])[1])

		hand[cardNum] += 1

		cardsSplit := strings.Split(lineSplit[1], "|")

		current := make(map[int]bool)

		for _, cardInHand := range strings.Fields(cardsSplit[1]) {
			cardVal, _ := strconv.Atoi(cardInHand)

			current[cardVal] = true
		}

		var cardTotal int

		for _, winningCard := range strings.Fields(cardsSplit[0]) {
			cardVal, _ := strconv.Atoi(winningCard)

			if _, ok := current[cardVal]; ok {
				cardTotal += 1
			}
		}

		for i := 1; i <= cardTotal; i++ {
			hand[cardNum+i] += hand[cardNum]
		}
	}

	var total int

	for _, v := range hand {
		total += v
	}

	fmt.Println(total)
}
