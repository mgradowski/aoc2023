package day03

import (
	"strconv"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Number struct {
	Value string `parser:"@Number"`
	Pos   lexer.Position
}

type Symbol struct {
	Value string `parser:"@Symbol"`
	Pos   lexer.Position
}

type Schematic struct {
	Numbers []*Number `parser:"( @@   "`
	Symbols []*Symbol `parser:"| @@ )*"`
}

type Position struct {
	Line   int
	Column int
}

var positionOffsets = []struct {
	DeltaLine   int
	DeltaColumn int
}{
	{+0, +1},
	{+0, -1},
	{+1, +0},
	{+1, +1},
	{+1, -1},
	{-1, +0},
	{-1, +1},
	{-1, -1},
}

func buildPositionOffsetMap(numbers []*Number) map[Position]int {
	result := make(map[Position]int)
	for _, number := range numbers {
		for i := range number.Value {
			position := Position{
				Line:   number.Pos.Line,
				Column: number.Pos.Column + i,
			}
			result[position] = number.Pos.Offset
		}
	}
	return result
}

func buildOffsetValueMap(numbers []*Number) map[int]int {
	result := make(map[int]int)
	for _, number := range numbers {
		value, _ := strconv.Atoi(number.Value)
		result[number.Pos.Offset] = value
	}
	return result
}

func calculateLegalValueSum(symbols []*Symbol, positionOffsetMap map[Position]int, offsetValueMap map[int]int) (result int) {
	uniqueLegalOffsets := make(map[int]struct{})
	for _, s := range symbols {
		for _, po := range positionOffsets {
			neighbouringPosition := Position{
				Line:   s.Pos.Line + po.DeltaLine,
				Column: s.Pos.Column + po.DeltaColumn,
			}
			if offset, ok := positionOffsetMap[neighbouringPosition]; ok {
				uniqueLegalOffsets[offset] = struct{}{}
			}
		}
	}
	for offset := range uniqueLegalOffsets {
		result += offsetValueMap[offset]
	}
	return result
}

func calculateGearRatioSum(symbols []*Symbol, positionOffsetMap map[Position]int, offsetValueMap map[int]int) (result int) {
	for _, s := range symbols {
		if s.Value != "*" {
			continue
		}
		uniqueNeighbouringOffsets := make(map[int]struct{})
		gearRatio := 1
		for _, po := range positionOffsets {
			neighbouringPosition := Position{
				Line:   s.Pos.Line + po.DeltaLine,
				Column: s.Pos.Column + po.DeltaColumn,
			}
			if offset, ok := positionOffsetMap[neighbouringPosition]; ok {
				uniqueNeighbouringOffsets[offset] = struct{}{}
			}
		}
		if len(uniqueNeighbouringOffsets) != 2 {
			continue
		}
		for offset := range uniqueNeighbouringOffsets {
			gearRatio *= offsetValueMap[offset]
		}
		result += gearRatio
	}
	return result
}

func (schematic *Schematic) Part1() (result int) {
	positionOffsetMap := buildPositionOffsetMap(schematic.Numbers)
	offsetValueMap := buildOffsetValueMap(schematic.Numbers)
	return calculateLegalValueSum(schematic.Symbols, positionOffsetMap, offsetValueMap)
}

func (schematic *Schematic) Part2() int {
	positionOffsetMap := buildPositionOffsetMap(schematic.Numbers)
	offsetValueMap := buildOffsetValueMap(schematic.Numbers)
	return calculateGearRatioSum(schematic.Symbols, positionOffsetMap, offsetValueMap)
}

func NewParser() *participle.Parser[Schematic] {
	lexerDefinition := lexer.MustSimple([]lexer.SimpleRule{
		{Name: "Whitespace", Pattern: `[\.\s]+`},
		{Name: "Number", Pattern: `\d+`},
		{Name: "Symbol", Pattern: `[^\w\.\d]`},
	})

	return participle.MustBuild[Schematic](
		participle.Elide("Whitespace"),
		participle.Lexer(lexerDefinition),
	)
}
