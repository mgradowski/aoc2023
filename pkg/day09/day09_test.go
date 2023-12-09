package day09_test

import (
	"testing"

	"mgradow.ski/aoc2023/pkg/day09"
)

const example = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestPart12(t *testing.T) {
	parser := day09.NewParser()

	report, err := parser.ParseString("example", example)
	if err != nil {
		t.Error(err)
	}
	if got, wanted := report.Part1(), 114; got != wanted {
		t.Errorf("expected report.Part1() to be %d; got %d", wanted, got)
	}
	if got, wanted := report.Part2(), 2; got != wanted {
		t.Errorf("expected report.Part2() to be %d; got %d", wanted, got)
	}
}
