package main

import (
	"fmt"
	"log"
	"os"

	"mgradow.ski/aoc2023/pkg/day07"
)

func main() {
	parser := day07.NewParser()

	game, err := parser.Parse("stdin", os.Stdin)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(game.Part2())
}
