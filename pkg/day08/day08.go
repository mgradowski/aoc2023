package day08

import (
	"cmp"
	"errors"
	"slices"
	"unicode/utf8"

	"github.com/EduardGomezEscandell/algo/algo"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Instruction uint8

type NodeId uint16

type Node struct {
	Id      NodeId `parser:"@NodeName '=' '(' "`
	LeftId  NodeId `parser:"@NodeName ','     "`
	RightId NodeId `parser:"@NodeName ')'     "`
}

type Network struct {
	Instructions []Instruction `parser:"@InstructionCode+ Whitespace"`
	Nodes        []Node        `parser:"@@+"`
}

var _ participle.Capture = new(Instruction)
var _ participle.Capture = new(NodeId)

var (
	ErrInvalidInstruction = errors.New("invalid instruction")
	ErrInvalidNodeName    = errors.New("invalid node name")
	ErrNodeNotFound       = errors.New("node not found")
)

var (
	InstructionLeft, _  = NewInstruction('L')
	InstructionRight, _ = NewInstruction('R')
	NodeIdAAA, _        = NewNodeId("AAA")
	NodeIdZZZ, _        = NewNodeId("ZZZ")
)

func NewInstruction(instructionCode rune) (result Instruction, err error) {
	switch instructionCode {
	case 'L':
		result = 0
	case 'R':
		result = 1
	default:
		err = ErrInvalidInstruction
	}
	return
}

func (instruction *Instruction) Capture(values []string) error {
	instructionCode, _ := utf8.DecodeRuneInString(values[0])
	if instructionCode == utf8.RuneError {
		return ErrInvalidInstruction
	}
	result, err := NewInstruction(instructionCode)
	*instruction = result
	return err
}

func NewNodeId(nodeName string) (result NodeId, err error) {
	if len(nodeName) != 3 {
		err = ErrInvalidNodeName
		return
	}
	nodeId := 0
	for _, letter := range nodeName {
		if letter < 'A' || letter > 'Z' {
			err = ErrInvalidNodeName
			return
		}
		nodeId = int(letter-'A') + int('Z'-'A'+1)*nodeId
	}
	result = NodeId(nodeId)
	return
}

func (nodeId *NodeId) Capture(values []string) error {
	result, err := NewNodeId(values[0])
	*nodeId = result
	return err
}

func (nodeId NodeId) EndsWithA() bool {
	return int(nodeId)%int('Z'-'A'+1) == 0
}

func (nodeId NodeId) EndsWithZ() bool {
	return int(nodeId)%int('Z'-'A'+1) == int('Z'-'A')
}

func (network *Network) Part1() (result int, err error) {
	slices.SortFunc[[]Node](
		network.Nodes,
		func(n1, n2 Node) int {
			return cmp.Compare[NodeId](n1.Id, n2.Id)
		},
	)
	currentNodeIndex := 0
	for {
		for _, instruction := range network.Instructions {
			var nextNodeId NodeId
			switch instruction {
			case InstructionLeft:
				nextNodeId = network.Nodes[currentNodeIndex].LeftId
			case InstructionRight:
				nextNodeId = network.Nodes[currentNodeIndex].RightId
			}
			nextNodeIndex, ok := slices.BinarySearchFunc[[]Node](
				network.Nodes,
				nextNodeId,
				func(haystack Node, needle NodeId) int {
					return cmp.Compare[NodeId](haystack.Id, needle)
				},
			)
			if !ok {
				err = ErrNodeNotFound
				return
			}
			currentNodeIndex = nextNodeIndex
			result++
			if network.Nodes[currentNodeIndex].Id == NodeIdZZZ {
				return
			}
		}
	}
}

func (network *Network) Part2() (result int, err error) {
	slices.SortFunc[[]Node](
		network.Nodes,
		func(n1, n2 Node) int {
			return cmp.Compare[NodeId](n1.Id, n2.Id)
		},
	)

	startNodeIndices := make([]int, 0)
	for i, node := range network.Nodes {
		if node.Id.EndsWithA() {
			startNodeIndices = append(startNodeIndices, i)
		}
	}

	loopLengths := make([]int, 0)
	for _, currentNodeIndex := range startNodeIndices {
		loopLength := 0
	outer:
		for {
			for _, instruction := range network.Instructions {
				var nextNodeId NodeId
				switch instruction {
				case InstructionLeft:
					nextNodeId = network.Nodes[currentNodeIndex].LeftId
				case InstructionRight:
					nextNodeId = network.Nodes[currentNodeIndex].RightId
				}
				nextNodeIndex, ok := slices.BinarySearchFunc[[]Node](
					network.Nodes,
					nextNodeId,
					func(haystack Node, needle NodeId) int {
						return cmp.Compare[NodeId](haystack.Id, needle)
					},
				)
				if !ok {
					err = ErrNodeNotFound
					return
				}
				currentNodeIndex = nextNodeIndex
				loopLength++
				if network.Nodes[currentNodeIndex].Id.EndsWithZ() {
					break outer
				}
			}
		}
		loopLengths = append(loopLengths, loopLength)
	}
	result = algo.LCM[int](loopLengths[0], loopLengths[1], loopLengths[2:]...)
	return
}

func NewParser() *participle.Parser[Network] {
	lexerDefinition := lexer.MustStateful(lexer.Rules{
		"Root": {
			{Name: "InstructionCode", Pattern: `[LR]`},
			{Name: "Whitespace", Pattern: `\s+`, Action: lexer.Push("Nodes")},
		},
		"Nodes": {
			{Name: "NodeName", Pattern: `[A-Z]{3}`},
			{Name: "Interpunction", Pattern: `[(),=]`},
			{Name: "Whitespace", Pattern: `\s+`},
		},
	})

	return participle.MustBuild[Network](
		participle.Elide("Whitespace"),
		participle.Lexer(lexerDefinition),
	)
}
