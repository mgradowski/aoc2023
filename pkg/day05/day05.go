package day05

import (
	"errors"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"mgradow.ski/aoc2023/pkg/collections"
)

type Range struct {
	Start  int `parser:"@Int"`
	Length int `parser:"@Int"`
}

type MapRange struct {
	DestinationRangeStart int `parser:"@Int"`
	SourceRangeStart      int `parser:"@Int"`
	RangeLength           int `parser:"@Int"`
}

type Map struct {
	MapRanges []*MapRange `parser:"Name KeywordMap @@*"`
}

type Almanac struct {
	Seeds []Range `parser:"KeywordSeeds @@+"`
	Maps  []*Map  `parser:"@@*"`
}

var ErrOutOfRange = errors.New("input out of range")

func RangeFromEndpoints(endpoint1 int, endpoint2 int) Range {
	if endpoint1 > endpoint2 {
		endpoint1, endpoint2 = endpoint2, endpoint1
	}
	return Range{
		Start:  endpoint1,
		Length: endpoint2 - endpoint1 + 1,
	}
}

func (range_ *Range) Endpoints() (int, int) {
	return range_.Start, range_.Start + range_.Length - 1
}

func (mapRange *MapRange) Map(n int) (int, error) {
	if n < mapRange.SourceRangeStart || n >= mapRange.SourceRangeStart+mapRange.RangeLength {
		return 0, ErrOutOfRange
	}
	return mapRange.DestinationRangeStart + (n - mapRange.SourceRangeStart), nil
}

func (mapRange *MapRange) SourceEndpoints() (int, int) {
	sourceRange := Range{Start: mapRange.SourceRangeStart, Length: mapRange.RangeLength}
	return sourceRange.Endpoints()
}

// Maps the overlap between mapRange and range_ to a new range; returns
// the the non-overlapping part of range_ unchanged in remainder. If ok is false,
// result should be ignored because there was no overlap between range_ and mapRange.
func (mapRange *MapRange) MapRange(range_ Range) (result Range, remainder []Range, ok bool) {
	sourceRangeLeft, sourceRangeRight := mapRange.SourceEndpoints()
	rangeLeft, rangeRight := range_.Endpoints()

	if rangeRight < sourceRangeLeft || rangeLeft > sourceRangeRight {
		// No overlap, remainder-only
		remainder = append(
			remainder,
			range_,
		)
		return
	}
	if rangeLeft >= sourceRangeLeft && rangeRight <= sourceRangeRight {
		// range is a subset of mapRange, no remainder
		resultStart, _ := mapRange.Map(rangeLeft)
		result = Range{
			Start:  resultStart,
			Length: range_.Length,
		}
		ok = true
		return
	}
	if rangeLeft < sourceRangeLeft && rangeRight > sourceRangeRight {
		// mapRange is a subset of range, result & remainder
		result = Range{
			Start:  mapRange.DestinationRangeStart,
			Length: mapRange.RangeLength,
		}
		remainder = append(remainder, RangeFromEndpoints(rangeLeft, sourceRangeLeft-1))
		remainder = append(remainder, RangeFromEndpoints(sourceRangeRight+1, rangeRight))
		ok = true
		return
	}
	if rangeLeft < sourceRangeLeft && rangeRight >= sourceRangeLeft {
		// Left overlap, result & remainder
		result = Range{
			Start:  mapRange.DestinationRangeStart,
			Length: rangeRight - sourceRangeLeft + 1,
		}
		remainder = append(remainder, RangeFromEndpoints(rangeLeft, sourceRangeLeft-1))
		ok = true
		return
	}
	// Right overlap, result & remainder
	resultStart, _ := mapRange.Map(rangeLeft)
	result = Range{
		Start:  resultStart,
		Length: sourceRangeRight - rangeLeft + 1,
	}
	remainder = append(remainder, RangeFromEndpoints(sourceRangeRight+1, rangeRight))
	ok = true
	return
}

func (map_ *Map) Map(n int) int {
	for _, mapRange := range map_.MapRanges {
		result, err := mapRange.Map(n)
		if err == nil {
			return result
		}
	}
	return n
}

func (map_ *Map) MapRange(range_ Range) (result []Range) {
	remainderStack := collections.NewStack[Range]()
	remainderStackLast := collections.NewStack[Range]()
	remainderStackLast.Push(range_)
	for _, mapRange := range map_.MapRanges {
		for {
			range_, err := remainderStackLast.Pop()
			if err != nil {
				break
			}
			mapResult, remainder, ok := mapRange.MapRange(range_)
			if ok {
				result = append(result, mapResult)
			}
			for _, r := range remainder {
				remainderStack.Push(r)
			}
		}
		remainderStackLast, remainderStack = remainderStack, remainderStackLast
	}
	for {
		range_, err := remainderStackLast.Pop()
		if err != nil {
			break
		}
		result = append(result, range_)
	}
	return
}

func (almanac *Almanac) Map(n int) int {
	result := n
	for _, map_ := range almanac.Maps {
		result = map_.Map(result)
	}
	return result
}

func (almanac *Almanac) MapRange(range_ Range) (result []Range) {
	resultStack := collections.NewStack[Range]()
	resultStackLast := collections.NewStack[Range]()
	resultStackLast.Push(range_)
	for _, map_ := range almanac.Maps {
		for {
			range_, err := resultStackLast.Pop()
			if err != nil {
				break
			}
			resultStack.Push(map_.MapRange(range_)...)
		}
		resultStackLast, resultStack = resultStack, resultStackLast
	}
	for {
		range_, err := resultStackLast.Pop()
		if err != nil {
			break
		}
		result = append(result, range_)
	}
	return
}

func (almanac *Almanac) Part1() (result int) {
	for i, seedPair := range almanac.Seeds {
		for j, seed := range []int{seedPair.Start, seedPair.Length} {
			location := almanac.Map(seed)
			if i == 0 && j == 0 || location < result {
				result = location
			}
		}
	}
	return
}

func (almanac *Almanac) Part2() (result int) {
	for i, seed := range almanac.Seeds {
		for j, location := range almanac.MapRange(seed) {
			if i == 0 && j == 0 || location.Start < result {
				result = location.Start
			}
		}
	}
	return
}

func NewParser() *participle.Parser[Almanac] {
	lexerDefinition := lexer.MustSimple([]lexer.SimpleRule{
		{Name: "Int", Pattern: `\d+`},
		{Name: "KeywordMap", Pattern: `map:`},
		{Name: "KeywordSeeds", Pattern: `seeds:`},
		{Name: "Name", Pattern: `[a-z-]+`},
		{Name: "Whitespace", Pattern: `\s+`},
	})

	return participle.MustBuild[Almanac](
		participle.Elide("Whitespace"),
		participle.Lexer(lexerDefinition),
	)
}
