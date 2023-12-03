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
	expectedChunks := []struct {
		leadingSymbolCount  int
		trailingSymbolCount int
		value               string
	}{
		{
			leadingSymbolCount:  0,
			trailingSymbolCount: 2,
			value:               "123",
		},
		{
			leadingSymbolCount:  0,
			trailingSymbolCount: 1,
			value:               "456",
		},
	}

	parser := day03.NewParser()
	schematic, err := parser.ParseString("test", input)
	if err != nil {
		t.Error(err)
	}

	if nchunks := len(schematic.Chunks); nchunks != 2 {
		t.Errorf("expected 2 chunks; got %d", nchunks)
	}
	for i, expectedChunk := range expectedChunks {
		if nleadingsymbols := len(schematic.Chunks[i].LeadingSymbols); nleadingsymbols != expectedChunk.leadingSymbolCount {
			t.Errorf("expected %d leading symbols in chunk %d; got %d", expectedChunk.leadingSymbolCount, i, nleadingsymbols)
		}
		if ntrailingsymbols := len(schematic.Chunks[i].TrailingSymbols); ntrailingsymbols != expectedChunk.trailingSymbolCount {
			t.Errorf("expected %d trailing symbols in chunk %d; got %d", expectedChunk.trailingSymbolCount, i, ntrailingsymbols)
		}
		if value := schematic.Chunks[i].Number.Value; value != expectedChunk.value {
			t.Errorf("expected value in chunk %d to be %s; got %s", i, expectedChunk.value, value)
		}
	}
}
