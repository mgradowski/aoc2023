package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
	"mgradow.ski/aoc2023/pkg/day02"
)

var defaultMaxColors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	parser := participle.MustBuild[day02.Game]()
	scanner := bufio.NewScanner(os.Stdin)
	result := 0

	for scanner.Scan() {
		game, err := parser.ParseString("stdin", scanner.Text())
		if err != nil {
			log.Panic(err)
		}

		if !game.IsLegal(defaultMaxColors) {
			continue
		}
		result += game.Id
	}

	fmt.Println(result)
}
