package day03_test

import (
	"strings"
	"testing"

	"github.com/mgradowski/aoc2023/pkg/day03"
)

const input = `
..123*...
...*456*.`

func TestLexer(t *testing.T) {
	expectedTokenTypes := []string{
		"Whitespace",
		"Number",
		"Symbol",
		"Whitespace",
		"Symbol",
		"Number",
		"Symbol",
		"Whitespace",
		"EOF",
	}

	definition := day03.NewParser().Lexer()
	reader := strings.NewReader(input)
	lexer, err := definition.Lex("test", reader)
	if err != nil {
		t.Error(err)
	}

	for _, expectedTokenType := range expectedTokenTypes {
		token, err := lexer.Next()
		if err != nil {
			t.Error(err)
		}
		if token.Type != definition.Symbols()[expectedTokenType] {
			t.Errorf("expected token type %s; got %v", expectedTokenType, token.Type)
		}
	}
}

func TestParser(t *testing.T) {
	expectedNumbers := []string{
		"123",
		"456",
	}
	expectedSymbols := []string{
		"*",
		"*",
		"*",
	}

	parser := day03.NewParser()
	schematic, err := parser.ParseString("test", input)
	if err != nil {
		t.Error(err)
	}

	if len(schematic.Numbers) != len(expectedNumbers) {
		t.Errorf(
			"expected len(schematic.Numbers) to be %d; got %d",
			len(expectedNumbers),
			len(schematic.Numbers),
		)
	}
	for i, expectedNumber := range expectedNumbers {
		if schematic.Numbers[i].Value != expectedNumber {
			t.Errorf(
				"expected number %d to be %s; got %s",
				i,
				expectedNumber,
				schematic.Numbers[i].Value,
			)
		}
	}

	if len(schematic.Symbols) != len(expectedSymbols) {
		t.Errorf(
			"expected len(schematic.Symbols) to be %d; got %d",
			len(expectedSymbols),
			len(schematic.Symbols),
		)
	}
	for i, expectedSymbol := range expectedSymbols {
		if schematic.Symbols[i].Value != expectedSymbol {
			t.Errorf(
				"expected symbol %d to be %s; got %s",
				i,
				expectedSymbol,
				schematic.Symbols[i].Value,
			)
		}
	}
}
