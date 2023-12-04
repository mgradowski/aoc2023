package main

import (
	"fmt"
	"log"
	"os"

	"mgradow.ski/aoc2023/pkg/day04"
)

func main() {
	parser := day04.NewParser()

	cards, err := parser.Parse("stdin", os.Stdin)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(cards.Part1())
}
