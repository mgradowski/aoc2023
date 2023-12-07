package day07

import (
	"slices"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Card struct {
	value int8
}

var _ participle.Capture = &Card{}

type Hand struct {
	Cards []Card `parser:"@Card @Card @Card @Card @Card"`
	Bid   int    `parser:"@Bid"`
}

type Game struct {
	Hands []*Hand `parser:"@@+"`
}

type HandSummary struct {
	CardCounts [13]uint8
}

const (
	card2 int8 = iota
	card3 int8 = iota
	card4 int8 = iota
	card5 int8 = iota
	card6 int8 = iota
	card7 int8 = iota
	card8 int8 = iota
	card9 int8 = iota
	cardT int8 = iota
	cardJ int8 = iota
	cardQ int8 = iota
	cardK int8 = iota
	cardA int8 = iota
)

type HandType uint8

const (
	HandHighCard     HandType = iota
	HandOnePair      HandType = iota
	HandTwoPair      HandType = iota
	HandThreeOfAKind HandType = iota
	HandFullHouse    HandType = iota
	HandFourOfAKind  HandType = iota
	HandFiveOfAKind  HandType = iota
)

func (card *Card) Capture(values []string) error {
	switch values[0][0] {
	case '2':
		card.value = card2
	case '3':
		card.value = card3
	case '4':
		card.value = card4
	case '5':
		card.value = card5
	case '6':
		card.value = card6
	case '7':
		card.value = card7
	case '8':
		card.value = card8
	case '9':
		card.value = card9
	case 'T':
		card.value = cardT
	case 'J':
		card.value = cardJ
	case 'Q':
		card.value = cardQ
	case 'K':
		card.value = cardK
	case 'A':
		card.value = cardA
	}
	return nil
}

func (hand *Hand) Summarize() (result HandSummary) {
	for _, card := range hand.Cards {
		result.CardCounts[card.value]++
	}
	return
}

func (handSummary HandSummary) HandType() HandType {
	foundTwo, foundThree := false, false
	for _, count := range handSummary.CardCounts {
		if count == 5 {
			return HandFiveOfAKind
		}
		if count == 4 {
			return HandFourOfAKind
		}
		if count == 3 {
			foundThree = true
		}
		if count == 2 && foundTwo {
			return HandTwoPair
		}
		if count == 2 {
			foundTwo = true
		}
	}
	if foundTwo && foundThree {
		return HandFullHouse
	}
	if foundThree {
		return HandThreeOfAKind
	}
	if foundTwo {
		return HandOnePair
	}
	return HandHighCard
}

func (handSummary HandSummary) WildcardHandType() HandType {
	wildcardCount := handSummary.CardCounts[cardJ]

	if wildcardCount == 0 {
		return handSummary.HandType()
	}
	if wildcardCount == 5 || wildcardCount == 4 {
		return HandFiveOfAKind
	}
	if wildcardCount == 3 {
		for _, count := range handSummary.CardCounts {
			if count == 2 {
				return HandFiveOfAKind
			}
		}
		return HandFourOfAKind
	}
	if wildcardCount == 2 {
		for i, count := range handSummary.CardCounts {
			if count == 3 {
				return HandFiveOfAKind
			}
			if count == 2 && i != int(cardJ) {
				return HandFourOfAKind
			}
		}
		return HandThreeOfAKind
	}
	foundTwo := false
	for _, count := range handSummary.CardCounts {
		if count == 4 {
			return HandFiveOfAKind
		}
		if count == 3 {
			return HandFourOfAKind
		}
		if count == 2 && foundTwo {
			return HandFullHouse
		}
		if count == 2 {
			foundTwo = true
		}
	}
	if foundTwo {
		return HandThreeOfAKind
	}
	return HandOnePair
}

func cmpHands1(h1, h2 *Hand) int {
	handType1 := h1.Summarize().HandType()
	handType2 := h2.Summarize().HandType()
	if handType1 < handType2 {
		return -1
	}
	if handType1 > handType2 {
		return +1
	}
	for i := range h1.Cards {
		if h1.Cards[i].value < h2.Cards[i].value {
			return -1
		}
		if h1.Cards[i].value > h2.Cards[i].value {
			return +1
		}
	}
	return 0
}

func cmpHands2(h1, h2 *Hand) int {
	handType1 := h1.Summarize().WildcardHandType()
	handType2 := h2.Summarize().WildcardHandType()
	if handType1 < handType2 {
		return -1
	}
	if handType1 > handType2 {
		return +1
	}
	for i := range h1.Cards {
		val1 := h1.Cards[i].value
		if val1 == cardJ {
			val1 = card2 - 1
		}
		val2 := h2.Cards[i].value
		if val2 == cardJ {
			val2 = card2 - 1
		}
		if val1 < val2 {
			return -1
		}
		if val1 > val2 {
			return +1
		}
	}
	return 0
}

func (game *Game) Part1() (result int) {
	slices.SortFunc(game.Hands, cmpHands1)

	for i, hand := range game.Hands {
		rank := i + 1
		result += rank * hand.Bid
	}
	return
}

func (game *Game) Part2() (result int) {
	slices.SortFunc(game.Hands, cmpHands2)

	for i, hand := range game.Hands {
		rank := i + 1
		result += rank * hand.Bid
	}
	return
}

func NewParser() *participle.Parser[Game] {
	lexerDefinition := lexer.MustStateful(lexer.Rules{
		"Root": {
			{Name: "Card", Pattern: `[23456789TJQKA]`, Action: nil},
			{Name: "Whitespace", Pattern: `\s+`, Action: lexer.Push("Bid")},
		},
		"Bid": {
			{Name: "Bid", Pattern: `\d+`, Action: nil},
			{Name: "Whitespace", Pattern: `\s+`, Action: lexer.Pop()},
		},
	})

	return participle.MustBuild[Game](
		participle.Elide("Whitespace"),
		participle.Lexer(lexerDefinition),
	)
}
