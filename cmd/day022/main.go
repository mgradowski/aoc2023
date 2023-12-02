package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/mgradowski/aoc2023/pkg/day02"
)

func main() {
	parser := participle.MustBuild[day02.Game]()
	scanner := bufio.NewScanner(os.Stdin)
	result := 0

	for scanner.Scan() {
		game, err := parser.ParseString("stdin", scanner.Text())
		if err != nil {
			log.Panic(err)
		}
		result += game.Power()
	}

	fmt.Println(result)
}
