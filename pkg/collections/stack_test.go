package collections_test

import (
	"testing"

	"github.com/mgradowski/aoc2023/pkg/collections"
)

func TestStack(t *testing.T) {
	stack := collections.NewStack[int]()
	el1 := 1
	el2 := 2
	el3 := 3

	if sz := stack.Len(); sz != 0 {
		t.Errorf("expected stack.Len() to be 0; got %d", sz)
	}

	stack.Push(el1, el2)
	stack.Push(el3)
	if sz := stack.Len(); sz != 3 {
		t.Errorf("expected stack.Len() to be 3; got %d", sz)
	}

	for _, el := range []int{el3, el2, el1} {
		val, err := stack.Pop()
		if err != nil {
			t.Errorf("expected stack.Pop() error to be nil; got %v", err)
		}
		if val != el {
			t.Errorf("expected stack.Pop() to be %d; got %d", el, val)
		}
	}

	if _, err := stack.Pop(); err != collections.ErrStackEmpty {
		t.Errorf("expected stack.Pop() error to be `%v`; got `%v`", collections.ErrStackEmpty, err)
	}
}
