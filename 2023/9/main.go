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

	var total int

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("unable to read line: %w", err)
		}

		line = strings.TrimSuffix(line, "\n")

		var nums []int

		for _, numStr := range strings.Fields(line) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal("unable to convert string to int: %w", err)
			}

			nums = append(nums, num)
		}

		var (
			currentNums = nums
			markerNums  = []int{nums[0]}
		)

		for {
			var (
				nextNums []int
				zeroNums = true
			)

			for i := 0; i < len(currentNums)-1; i++ {
				nextNum := currentNums[i+1] - currentNums[i]

				// track if all nums for next iteration are zero
				zeroNums = zeroNums && nextNum == 0

				// if first comparison, add to markerNums
				if i == 0 {
					markerNums = append(markerNums, nextNum)
				}

				nextNums = append(nextNums, nextNum)
			}

			// end loop if all nums are zero
			if zeroNums {
				break
			}

			currentNums = nextNums
		}

		// calculate the first number in the sequence
		historyPrediction := markerNums[len(markerNums)-1]
		for i := len(markerNums) - 2; i >= 0; i-- {
			historyPrediction = markerNums[i] - historyPrediction
		}

		total += historyPrediction
	}

	fmt.Println(total)
}
