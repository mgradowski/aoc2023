package main

import (
	"fmt"
	"log"
	"os"

	"mgradow.ski/aoc2023/pkg/day05"
)

func main() {
	parser := day05.NewParser()

	almanac, err := parser.Parse("stdin", os.Stdin)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(almanac.Part2())
}
