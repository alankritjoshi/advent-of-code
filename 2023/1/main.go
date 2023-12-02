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

var word_to_digit map[string]rune = map[string]rune{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
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

	total := 0

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("unable to read line: %w", err)
		}

		var (
			twoDigits = TwoDigits{
				first:  -1,
				second: -1,
			}
			word strings.Builder
		)

		for _, c := range line {
			var digit rune

			if unicode.IsDigit(c) {
				word.Reset()

				digit = c - '0'
			} else {
				word.WriteRune(c)

				var ok bool

				digit, ok = twoDigits.maybeGetDigit(word.String())
				if !ok {
					continue
				}
			}

			twoDigits.updateDigit(digit)
		}

		lineTotal := twoDigits.Int()

		total += lineTotal
	}

	fmt.Println(total)
}

type TwoDigits struct {
	first  rune
	second rune
}

func (twoDigits *TwoDigits) updateDigit(digit rune) {
	if twoDigits.first == -1 {
		twoDigits.first = digit
	}
	twoDigits.second = digit
}

func (twoDigits *TwoDigits) maybeGetDigit(searchStr string) (rune, bool) {
	for w, digit := range word_to_digit {
		if strings.HasSuffix(searchStr, w) {
			twoDigits.updateDigit(digit)

			return digit, true
		}
	}

	return -1, false
}

func (twoDigits *TwoDigits) String() string {
	if twoDigits.first == -1 {
		return ""
	}

	if twoDigits.second == -1 {
		return fmt.Sprintf("%d%d\n", int(twoDigits.first), int(twoDigits.first))
	}

	return fmt.Sprintf("%d%d\n", int(twoDigits.first), int(twoDigits.second))
}

func (twoDigits *TwoDigits) Int() int {
	if twoDigits.first == -1 || twoDigits.second == -1 {
		return 0
	}

	return int(twoDigits.first)*10 + int(twoDigits.second)
}
