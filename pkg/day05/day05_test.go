package day05_test

import (
	"testing"

	"mgradow.ski/aoc2023/pkg/day05"
)

const example = `
seeds: 79 14 55 13
seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestMap(t *testing.T) {
	parser := day05.NewParser()

	almanac, err := parser.ParseString("test", example)
	if err != nil {
		t.Error(err)
	}
	range_ := almanac.Maps[0]

	cases := []struct {
		Src int
		Dst int
	}{
		{79, 81},
		{14, 14},
		{55, 57},
		{13, 13},
	}

	for _, case_ := range cases {
		got := range_.Map(case_.Src)
		if got != case_.Dst {
			t.Errorf("expected %d to map to %d; got %d", case_.Src, case_.Dst, got)
		}
	}
}

func TestPart1(t *testing.T) {
	parser := day05.NewParser()

	almanac, err := parser.ParseString("test", example)
	if err != nil {
		t.Error(err)
	}
	got := almanac.Part1()
	if got != 35 {
		t.Errorf("expected the answer to be 35; got %d", got)
	}
}

func TestMapRangeMapRange(t *testing.T) {
	mapRange := day05.MapRange{
		SourceRangeStart:      10,
		DestinationRangeStart: 100,
		RangeLength:           10,
	}

	for _, range_ := range []day05.Range{
		day05.RangeFromEndpoints(0, 9),
		day05.RangeFromEndpoints(20, 29),
	} {
		_, remainder, ok := mapRange.MapRange(range_)
		if ok {
			t.Error("expected no result")
		}
		if length := len(remainder); length != 1 {
			t.Errorf("expected remainder length to be 1; got %d", length)
		}
		if range_ != remainder[0] {
			t.Errorf("expected remainder[0] to be %#v; got %#v", range_, remainder[0])
		}

	}
	{
		range_ := day05.RangeFromEndpoints(10, 15)
		wanted := day05.RangeFromEndpoints(100, 105)
		result, remainder, ok := mapRange.MapRange(range_)
		if !ok {
			t.Error("expected a result")
		}
		if remainder != nil {
			t.Error("expected no remainder")
		}
		if result != wanted {
			t.Errorf("expected result to be %#v; got %#v", wanted, result)
		}
	}
	{
		range_ := day05.RangeFromEndpoints(0, 29)
		wanted := day05.RangeFromEndpoints(100, 109)
		result, remainder, ok := mapRange.MapRange(range_)
		if !ok {
			t.Error("expected a result")
		}
		if length := len(remainder); length != 2 {
			t.Errorf("len(remainder) to be 2; got %d", length)
		}
		if result != wanted {
			t.Errorf("expected result to be %#v; got %#v", wanted, result)
		}
		for i, wanted := range []day05.Range{
			day05.RangeFromEndpoints(0, 9),
			day05.RangeFromEndpoints(20, 29),
		} {
			if wanted != remainder[i] {
				t.Errorf("expected remainder[%d] to be %#v; got %#v", i, wanted, remainder[i])
			}
		}
	}
	for _, case_ := range []struct {
		Range           day05.Range
		WantedResult    day05.Range
		WantedRemainder day05.Range
	}{
		{
			Range:           day05.RangeFromEndpoints(0, 15),
			WantedResult:    day05.RangeFromEndpoints(100, 105),
			WantedRemainder: day05.RangeFromEndpoints(0, 9),
		},
		{
			Range:           day05.RangeFromEndpoints(15, 29),
			WantedResult:    day05.RangeFromEndpoints(105, 109),
			WantedRemainder: day05.RangeFromEndpoints(20, 29),
		},
	} {
		result, remainder, ok := mapRange.MapRange(case_.Range)
		if !ok {
			t.Error("expected a result")
		}
		if length := len(remainder); length != 1 {
			t.Errorf("len(remainder) to be 1; got %d", length)
		}
		if result != case_.WantedResult {
			t.Errorf("expected result to be %#v; got %#v", case_.WantedResult, result)
		}
		if remainder[0] != case_.WantedRemainder {
			t.Errorf("expected remainder[0] to be %#v; got %#v", case_.WantedRemainder, remainder[0])
		}
	}
}

func TestPart2(t *testing.T) {
	parser := day05.NewParser()

	almanac, err := parser.ParseString("test", example)
	if err != nil {
		t.Error(err)
	}
	got := almanac.Part2()
	if got != 46 {
		t.Errorf("expected the answer to be 35; got %d", got)
	}
}
