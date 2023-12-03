package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mgradowski/aoc2023/pkg/day03"
)

func main() {
	parser := day03.NewParser()
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Panic(err)
	}

	schematic, err := parser.ParseString("stdin", string(stdin))
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(schematic.Part2())
}
