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

type HandRank int

// order of hand ranks is important as it increases from iota onwards
const (
	ordinary HandRank = iota // lowest rank
	highCard
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind // highest rank
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
	cardCounts := h.counts()

	var jokers int
	var hasPair bool
	var hasTriplet bool
	var hasQuad bool

	for card, cardCount := range cardCounts {
		if card.face == "J" {
			jokers = cardCount
		} else if cardCount == 5 {
			// AAAAA
			return true
		} else if cardCount == 4 {
			hasQuad = true
		} else if cardCount == 3 {
			hasTriplet = true
		} else if cardCount == 2 {
			hasPair = true
		}
	}

	// JAAAA
	if jokers == 1 && hasQuad {
		return true
	}

	// JJQQQ
	if jokers == 2 && hasTriplet {
		return true
	}

	// JJJQQ
	if jokers == 3 && hasPair {
		return true
	}

	// JJJJQ
	if jokers == 4 {
		return true
	}

	// JJJJJ
	if jokers == 5 {
		return true
	}

	return false
}

func (h hand) isFourOfAKind() bool {
	cardCounts := h.counts()

	var jokers int
	var hasTriplet bool
	var naturalPairs []card

	for card, cardCount := range cardCounts {
		if card.face == "J" {
			jokers = cardCount
		} else if cardCount == 4 {
			// AAAA2, AAAAJ
			return true
		} else if cardCount == 3 {
			hasTriplet = true
		} else if cardCount == 2 {
			naturalPairs = append(naturalPairs, card)
		}
	}

	// JKKKQ
	if jokers == 1 && hasTriplet {
		return true
	}

	// JJKKQ
	if jokers == 2 && len(naturalPairs) == 1 {
		return true
	}

	// JJJAQ
	if jokers == 3 && len(naturalPairs) == 0 {
		return true
	}

	// JJJJQ
	if jokers == 4 {
		return true
	}

	return false
}

func (h hand) isFullHouse() bool {
	cardCounts := h.counts()

	var jokers int
	var naturalPairs []card
	var hasTriplet bool

	for card, cardCount := range cardCounts {
		if card.face == "J" {
			jokers = cardCount
		} else if cardCount == 3 {
			hasTriplet = true
		} else if cardCount == 2 {
			naturalPairs = append(naturalPairs, card)
		}
	}

	// AAAKK
	if jokers == 0 && hasTriplet && len(naturalPairs) == 1 {
		return true
	}

	// JKKAA
	if jokers == 1 && len(naturalPairs) == 2 {
		return true
	}

	// JJKKK
	if jokers == 2 && hasTriplet {
		return true
	}

	// JJQKK
	if jokers == 2 && len(naturalPairs) == 1 {
		return true
	}

	// JJJQQ
	if jokers == 3 && len(naturalPairs) == 1 {
		return true
	}

	return false
}

func (h hand) isThreeOfAKind() bool {
	cardCounts := h.counts()

	var jokers int
	var naturalPairs int
	var hasTriplet bool

	for card, cardCount := range cardCounts {
		if card.face == "J" {
			jokers = cardCount
		} else if cardCount == 3 {
			hasTriplet = true
		} else if cardCount == 2 {
			naturalPairs += 1
		}
	}

	// KKKAQ
	if jokers == 0 && hasTriplet && naturalPairs == 0 {
		return true
	}

	// JKKAQ
	if jokers == 1 && naturalPairs == 1 {
		return true
	}

	// JJKQA
	if jokers == 2 && naturalPairs == 0 {
		return true
	}

	// JJJKA
	if jokers == 3 && naturalPairs == 0 {
		return true
	}

	return false
}

func (h hand) isTwoPair() bool {
	cardCounts := h.counts()

	var jokers int
	var naturalPairs int

	for card, cardCount := range cardCounts {
		if card.face == "J" {
			jokers = cardCount
		} else if cardCount == 2 {
			naturalPairs += 1
		}
	}

	// KKAQQ, JKKQQ
	if naturalPairs == 2 {
		return true
	}

	// JKKAQ
	if jokers == 1 && naturalPairs == 1 {
		return true
	}

	// JJKKQ
	if jokers == 2 && naturalPairs == 1 {
		return true
	}

	return false
}

func (h hand) isOnePair() bool {
	cardCounts := h.counts()

	var jokers int
	var naturalPairs int

	for card, cardCount := range cardCounts {
		if card.face == "J" {
			jokers = cardCount
		} else if cardCount == 2 {
			naturalPairs += 1
		}
	}

	// KKQA2
	if jokers == 0 && naturalPairs == 1 {
		return true
	}

	// J2345
	if jokers == 1 && naturalPairs == 0 {
		return true
	}

	return false
}

func (h hand) isHighCard() bool {
	cardCounts := h.counts()

	for _, cardCount := range cardCounts {
		if cardCount > 1 {
			// 11344, JJ234
			return false
		}
	}

	return true
}

func (h hand) strength() int {
	if h.isFiveOfAKind() {
		return int(fiveOfAKind)
	} else if h.isFourOfAKind() {
		return int(fourOfAKind)
	} else if h.isFullHouse() {
		return int(fullHouse)
	} else if h.isThreeOfAKind() {
		return int(threeOfAKind)
	} else if h.isTwoPair() {
		return int(twoPair)
	} else if h.isOnePair() {
		return int(onePair)
	} else if h.isHighCard() {
		return int(highCard)
	} else {
		return int(ordinary)
	}
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

	var totalWinnings int
	for i, h := range hands {
		fmt.Printf("hand %d: %s %d\n", i+1, h, h.strength())
		totalWinnings += (i + 1) * h.bids
	}

	fmt.Println(totalWinnings)
}
