package day04

import (
	"slices"

	"github.com/alecthomas/participle/v2"
	"github.com/mgradowski/aoc2023/pkg/collections"
)

type Card struct {
	Id             int   `parser:"'Card' @Int ':'"`
	Tickets        []int `parser:"@Int+ '|'"`
	WinningNumbers []int `parser:"@Int+"`
}

type Cards struct {
	Cards []*Card `parser:"@@*"`
}

func (card *Card) matchCount() int {
	slices.Sort(card.Tickets)
	slices.Sort(card.WinningNumbers)

	matches := 0
	for i, j := 0, 0; i < len(card.Tickets) && j < len(card.WinningNumbers); {
		if card.Tickets[i] == card.WinningNumbers[j] {
			matches++
			i++
			j++
		} else if card.Tickets[i] < card.WinningNumbers[j] {
			i++
		} else {
			j++
		}
	}
	return matches
}

func (cards *Cards) matchCounts() []int {
	result := make([]int, len(cards.Cards))
	for i, card := range cards.Cards {
		result[i] = card.matchCount()
	}
	return result
}

func (cards *Cards) Part1() int {
	result := 0
	for _, card := range cards.Cards {
		matches := card.matchCount()

		if matches == 0 {
			continue
		}

		value := 1
		for i := 1; i < matches; i++ {
			value *= 2
		}
		result += value
	}
	return result
}

func (cards *Cards) Part2() int {
	allCards := 0
	copyStack := collections.NewStack[*Card]()
	copyStack.Push(cards.Cards...)
	matchCounts := cards.matchCounts()

	for {
		card, err := copyStack.Pop()
		if err != nil {
			break
		}
		allCards++
		copyStack.Push(cards.Cards[card.Id : card.Id+matchCounts[card.Id-1]]...)
	}

	return allCards
}

func NewParser() *participle.Parser[Cards] {
	return participle.MustBuild[Cards]()
}
