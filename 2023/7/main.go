package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type card struct {
	face  string
	value int
}

var (
	two   = card{"2", 2}
	three = card{"3", 3}
	four  = card{"4", 4}
	five  = card{"5", 5}
	six   = card{"6", 6}
	seven = card{"7", 7}
	eight = card{"8", 8}
	nine  = card{"9", 9}
	ten   = card{"T", 10}
	jack  = card{"J", 1} // weakened from 11 to 1 for part 2
	queen = card{"Q", 12}
	king  = card{"K", 13}
	ace   = card{"A", 14}
)

func stringToCard(s string) (card, error) {
	switch s {
	case "2":
		return two, nil
	case "3":
		return three, nil
	case "4":
		return four, nil
	case "5":
		return five, nil
	case "6":
		return six, nil
	case "7":
		return seven, nil
	case "8":
		return eight, nil
	case "9":
		return nine, nil
	case "T":
		return ten, nil
	case "J":
		return jack, nil
	case "Q":
		return queen, nil
	case "K":
		return king, nil
	case "A":
		return ace, nil
	default:
		return card{}, fmt.Errorf("invalid card string: %s", s)
	}
}

func (c card) String() string {
	return fmt.Sprintf(c.face)
}

type hand struct {
	cards []card
	bids  int
}

func (h hand) counts() map[card]int {
	cardCounts := map[card]int{}

	for _, c := range h.cards {
		cardCounts[c]++
	}

	return cardCounts
}

func (h hand) isInvalid() bool {
	return len(h.cards) != 5
}

func (h hand) isFiveOfAKind() bool {
	for _, c := range h.cards {
		if c.face == "J" {
			return false
		}
	}

	return len(h.counts()) == 1
}

func (h hand) isFourOfAKind() bool {
	cardCounts := h.counts()

	if len(cardCounts) != 2 {
		return false
	}

	for _, cardCount := range cardCounts {
		if cardCount == 4 {
			return true
		}
	}

	return false
}

func (h hand) isFullHouse() bool {
	cardCounts := h.counts()

	if len(cardCounts) != 2 {
		return false
	}

	for _, cardCount := range cardCounts {
		if cardCount != 3 && cardCount != 2 {
			return false
		}
	}

	return true
}

func (h hand) isThreeOfAKind() bool {
	cardCounts := h.counts()

	var hasJoker, hasPair bool

	for card, cardCount := range cardCounts {

		if cardCount == 3 {
			return true
		}

		if cardCount == 2 {
			hasPair = true
		}

		if card.face == "J" {
			hasJoker = true
		}
	}

	return hasPair && hasJoker
}

func (h hand) isTwoPair() bool {
	cardCounts := h.counts()

	var jokers, pairs int

	for card, cardCount := range cardCounts {
		if card.face == "J" {
			jokers = cardCount
		}
		if cardCount == 2 {
			pairs++
		}
	}

	if pairs == 2 {
		return true
	}

	if pairs == 1 && jokers >= 1 {
		return true
	}

	return false
}

func (h hand) isOnePair() bool {
	cardCounts := h.counts()

	var jokers int
	for card, cardCount := range cardCounts {
		if card.face == "J" {
			jokers = cardCount
		}
		if cardCount == 2 {
			return true
		}
	}

	if jokers >= 1 {
		return true
	}

	return false
}

func (h hand) isHighCard() bool {
	cardCounts := h.counts()

	for card, cardCount := range cardCounts {
		if card.face == "J" {
			return false
		}

		if cardCount > 1 {
			return false
		}
	}
}

func (h hand) strength() int {
	var strength int

	if h.isHighCard() {
		strength = 1
	} else if h.isOnePair() {
		strength = 2
	} else if h.isTwoPair() {
		strength = 3
	} else if h.isThreeOfAKind() {
		strength = 4
	} else if h.isFullHouse() {
		strength = 5
	} else if h.isFourOfAKind() {
		strength = 6
	} else if h.isFiveOfAKind() {
		strength = 7
	} else {
		strength = 0
	}

	return strength
}

func (h hand) String() string {
	var cardsStrBuilder strings.Builder
	for _, c := range h.cards {
		cardsStrBuilder.WriteString(c.String())
	}

	normalCards := cardsStrBuilder.String()

	cardsStrBuilder.Reset()

	cards := h.cards

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].value < cards[j].value
	})

	for _, c := range cards {
		cardsStrBuilder.WriteString(c.String())
	}

	sortedCards := cardsStrBuilder.String()

	return fmt.Sprintf("cards: %s, sorted cards: %s, bids: %d", normalCards, sortedCards, h.bids)
}

func newHand(handStr, bidsStr string) (*hand, error) {
	cards := make([]card, 0)

	for _, c := range handStr {
		card, err := stringToCard(string(c))
		if err != nil {
			return nil, fmt.Errorf("failed to parse card %s in hand: %w", string(c), err)
		}

		cards = append(cards, card)
	}

	bids, err := strconv.Atoi(bidsStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse bids %s as integer: %w", bidsStr, err)
	}

	return &hand{
		cards: cards,
		bids:  bids,
	}, nil
}

func sortHands(hands []*hand) {
	sort.Slice(hands, func(i, j int) bool {
		handIStrength := hands[i].strength()
		handJStrength := hands[j].strength()

		if handIStrength != handJStrength {
			return hands[i].strength() < hands[j].strength()
		}

		for k := 0; k < 5; k++ {
			if hands[i].cards[k].face != hands[j].cards[k].face {
				return hands[i].cards[k].value < hands[j].cards[k].value
			}
		}

		return false
	})
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

	hands := make([]*hand, 0)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("unable to read line: %w", err)
		}

		line = strings.TrimSuffix(line, "\n")

		lineSplit := strings.Split(line, " ")

		h, err := newHand(lineSplit[0], lineSplit[1])
		if err != nil {
			log.Fatalf("unable to create hand from line %s: %v", line, err)
		}

		hands = append(hands, h)
	}

	sortHands(hands)

	for _, h := range hands {
		fmt.Println(h, h.strength())
	}

	var totalWinnings int
	for i, h := range hands {
		totalWinnings += (i + 1) * h.bids
	}

	fmt.Println(totalWinnings)
}
