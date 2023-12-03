package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
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
		previousNums    []*number
		previousSymbols = make(map[int][]*number)

		total int
	)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("unable to read line: %w", err)
		}

		line = strings.TrimSuffix(line, "\n")

		var (
			currentNums    []*number
			currentSymbols = make(map[int][]*number)

			currentNum = NewNumber()
		)

		for i, r := range line {
			// digit
			if unicode.IsDigit(r) {
				currentNum.update(i, int(r-'0'))

				// last char in the line, so add it as a current number
				if i == len(line)-1 && currentNum.isValid() {
					currentNums = append(currentNums, currentNum)
				}

				continue
			}

			// record if valid gear
			if r != '.' && r == '*' {
				currentSymbols[i] = []*number{}
			}

			// occurence of a valid symbol indicates end of a valid number
			if currentNum.isValid() {
				currentNums = append(currentNums, currentNum)
			}

			currentNum = NewNumber()
		}

		// compare previous numbers with current symbols
		for _, l := range previousNums {
			for i := l.start - 1; i <= l.end+1; i++ {
				if _, ok := currentSymbols[i]; ok {
					currentSymbols[i] = append(currentSymbols[i], l)
				}
			}
		}

		for _, l := range currentNums {
			// check with both previous and current symbols
			for i := l.start - 1; i <= l.end+1; i++ {
				if _, ok := currentSymbols[i]; ok {
					currentSymbols[i] = append(currentSymbols[i], l)
				}
				if _, ok := previousSymbols[i]; ok {
					previousSymbols[i] = append(previousSymbols[i], l)
				}
			}
		}

		// process previous symbols as both sides of it are now checked
		for _, l := range previousSymbols {
			if len(l) == 2 {
				total += l[0].value * l[1].value
			}
		}

		previousNums = currentNums
		previousSymbols = currentSymbols
	}

	// process last line symbols
	for _, l := range previousSymbols {
		if len(l) == 2 {
			total += l[0].value * l[1].value
		}
	}

	fmt.Println(total)
}

type number struct {
	start int
	end   int
	value int
}

func NewNumber() *number {
	return &number{
		start: -1,
		end:   -1,
	}
}

func (l *number) update(pos int, val int) {
	if l.start == -1 {
		l.start = pos
	}

	l.end = pos

	l.value = l.value*10 + val
}

func (l number) isValid() bool {
	return l.start != -1 && l.end != -1
}

func (l number) String() string {
	return fmt.Sprintf("start: %d, end: %d, value: %d\n", l.start, l.end, l.value)
}
