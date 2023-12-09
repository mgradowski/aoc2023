package day09

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type ValueHistory struct {
	Readings []int `parser:"@Int+ (EOL | EOF)"`
}

type Report struct {
	ValueHistories []ValueHistory `parser:"@@*"`
}

// Predicts both the previous and next values from a given history.
func predictRecursive(values []int) (int, int) {
	isBase := true
	for _, value := range values {
		if value != 0 {
			isBase = false
			break
		}
	}
	if isBase {
		return 0, 0
	}
	diffSequence := make([]int, len(values)-1)
	for i := range diffSequence {
		diffSequence[i] = values[i+1] - values[i]
	}
	previous, next := predictRecursive(diffSequence)
	return values[0] - previous, values[len(values)-1] + next
}

func (report *Report) Part1() (result int) {
	for _, vh := range report.ValueHistories {
		_, next := predictRecursive(vh.Readings)
		result += next
	}
	return
}

func (report *Report) Part2() (result int) {
	for _, vh := range report.ValueHistories {
		previous, _ := predictRecursive(vh.Readings)
		result += previous
	}
	return
}

func NewParser() *participle.Parser[Report] {
	lexerDefinition := lexer.MustSimple([]lexer.SimpleRule{
		{Name: "Int", Pattern: `(-?[1-9]\d*|0)`},
		{Name: "EOL", Pattern: `\n`},
		{Name: "Whitespace", Pattern: `\s+`},
	})

	return participle.MustBuild[Report](
		participle.Elide("Whitespace"),
		participle.Lexer(lexerDefinition),
	)
}
