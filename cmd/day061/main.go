package main

import (
	"fmt"
	"log"
	"os"

	"mgradow.ski/aoc2023/pkg/day06"
)

func main() {
	parser := day06.NewParser()

	leaderboard, err := parser.Parse("stdin", os.Stdin)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(leaderboard.Part1())
}
