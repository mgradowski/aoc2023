package main

import (
	"fmt"
	"log"
	"os"

	"mgradow.ski/aoc2023/pkg/day09"
)

func main() {
	parser := day09.NewParser()

	report, err := parser.Parse("stdin", os.Stdin)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(report.Part1())
}
