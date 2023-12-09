package main

import (
	"fmt"
	"log"
	"os"

	"mgradow.ski/aoc2023/pkg/day08"
)

func main() {
	parser := day08.NewParser()

	network, err := parser.Parse("stdin", os.Stdin)
	if err != nil {
		log.Panic(err)
	}
	result, err := network.Part1()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(result)
}
