package day08_test

import (
	"testing"

	"mgradow.ski/aoc2023/pkg/day08"
)

const example = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

func TestParser(t *testing.T) {
	parser := day08.NewParser()
	network, err := parser.ParseString("example", example)
	if err != nil {
		t.Error(err)
	}
	for i, wanted := range []day08.Instruction{
		day08.InstructionLeft,
		day08.InstructionLeft,
		day08.InstructionRight,
	} {
		if got := network.Instructions[i]; got != wanted {
			t.Errorf("expected network.Instructions[%d] to be %d; got %d", i, got, wanted)
		}
	}
	aaa, _ := day08.NewNodeId("AAA")
	bbb, _ := day08.NewNodeId("BBB")
	zzz, _ := day08.NewNodeId("ZZZ")
	for i, wanted := range []day08.Node{
		{aaa, bbb, bbb},
		{bbb, aaa, zzz},
		{zzz, zzz, zzz},
	} {
		if got := network.Nodes[i]; got != wanted {
			t.Errorf("expected network.Nodes[%d] to be %#v; got %#v", i, got, wanted)
		}
	}
}

func TestPart1(t *testing.T) {
	parser := day08.NewParser()
	network, err := parser.ParseString("example", example)
	if err != nil {
		t.Error(err)
	}
	wanted := 6
	got, err := network.Part1()
	if err != nil {
		t.Error(err)
	}
	if got != wanted {
		t.Errorf("expected network.Part1() to be %d; got %d", wanted, got)
	}
}

func TestEndsWithAZ(t *testing.T) {
	aaa, _ := day08.NewNodeId("AAA")
	bbb, _ := day08.NewNodeId("BBB")
	zzz, _ := day08.NewNodeId("ZZZ")

	for _, case_ := range []struct {
		Input     day08.NodeId
		EndsWithA bool
		EndsWithZ bool
	}{
		{aaa, true, false},
		{bbb, false, false},
		{zzz, false, true},
	} {
		if got := case_.Input.EndsWithA(); got != case_.EndsWithA {
			t.Errorf("expected %#v.EndsWithA() to be %v; got %v", case_.Input, case_.EndsWithA, got)
		}
		if got := case_.Input.EndsWithZ(); got != case_.EndsWithZ {
			t.Errorf("expected %#v.EndsWithZ() to be %v; got %v", case_.Input, case_.EndsWithZ, got)
		}
	}
}
